package workload

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/idempotency"
	"github.com/ledger/skewed-ledger/pkg/ledger"
	"github.com/ledger/skewed-ledger/pkg/telemetry"
)

type CellConfig struct {
	StrategyName string
	Strategy     cc.Strategy
	SkewTheta    float64
	Concurrency  int
	Duration     time.Duration
	Warmup       time.Duration
	EnableStageA bool
	UseUnlogged  bool
}

type BenchmarkRunner struct {
	db                 *sql.DB
	idempotencyManager *idempotency.IdempotencyManager
	dbCleaner          *DBCleaner
	statsCollector     *telemetry.DBStatsCollector
}

func NewBenchmarkRunner(db *sql.DB, retryController *cc.UnifiedRetryController) *BenchmarkRunner {
	return &BenchmarkRunner{
		db:                 db,
		idempotencyManager: idempotency.NewIdempotencyManager(db, retryController),
		dbCleaner:          NewDBCleaner(db),
		statsCollector:     telemetry.NewDBStatsCollector(db),
	}
}

func (r *BenchmarkRunner) RunCell(ctx context.Context, cfg CellConfig) (*telemetry.CellMetricsReport, error) {
	log.Printf("=== Starting Experiment Cell: Strategy=%s, Skew=%.2f, Concurrency=%d, StageA=%v, Unlogged=%v ===",
		cfg.StrategyName, cfg.SkewTheta, cfg.Concurrency, cfg.EnableStageA, cfg.UseUnlogged)

	// 1. Storage & Autovacuum setup
	if err := r.dbCleaner.DisableAutovacuum(ctx); err != nil {
		return nil, fmt.Errorf("failed to disable autovacuum: %w", err)
	}

	if err := r.dbCleaner.ToggleUnloggedVariant(ctx, cfg.UseUnlogged); err != nil {
		return nil, fmt.Errorf("failed to set unlogged variant: %w", err)
	}

	// 2. Seed Accounts (10,000 accounts with $1,000.00 each)
	log.Println("[Runner] Seeding 10,000 accounts...")
	if _, err := r.db.ExecContext(ctx, "SELECT seed_accounts(10000, 1000.00)"); err != nil {
		return nil, fmt.Errorf("failed to seed accounts: %w", err)
	}

	// 3. Pre-run Dead Tuple Count
	preDeadTuples, err := r.statsCollector.GetDeadTupleCount(ctx)
	if err != nil {
		log.Printf("[Runner Warning] Failed to get initial dead tuple count: %v", err)
	}

	// 4. Instantiate Zipf generator
	zipfGen := NewZipfGenerator(10000, cfg.SkewTheta)

	// Operational atomic counters
	var totalRequests int64
	var committedTxns int64
	var clientFailures int64
	var internalRetries int64
	var deadlockCount int64

	// Duration arrays for percentile telemetry
	var appLatencies []time.Duration
	var dbLatencies []time.Duration
	var walProxies []time.Duration
	var e2eLatencies []time.Duration
	var latMutex sync.Mutex

	recordMetrics := func(res *ledger.TxResult, stats *cc.RetryStats, e2e time.Duration, err error) {
		atomic.AddInt64(&totalRequests, 1)
		if stats != nil {
			atomic.AddInt64(&internalRetries, int64(stats.InternalRetries))
			atomic.AddInt64(&deadlockCount, int64(stats.DeadlockCount))
		}

		if err == nil && res != nil {
			atomic.AddInt64(&committedTxns, 1)
			latMutex.Lock()
			appLatencies = append(appLatencies, res.AppLatency)
			dbLatencies = append(dbLatencies, res.DBLatency)
			walProxies = append(walProxies, res.WALWaitProxy)
			e2eLatencies = append(e2eLatencies, e2e)
			latMutex.Unlock()
		} else {
			atomic.AddInt64(&clientFailures, 1)
		}
	}

	// 5. Worker Function
	worker := func(workerID int, stopCh <-chan struct{}, isWarmup *atomic.Bool) {
		clientID := fmt.Sprintf("client-%d", workerID)

		for {
			select {
			case <-stopCh:
				return
			default:
				debitedID := zipfGen.NextAccountID()
				creditedID := zipfGen.NextAccountID()
				for debitedID == creditedID {
					creditedID = zipfGen.NextAccountID()
				}

				params := ledger.TransferParams{
					ClientID:          clientID,
					IdempotencyKey:    uuid.New().String(),
					DebitedAccountID:  debitedID,
					CreditedAccountID: creditedID,
					Amount:            1.00 + rand.Float64()*10.00,
				}

				opts := cc.TransferOptions{EnableStageA: cfg.EnableStageA}

				e2eStart := time.Now()
				res, stats, err := r.idempotencyManager.ProcessTransfer(ctx, cfg.Strategy, params, opts)
				e2eDuration := time.Since(e2eStart)

				if !isWarmup.Load() {
					recordMetrics(res, stats, e2eDuration, err)
				}
			}
		}
	}

	// 6. Launch Concurrency Workers
	stopCh := make(chan struct{})
	var isWarmup atomic.Bool
	isWarmup.Store(true)

	var wg sync.WaitGroup
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, stopCh, &isWarmup)
		}(i)
	}

	// 7. Execute Warm-up Period
	if cfg.Warmup > 0 {
		log.Printf("[Runner] Executing warm-up period (%v)...", cfg.Warmup)
		time.Sleep(cfg.Warmup)
	}

	// 8. Execute Measured Cell Run Window
	log.Printf("[Runner] Starting measured run window (%v)...", cfg.Duration)
	isWarmup.Store(false)
	time.Sleep(cfg.Duration)

	// Stop workers
	close(stopCh)
	wg.Wait()
	log.Println("[Runner] Measured run window completed.")

	// 9. Post-run Dead Tuple Count
	postDeadTuples, err := r.statsCollector.GetDeadTupleCount(ctx)
	if err != nil {
		log.Printf("[Runner Warning] Failed to get post dead tuple count: %v", err)
	}

	deadTuplesDiff := postDeadTuples - preDeadTuples
	if deadTuplesDiff < 0 {
		deadTuplesDiff = 0
	}

	// 10. Inter-cell Cleanup
	if err := r.dbCleaner.ExplicitVacuum(ctx); err != nil {
		log.Printf("[Runner Warning] Inter-cell explicit VACUUM failed: %v", err)
	}

	// 11. Compile Metrics Report
	durationSec := cfg.Duration.Seconds()
	tps := float64(committedTxns) / durationSec

	var abortRate float64
	if totalRequests > 0 {
		abortRate = (float64(internalRetries+clientFailures) / float64(totalRequests)) * 100.0
	}

	var deadTuplesPerCommit float64
	if committedTxns > 0 {
		deadTuplesPerCommit = float64(deadTuplesDiff) / float64(committedTxns)
	}

	return &telemetry.CellMetricsReport{
		Strategy:              cfg.StrategyName,
		SkewTheta:             cfg.SkewTheta,
		Concurrency:           cfg.Concurrency,
		Duration:              cfg.Duration,
		TotalRequests:         totalRequests,
		CommittedTxns:         committedTxns,
		ClientVisibleFailures: clientFailures,
		InternalRetries:       internalRetries,
		DeadlockCount:         deadlockCount,
		Throughput:            tps,
		AbortRetryRate:        abortRate,
		RealizedHotRowHitRate: zipfGen.RealizedHotRowHitRate(),
		DeadTuplesGenerated:   deadTuplesDiff,
		DeadTuplesPerCommit:   deadTuplesPerCommit,
		AppLatency:            telemetry.CalculatePercentiles(appLatencies),
		DBLatency:             telemetry.CalculatePercentiles(dbLatencies),
		WALWaitProxy:          telemetry.CalculatePercentiles(walProxies),
		ClientEndToEnd:        telemetry.CalculatePercentiles(e2eLatencies),
		RuntimeSettings:       telemetry.GetGCRuntimeSettings(),
	}, nil
}
