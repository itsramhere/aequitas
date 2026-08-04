package ledger

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// InsertEntriesAndCommit handles steps 4, 5, and 6 of the canonical transaction body:
// 4. Inserts parent transaction row and paired debit/credit entry rows into immutable entries table.
// 5. Updates idempotency key state to 'committed' (if Stage A is enabled).
// 6. Commits the transaction.
func RecordEntriesAndCommit(tx *sql.Tx, params TransferParams, enableStageA bool) (uuid.UUID, time.Duration, error) {
	txID := uuid.New()
	now := time.Now()

	// Step 4a: Insert parent transaction record
	_, err := tx.Exec(`
		INSERT INTO transactions (id, client_id, idempotency_key, created_at)
		VALUES ($1, $2, $3, $4)
	`, txID, params.ClientID, params.IdempotencyKey, now)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("failed to insert transaction record: %w", err)
	}

	// Step 4b: Insert paired debit/credit rows
	// Debit entry (debited account)
	_, err = tx.Exec(`
		INSERT INTO entries (transaction_id, account_id, amount, entry_type, created_at)
		VALUES ($1, $2, $3, 'debit', $4)
	`, txID, params.DebitedAccountID, params.Amount, now)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("failed to insert debit entry: %w", err)
	}

	// Credit entry (credited account)
	_, err = tx.Exec(`
		INSERT INTO entries (transaction_id, account_id, amount, entry_type, created_at)
		VALUES ($1, $2, $3, 'credit', $4)
	`, txID, params.CreditedAccountID, params.Amount, now)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("failed to insert credit entry: %w", err)
	}

	// Step 5: Update idempotency key to 'committed' if Stage A was executed
	if enableStageA {
		res, err := tx.Exec(`
			UPDATE idempotency_keys
			SET state = 'committed', updated_at = CURRENT_TIMESTAMP
			WHERE client_id = $1 AND idempotency_key = $2 AND state = 'pending'
		`, params.ClientID, params.IdempotencyKey)
		if err != nil {
			return uuid.Nil, 0, fmt.Errorf("failed to update idempotency key state: %w", err)
		}
		rows, err := res.RowsAffected()
		if err != nil || rows == 0 {
			return uuid.Nil, 0, fmt.Errorf("failed to transition idempotency key to committed (rows affected = %d)", rows)
		}
	}

	// Step 6: Commit transaction & measure WAL fsync commit proxy duration
	commitStart := time.Now()
	if err := tx.Commit(); err != nil {
		return uuid.Nil, 0, fmt.Errorf("transaction commit failed: %w", err)
	}
	walProxyDuration := time.Since(commitStart)

	return txID, walProxyDuration, nil
}
