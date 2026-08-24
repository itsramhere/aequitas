package workload

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
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
	// StatementTimeout, when non-empty, is applied per transaction via
	// SET LOCAL (ADR-07). Must satisfy MaxAttempts*timeout < pending key TTL.
	StatementTimeout string
	// DuplicateRate in [0,1]: probability that a request replays one of the
	// worker's own recent idempotency keys instead of minting a fresh UUID.
	// Only meaningful with EnableStageA; without it, duplicate handling,
	// cached responses, and 429 collisions are unreachable by construction.
	DuplicateRate float64
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
	// Restore default autovacuum when the cell (and thus the whole run) ends,
	// so suppression never leaks past the benchmark harness (ADR-11).
	defer func() {
		if err := r.dbCleaner.RestoreAutovacuum(context.Background()); err != nil {
			log.Printf("[Runner Warning] Failed to restore autovacuum: %v", err)
		}
	}()

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

	// Phase deadlines: a request's metrics bucket is decided by its START time,
	// so requests can neither straddle the warm-up boundary into the measured
	// window nor drain past its end into the throughput denominator.
	phaseStart := time.Now()
	warmupEnd := phaseStart.Add(cfg.Warmup)
	measuredEnd := warmupEnd.Add(cfg.Duration)

	// Operational atomic counters
	var totalRequests int64
	var committedTxns int64
	var clientFailures int64
	var internalRetries int64
	var deadlockCount int64
	var collisionRejections int64
	var cachedHits int64
	var insufficientFundsCount int64

	// Duration arrays for percentile telemetry
	var appLatencies []time.Duration
	var dbLatencies []time.Duration
	var ccWaitLatencies []time.Duration
	var walProxies []time.Duration
	var e2eLatencies []time.Duration
	var collisionLatencies []time.Duration
	var latMutex sync.Mutex

	recordMetrics := func(res *ledger.TxResult, stats *cc.RetryStats, e2e time.Duration, err error) {
		atomic.AddInt64(&totalRequests, 1)
		if stats != nil {
			atomic.AddInt64(&internalRetries, int64(stats.InternalRetries))
			atomic.AddInt64(&deadlockCount, int64(stats.DeadlockCount))
			atomic.AddInt64(&insufficientFundsCount, int64(stats.InsufficientFundsCount))
		}

		if err != nil {
			if errors.Is(err, ledger.ErrProcessingRetryLater) {
				// Pending-collision rejection: a deliberate 429-class outcome that
				// must not be counted as a CC failure (ADR-17).
				atomic.AddInt64(&collisionRejections, 1)
				latMutex.Lock()
				collisionLatencies = append(collisionLatencies, e2e)
				latMutex.Unlock()
			} else {
				atomic.AddInt64(&clientFailures, 1)
			}
			return
		}
		if res == nil {
			atomic.AddInt64(&clientFailures, 1)
			return
		}

		// Latency hygiene: only fresh committed executions carry real strategy
		// latency measurements. Cached duplicates carry zero-valued durations by
		// construction (the payload persists IDs/status only), and business
		// rejections never entered a Stage B transaction — appending either as
		// latency samples fabricated sub-microsecond percentiles.
		if res.IsCached {
			atomic.AddInt64(&cachedHits, 1)
			return
		}
		if res.Status == "committed" {
			atomic.AddInt64(&committedTxns, 1)
			latMutex.Lock()
			appLatencies = append(appLatencies, res.AppLatency)
			dbLatencies = append(dbLatencies, res.DBLatency)
			ccWaitLatencies = append(ccWaitLatencies, res.CCWaitLatency)
			walProxies = append(walProxies, res.WALWaitProxy)
			e2eLatencies = append(e2eLatencies, e2e)
			latMutex.Unlock()
		}
	}

	// recentKeyRingSize bounds how far back a replay may reach; keys stay
	// worker-local, which preserves the (client_id, idempotency_key) uniqueness
	// domain because client IDs are also worker-local.
	const recentKeyRingSize = 256

	// 5. Worker Function (each worker gets a dedicated Zipf sampler with its
	// own PRNG to avoid shared math/rand mutex contention at high concurrency)
	worker := func(workerID int, stopCh <-chan struct{}) {
		clientID := fmt.Sprintf("client-%d", workerID)
		sampler := zipfGen.NewSampler(int64(workerID) + 1)
		amountRnd := rand.New(rand.NewSource(int64(workerID) + 1_000_000))
		recentKeys := make([]string, 0, recentKeyRingSize)

		for {
			select {
			case <-stopCh:
				return
			default:
			}

			now := time.Now()
			if !now.Before(measuredEnd) {
				return
			}
			inWarmup := now.Before(warmupEnd)

			debitedID := sampler.NextAccountID()
			creditedID := sampler.NextAccountID()
			for debitedID == creditedID {
				creditedID = sampler.NextAccountID()
			}

			// Quantize amounts to whole cents so every amount is stored and
			// summed identically on both sides of the auditor's reconciliation
			// (NUMERIC(18,4)); full-precision float64 amounts previously made
			// balance-storage rounding diverge from entry-log rounding.
			amount := math.Round((1.00+amountRnd.Float64()*10.00)*100) / 100

			params := ledger.TransferParams{
				ClientID:          clientID,
				IdempotencyKey:    uuid.New().String(),
				DebitedAccountID:  debitedID,
				CreditedAccountID: creditedID,
				Amount:            amount,
			}

			// Duplicate injection (ADR-06 revision): with probability p, replay
			// one of this worker's own recent keys so the idempotency machinery
			// (unique-index collisions, cached responses, 429s, failed-key
			// reclaim) is actually exercised and auditable.
			if cfg.EnableStageA && cfg.DuplicateRate > 0 && len(recentKeys) > 0 && amountRnd.Float64() < cfg.DuplicateRate {
				params.IdempotencyKey = recentKeys[amountRnd.Intn(len(recentKeys))]
			} else if cfg.EnableStageA {
				if len(recentKeys) < recentKeyRingSize {
					recentKeys = append(recentKeys, params.IdempotencyKey)
				} else {
					recentKeys[amountRnd.Intn(recentKeyRingSize)] = params.IdempotencyKey
				}
			}

			opts := cc.TransferOptions{
				EnableStageA:     cfg.EnableStageA,
				StatementTimeout: cfg.StatementTimeout,
			}

			e2eStart := time.Now()
			res, stats, err := r.idempotencyManager.ProcessTransfer(ctx, cfg.Strategy, params, opts)
			e2eDuration := time.Since(e2eStart)

			if !inWarmup {
				recordMetrics(res, stats, e2eDuration, err)
			}
		}
	}

	// 6. Launch Concurrency Workers
	stopCh := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, stopCh)
		}(i)
	}

	// 7. Wait out the warm-up period BEFORE opening the measured window. The
	// per-second sampler must start exactly at warmupEnd: launching it early
	// emitted zero-TPS ticks labelled as measured seconds and shifted the
	// whole time_series array relative to the reported window.
	if rest := time.Until(warmupEnd); rest > 0 {
		log.Printf("[Runner] Executing warm-up period (%v)...", rest)
		time.Sleep(rest)
	}

	// Per-second time-series sampler over the measured window
	log.Printf("[Runner] Starting measured run window (%v)...", cfg.Duration)

	var timeSeries []telemetry.TimeSeriesPoint
	tsStopCh := make(chan struct{})
	var tsWg sync.WaitGroup

	tsWg.Add(1)
	go func() {
		defer tsWg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var lastTxns int64
		var lastRetries int64
		second := 0

		for {
			select {
			case <-tsStopCh:
				return
			case <-ticker.C:
				if !time.Now().Before(measuredEnd) {
					return
				}
				second++
				curTxns := atomic.LoadInt64(&committedTxns)
				curRetries := atomic.LoadInt64(&internalRetries)

				deltaTxns := curTxns - lastTxns
				deltaRetries := curRetries - lastRetries

				lastTxns = curTxns
				lastRetries = curRetries

				timeSeries = append(timeSeries, telemetry.TimeSeriesPoint{
					Second:  second,
					TPS:     float64(deltaTxns),
					Retries: deltaRetries,
				})
			}
		}
	}()

	// 8. Wait for the measured window to elapse (workers self-retire at
	// measuredEnd), then shut down cleanly.
	for time.Now().Before(measuredEnd) {
		time.Sleep(50 * time.Millisecond)
	}
	close(stopCh)
	wg.Wait()

	close(tsStopCh)
	tsWg.Wait()
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

	var retriesPerRequest float64
	var clientFailureRate float64
	if totalRequests > 0 {
		retriesPerRequest = float64(internalRetries) / float64(totalRequests)
		clientFailureRate = (float64(clientFailures) / float64(totalRequests)) * 100.0
	}

	var deadTuplesPerCommit float64
	if committedTxns > 0 {
		deadTuplesPerCommit = float64(deadTuplesDiff) / float64(committedTxns)
	}

	return &telemetry.CellMetricsReport{
		Strategy:               cfg.StrategyName,
		SkewTheta:              cfg.SkewTheta,
		Concurrency:            cfg.Concurrency,
		Duration:               cfg.Duration,
		TotalRequests:          totalRequests,
		CommittedTxns:          committedTxns,
		CollisionRejections:    collisionRejections,
		ClientVisibleFailures:  clientFailures,
		CachedHits:             cachedHits,
		InternalRetries:        internalRetries,
		DeadlockCount:          deadlockCount,
		InsufficientFundsCount: insufficientFundsCount,
		Throughput:             tps,
		RetriesPerRequest:      retriesPerRequest,
		ClientFailureRate:      clientFailureRate,
		RealizedHotRowHitRate:  zipfGen.RealizedHotRowHitRate(),
		DeadTuplesGenerated:    deadTuplesDiff,
		DeadTuplesPerCommit:    deadTuplesPerCommit,
		AppLatency:             telemetry.CalculatePercentiles(appLatencies),
		DBLatency:              telemetry.CalculatePercentiles(dbLatencies),
		CCWaitLatency:          telemetry.CalculatePercentiles(ccWaitLatencies),
		WALWaitProxy:           telemetry.CalculatePercentiles(walProxies),
		ClientEndToEnd:         telemetry.CalculatePercentiles(e2eLatencies),
		ClientBackoffRejection: telemetry.CalculatePercentiles(collisionLatencies),
		RuntimeSettings:        telemetry.GetGCRuntimeSettings(),
		TimeSeries:             timeSeries,
	}, nil
}
