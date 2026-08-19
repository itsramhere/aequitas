package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/workload"
)

func main() {
	dbConnStr := flag.String("db", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable", "PostgreSQL Connection String")
	strategyName := flag.String("strategy", "SSI", "CC Strategy (SSI, PESSIMISTIC, OCC)")
	skewTheta := flag.Float64("skew", 0.8, "Zipfian Skew Coefficient Theta (0.0 to 1.2)")
	concurrency := flag.Int("concurrency", 50, "Concurrent client workers")
	duration := flag.Duration("duration", 5*time.Minute, "Measured run duration")
	warmup := flag.Duration("warmup", 1*time.Minute, "Warm-up duration to discard")
	enableStageA := flag.Bool("enable-stage-a", false, "Enable Stage A idempotency key insert (false for Exp Set 1 bare CC path)")
	useUnlogged := flag.Bool("unlogged", false, "Use UNLOGGED table diagnostic variant")
	outputFile := flag.String("output", "", "JSON file path to save cell metrics report")
	flag.Parse()

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
		strat = cc.NewPessimisticStrategy("2000ms")
	case "OCC":
		strat = cc.NewOCCStrategy()
	case "ADAPTIVE":
		adaptiveStrat := cc.NewAdaptiveStrategy("2000ms")
		adaptiveStrat.Init(context.Background())
		defer adaptiveStrat.Stop()
		strat = adaptiveStrat
	default:
		log.Fatalf("Unknown Strategy: %s", *strategyName)
	}

	retryController := cc.NewUnifiedRetryController(5, 5*time.Millisecond, 100*time.Millisecond)
	runner := workload.NewBenchmarkRunner(db, retryController)

	cfg := workload.CellConfig{
		StrategyName: *strategyName,
		Strategy:     strat,
		SkewTheta:    *skewTheta,
		Concurrency:  *concurrency,
		Duration:     *duration,
		Warmup:       *warmup,
		EnableStageA: *enableStageA,
		UseUnlogged:  *useUnlogged,
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
