package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/idempotency"
	"github.com/ledger/skewed-ledger/pkg/ledger"
	_ "github.com/lib/pq"
)

func main() {
	dbConnStr := flag.String("db", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable", "PostgreSQL Connection String")
	port := flag.Int("port", 8080, "HTTP Server Port")
	maxConns := flag.Int("max-conns", 275, "Max Open DB Connections (strictly > 250 clients)")
	strategyName := flag.String("strategy", "SSI", "Concurrency Control Strategy (SSI, PESSIMISTIC, OCC, ADAPTIVE)")
	lockTimeout := flag.String("lock-timeout", "2000ms", "Pessimistic lock_timeout setting")
	enableStageA := flag.Bool("enable-stage-a", true, "Enable Stage A Idempotency Key pending insert (set false for Set 1 bare CC path)")
	flag.Parse()

	db, err := sql.Open("postgres", *dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Hardcode/size connection pool strictly above maximum tested concurrency (250 clients)
	db.SetMaxOpenConns(*maxConns)
	db.SetMaxIdleConns(*maxConns)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Printf("Connected to Postgres. DB Connection Pool sized to %d max open connections.", *maxConns)

	// Instantiate strategy
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
		log.Fatalf("Unknown CC Strategy: %s", *strategyName)
	}

	const statementTimeout = "1500ms"
	stmtTimeoutDur, err := time.ParseDuration(statementTimeout)
	if err != nil {
		log.Fatalf("Invalid statement timeout %q: %v", statementTimeout, err)
	}
	const pendingTTL = 10 * time.Second

	retryController := cc.NewUnifiedRetryController(5, 5*time.Millisecond, 100*time.Millisecond)

	// Statement timeout safety guard (ADR-07): checked against the retry
	// controller's actual MaxAttempts so the invariant cannot silently drift
	// when only one of the two values is changed.
	if time.Duration(retryController.MaxAttempts)*stmtTimeoutDur >= pendingTTL {
		log.Fatalf("safety constraint violation: MaxAttempts*statement_timeout (%v) must be < pendingTTL (%v)",
			time.Duration(retryController.MaxAttempts)*stmtTimeoutDur, pendingTTL)
	}

	// Feed per-attempt abort signals from the retry loop into the adaptive
	// controller's sliding window (ADR-05 / ADR-18).
	if adaptiveStrat, ok := strat.(*cc.AdaptiveStrategy); ok {
		retryController.OnRetriableAttempt = adaptiveStrat.NoteRetry
	}

	idempotencyMgr := idempotency.NewIdempotencyManager(db, retryController)

	// Start TTL cleaner for stale pending keys
	cleaner, err := idempotency.NewTTLCleaner(db, pendingTTL, stmtTimeoutDur, 5*time.Second)
	if err != nil {
		log.Fatalf("TTL cleaner misconfigured: %v", err)
	}
	cleanerCtx, cancelCleaner := context.WithCancel(context.Background())
	defer cancelCleaner()
	go cleaner.Start(cleanerCtx)

	mux := http.NewServeMux()
	mux.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var params ledger.TransferParams
		// Bound request bodies: legitimate TransferParams payloads are a few
		// hundred bytes; oversize bodies fail through the existing decode-error
		// path (400) instead of buffering unboundedly.
		r.Body = http.MaxBytesReader(w, r.Body, 8<<10)
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		opts := cc.TransferOptions{EnableStageA: *enableStageA, StatementTimeout: statementTimeout}
		res, _, err := idempotencyMgr.ProcessTransfer(r.Context(), strat, params, opts)

		// errors.Is traverses wrapped chains (ADR-16).
		if err != nil {
			switch {
			case errors.Is(err, ledger.ErrAccountNotFound),
				errors.Is(err, ledger.ErrSameAccountTransfer),
				errors.Is(err, ledger.ErrInvalidAmount):
				w.WriteHeader(http.StatusUnprocessableEntity)
			case errors.Is(err, ledger.ErrProcessingRetryLater):
				w.WriteHeader(http.StatusTooManyRequests)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		if res != nil && res.Status == "insufficient_funds" {
			// The retry controller reports business rejections as successful
			// results with a status field; map them to 422 like their error-
			// shaped siblings instead of an unqualified 200.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(res)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%d", *port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("Ledger Service running on %s [Strategy=%s, StageA=%v]", addr, strat.Type(), *enableStageA)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
