package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/idempotency"
	"github.com/ledger/skewed-ledger/pkg/workload"
	_ "github.com/lib/pq"
)

// Benchmark-harness safety constants (ADR-07). The pending key TTL and the
// per-transaction statement_timeout must satisfy
// MaxAttempts*statement_timeout < pendingTTL so a Stage B transaction can
// never outlive its idempotency key. The guard below is checked against the
// retry controller's actual MaxAttempts, not duplicated magic numbers.
const (
	statementTimeout = "1500ms"
	pendingTTL       = 10 * time.Second
)

func main() {
	dbConnStr := flag.String("db", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable", "PostgreSQL Connection String")
	strategyName := flag.String("strategy", "SSI", "CC Strategy (SSI, PESSIMISTIC, OCC, ADAPTIVE)")
	skewTheta := flag.Float64("skew", 0.8, "Zipfian Skew Coefficient Theta (0.0 to 1.2)")
	concurrency := flag.Int("concurrency", 50, "Concurrent client workers")
	duration := flag.Duration("duration", 5*time.Minute, "Measured run duration")
	warmup := flag.Duration("warmup", 1*time.Minute, "Warm-up duration to discard")
	enableStageA := flag.Bool("enable-stage-a", false, "Enable Stage A idempotency key insert (false for Exp Set 1 bare CC path)")
	duplicateRate := flag.Float64("duplicate-rate", 0.0, "Probability [0,1] of replaying a recent idempotency key (requires -enable-stage-a)")
	lockTimeout := flag.String("lock-timeout", "2000ms", "Pessimistic lock_timeout setting")
	useUnlogged := flag.Bool("unlogged", false, "Use UNLOGGED table diagnostic variant (unsupported: FK constraints block persistence changes)")
	outputFile := flag.String("output", "", "JSON file path to save cell metrics report")
	flag.Parse()

	if *duplicateRate < 0 || *duplicateRate > 1 {
		log.Fatalf("invalid -duplicate-rate %v: must be within [0, 1]", *duplicateRate)
	}
	if *duplicateRate > 0 && !*enableStageA {
		// Duplicate injection keys off the idempotency machinery; without
		// Stage A there is no key table to collide against, so the flag would
		// be silently meaningless.
		log.Fatalf("-duplicate-rate requires -enable-stage-a=true (no idempotency keys exist on the bare CC path)")
	}

	db, err := sql.Open("postgres", *dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Sized strictly above 250 max concurrency
	db.SetMaxOpenConns(275)
	db.SetMaxIdleConns(275)

	var strat cc.Strategy
	switch *strategyName {
	case "SSI":
		strat = cc.NewSSIStrategy()
	case "PESSIMISTIC":
		strat = cc.NewPessimisticStrategy(*lockTimeout)
	case "OCC":
		strat = cc.NewOCCStrategy()
	case "ADAPTIVE":
		adaptiveStrat := cc.NewAdaptiveStrategy(*lockTimeout)
		adaptiveStrat.Init(context.Background())
		defer adaptiveStrat.Stop()
		strat = adaptiveStrat
	default:
		log.Fatalf("Unknown Strategy: %s", *strategyName)
	}

	retryController := cc.NewUnifiedRetryController(5, 5*time.Millisecond, 100*time.Millisecond)

	// Fail fast on the ADR-07 invariant, using the controller's real config.
	stmtTimeoutDur, err := time.ParseDuration(statementTimeout)
	if err != nil {
		log.Fatalf("Invalid statement timeout %q: %v", statementTimeout, err)
	}
	if time.Duration(retryController.MaxAttempts)*stmtTimeoutDur >= pendingTTL {
		log.Fatalf("safety constraint violation: MaxAttempts*statement_timeout (%v) must be < pendingTTL (%v)",
			time.Duration(retryController.MaxAttempts)*stmtTimeoutDur, pendingTTL)
	}

	// Feed per-attempt abort signals from the retry loop into the adaptive
	// controller's sliding window (ADR-05 / ADR-18).
	if adaptiveStrat, ok := strat.(*cc.AdaptiveStrategy); ok {
		retryController.OnRetriableAttempt = adaptiveStrat.NoteRetry
	}

	// ADR-07 revision: the TTL cleaner now runs in the benchmark harness too —
	// previously it existed only in ledger_server, so expired pending keys were
	// never reclaimed during experiments.
	var ttlCleaner *idempotency.TTLCleaner
	if *enableStageA {
		ttlCleaner, err = idempotency.NewTTLCleaner(db, pendingTTL, stmtTimeoutDur, 5*time.Second)
		if err != nil {
			log.Fatalf("TTL cleaner misconfigured: %v", err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go ttlCleaner.Start(ctx)
	}

	runner := workload.NewBenchmarkRunner(db, retryController)

	cfg := workload.CellConfig{
		StrategyName:     *strategyName,
		Strategy:         strat,
		SkewTheta:        *skewTheta,
		Concurrency:      *concurrency,
		Duration:         *duration,
		Warmup:           *warmup,
		EnableStageA:     *enableStageA,
		UseUnlogged:      *useUnlogged,
		StatementTimeout: statementTimeout,
		DuplicateRate:    *duplicateRate,
	}

	report, err := runner.RunCell(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Benchmark run failed: %v", err)
	}

	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal report: %v", err)
	}

	log.Printf("\n=== CELL BENCHMARK REPORT ===\n%s\n", string(reportJSON))

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, reportJSON, 0644); err != nil {
			log.Fatalf("Failed to write report file: %v", err)
		}
		log.Printf("Report saved to %s", *outputFile)
	}
}
