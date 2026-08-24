package cc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type SSIStrategy struct{}

func NewSSIStrategy() *SSIStrategy {
	return &SSIStrategy{}
}

func (s *SSIStrategy) Type() StrategyType {
	return StrategySSI
}

func (s *SSIStrategy) IsolationLevel() sql.IsolationLevel {
	return sql.LevelSerializable
}

func (s *SSIStrategy) ExecuteTransfer(ctx context.Context, db *sql.DB, params ledger.TransferParams, opts TransferOptions) (*ledger.TxResult, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	start := time.Now()
	txOpts := &sql.TxOptions{Isolation: sql.LevelSerializable}
	tx, err := db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to begin SSI transaction: %w", err)
	}
	defer tx.Rollback()

	// dbStart marks the beginning of the full in-transaction window — including
	// the applyStatementTimeout SET round-trip below — matching Pessimistic so
	// DBLatency remains comparable across strategies (ADR-14).
	dbStart := time.Now()

	if err := applyStatementTimeout(ctx, tx, opts); err != nil {
		return nil, err
	}

	// Step 1: Plain SELECT balance of debited account (no FOR UPDATE in SSI)
	var debitedBalance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = $1", params.DebitedAccountID).Scan(&debitedBalance)
	if err == sql.ErrNoRows {
		return nil, ledger.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	// Step 2: Check funds sufficiency
	if debitedBalance-params.Amount < 0 {
		return nil, ledger.ErrInsufficientFunds
	}

	// Step 3: Apply debit and credit via plain UPDATE statements in ascending account ID order to eliminate AB-BA deadlocks
	if params.DebitedAccountID < params.CreditedAccountID {
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.DebitedAccountID)
		if err != nil {
			return nil, err
		}
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.CreditedAccountID)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.CreditedAccountID)
		if err != nil {
			return nil, err
		}
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.DebitedAccountID)
		if err != nil {
			return nil, err
		}
	}

	// Step 4, 5, 6: Write entry log, update idempotency state (if enabled), and commit
	txID, walWaitProxy, err := ledger.RecordEntriesAndCommit(tx, params, opts.EnableStageA)
	if err != nil {
		return nil, err
	}

	dbDuration := time.Since(dbStart)
	totalDuration := time.Since(start)

	return &ledger.TxResult{
		TransactionID:  txID,
		ClientID:       params.ClientID,
		IdempotencyKey: params.IdempotencyKey,
		Status:         "committed",
		Attempts:       1,
		AppLatency:     totalDuration - dbDuration,
		DBLatency:      dbDuration,
		WALWaitProxy:   walWaitProxy,
		IsCached:       false,
	}, nil
}
