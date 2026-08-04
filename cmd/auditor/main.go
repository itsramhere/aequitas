package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/auditor"
)

func main() {
	dbConnStr := flag.String("db", "postgres://postgres:postgres@localhost:5432/ledger?sslmode=disable", "PostgreSQL Connection String")
	initialBalance := flag.Float64("initial-balance", 1000.00, "Initial account balance for verification")
	outputFile := flag.String("output", "", "JSON file path to save audit report")
	flag.Parse()

	db, err := sql.Open("postgres", *dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	aud := auditor.NewAuditor(db)
	report, err := aud.RunAudit(context.Background(), *initialBalance)
	if err != nil {
		log.Fatalf("Audit execution failed: %v", err)
	}

	reportJSON, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal audit report: %v", err)
	}

	log.Printf("\n=== AUDIT REPORT ===\n%s\n", string(reportJSON))

	if *outputFile != "" {
		if err := os.WriteFile(*outputFile, reportJSON, 0644); err != nil {
			log.Fatalf("Failed to write audit report file: %v", err)
		}
		log.Printf("Audit report saved to %s", *outputFile)
	}

	if !report.Passed {
		os.Exit(1)
	}
}
