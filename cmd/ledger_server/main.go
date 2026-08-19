package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/idempotency"
	"github.com/ledger/skewed-ledger/pkg/ledger"
)

func main() {
	dbConnStr := flag.String("db", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable", "PostgreSQL Connection String")
	port := flag.Int("port", 8080, "HTTP Server Port")
	maxConns := flag.Int("max-conns", 275, "Max Open DB Connections (strictly > 250 clients)")
	strategyName := flag.String("strategy", "SSI", "Concurrency Control Strategy (SSI, PESSIMISTIC, OCC)")
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

	retryController := cc.NewUnifiedRetryController(5, 5*time.Millisecond, 100*time.Millisecond)
	idempotencyMgr := idempotency.NewIdempotencyManager(db, retryController)

	// Start TTL cleaner for stale pending keys
	cleaner, err := idempotency.NewTTLCleaner(db, 10*time.Second, 2*time.Second, 5*time.Second)
	if err == nil {
		go cleaner.Start(context.Background())
	}

	http.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var params ledger.TransferParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return
		}

		opts := cc.TransferOptions{EnableStageA: *enableStageA}
		res, _, err := idempotencyMgr.ProcessTransfer(r.Context(), strat, params, opts)
		if err != nil {
			if err == ledger.ErrInsufficientFunds || err == ledger.ErrAccountNotFound || err == ledger.ErrSameAccountTransfer {
				w.WriteHeader(http.StatusUnprocessableEntity)
			} else if err == ledger.ErrProcessingRetryLater {
				w.WriteHeader(http.StatusTooManyRequests)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Ledger Service running on %s [Strategy=%s, StageA=%v]", addr, strat.Type(), *enableStageA)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	os.Exit(0)
}
