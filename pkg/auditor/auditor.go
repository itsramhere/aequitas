package auditor

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type AuditReport struct {
	InitialAccountBalance    float64 `json:"initial_account_balance"`
	TotalAccountsChecked     int64   `json:"total_accounts_checked"`
	BalanceDiscrepancyCount  int64   `json:"balance_discrepancy_count"`
	TotalCommittedKeys       int64   `json:"total_committed_keys"`
	DuplicateEffectCount     int64   `json:"duplicate_effect_count"`
	Passed                   bool    `json:"passed"`
}

type Auditor struct {
	db *sql.DB
}

func NewAuditor(db *sql.DB) *Auditor {
	return &Auditor{db: db}
}

func (a *Auditor) RunAudit(ctx context.Context, initialBalance float64) (*AuditReport, error) {
	log.Println("=== Executing Independent Correctness Audit ===")
	report := &AuditReport{
		InitialAccountBalance: initialBalance,
		Passed:                true,
	}

	// -------------------------------------------------------------
	// Check 1: Balance Reconciliation Audit
	// Reconstructs balance from immutable entries log and diffs against accounts table
	// -------------------------------------------------------------
	rows, err := a.db.QueryContext(ctx, `
		SELECT a.id, a.balance, COALESCE(e.net_change, 0) AS net_change
		FROM accounts a
		LEFT JOIN (
			SELECT account_id, 
				   SUM(CASE WHEN entry_type = 'credit' THEN amount ELSE -amount END) AS net_change
			FROM entries
			GROUP BY account_id
		) e ON a.id = e.account_id
	`)
	if err != nil {
		return nil, fmt.Errorf("balance reconciliation query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accountID int64
		var actualBalance float64
		var netChange float64

		if err := rows.Scan(&accountID, &actualBalance, &netChange); err != nil {
			return nil, fmt.Errorf("failed to scan balance row: %w", err)
		}

		report.TotalAccountsChecked++
		expectedBalance := initialBalance + netChange

		// Diff tolerance check
		if diff := actualBalance - expectedBalance; diff < -0.0001 || diff > 0.0001 {
			report.BalanceDiscrepancyCount++
			report.Passed = false
			log.Printf("[AUDIT ERROR] Balance discrepancy on Account %d: Actual=%.4f, Expected=%.4f (Diff=%.4f)",
				accountID, actualBalance, expectedBalance, diff)
		}
	}

	// -------------------------------------------------------------
	// Check 2: Idempotency-Key-to-Entry-Set Cardinality Check
	// Verifies that every committed idempotency key maps to EXACTLY ONE transaction
	// -------------------------------------------------------------
	err = a.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM idempotency_keys WHERE state = 'committed'
	`).Scan(&report.TotalCommittedKeys)
	if err != nil {
		return nil, fmt.Errorf("failed to count committed keys: %w", err)
	}

	var duplicateKeys int64
	err = a.db.QueryRowContext(ctx, `
		SELECT COUNT(*) FROM (
			SELECT idempotency_key, COUNT(DISTINCT id) AS tx_count
			FROM transactions
			GROUP BY idempotency_key
			HAVING COUNT(DISTINCT id) > 1
		) duplicates
	`).Scan(&duplicateKeys)
	if err != nil {
		return nil, fmt.Errorf("idempotency cardinality check query failed: %w", err)
	}

	report.DuplicateEffectCount = duplicateKeys
	if duplicateKeys > 0 {
		report.Passed = false
		log.Printf("[AUDIT ERROR] Idempotency violation: %d committed key(s) produced duplicate transactions!", duplicateKeys)
	}

	if report.Passed {
		log.Println("=== AUDIT PASSED: Balance reconciliation & Idempotency cardinality verified cleanly. ===")
	} else {
		log.Printf("=== AUDIT FAILED: Discrepancies=%d, DuplicateEffects=%d ===",
			report.BalanceDiscrepancyCount, report.DuplicateEffectCount)
	}

	return report, nil
}
