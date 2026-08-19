package cc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type OCCStrategy struct{}

func NewOCCStrategy() *OCCStrategy {
	return &OCCStrategy{}
}

func (s *OCCStrategy) Type() StrategyType {
	return StrategyOCC
}

func (s *OCCStrategy) IsolationLevel() sql.IsolationLevel {
	return sql.LevelReadCommitted
}

func (s *OCCStrategy) ExecuteTransfer(ctx context.Context, db *sql.DB, params ledger.TransferParams, opts TransferOptions) (*ledger.TxResult, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	start := time.Now()
	txOpts := &sql.TxOptions{Isolation: sql.LevelReadCommitted}
	tx, err := db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to begin OCC transaction: %w", err)
	}
	defer tx.Rollback()

	dbStart := time.Now()

	// Step 1: Read debited account balance and version (plain SELECT, no locks)
	var debitedBalance float64
	var debitedVersion int64
	err = tx.QueryRowContext(ctx, "SELECT balance, version FROM accounts WHERE id = $1", params.DebitedAccountID).Scan(&debitedBalance, &debitedVersion)
	if err == sql.ErrNoRows {
		return nil, ledger.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	// Read credited account version (plain SELECT, no locks)
	var creditedBalance float64
	var creditedVersion int64
	err = tx.QueryRowContext(ctx, "SELECT balance, version FROM accounts WHERE id = $1", params.CreditedAccountID).Scan(&creditedBalance, &creditedVersion)
	if err == sql.ErrNoRows {
		return nil, ledger.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	// Step 2: Check funds sufficiency
	if debitedBalance-params.Amount < 0 {
		return nil, ledger.ErrInsufficientFunds
	}

	// Step 3: Atomic single-statement Compare-And-Swap (CAS) in ascending account ID order to eliminate AB-BA deadlocks
	type casUpdate struct {
		accountID int64
		version   int64
		isDebit   bool
	}

	u1 := casUpdate{accountID: params.DebitedAccountID, version: debitedVersion, isDebit: true}
	u2 := casUpdate{accountID: params.CreditedAccountID, version: creditedVersion, isDebit: false}
	if u1.accountID > u2.accountID {
		u1, u2 = u2, u1
	}

	for _, u := range []casUpdate{u1, u2} {
		var res sql.Result
		var err error
		if u.isDebit {
			res, err = tx.ExecContext(ctx, `
				UPDATE accounts 
				SET balance = balance - $1, version = version + 1, updated_at = CURRENT_TIMESTAMP 
				WHERE id = $2 AND version = $3
			`, params.Amount, u.accountID, u.version)
		} else {
			res, err = tx.ExecContext(ctx, `
				UPDATE accounts 
				SET balance = balance + $1, version = version + 1, updated_at = CURRENT_TIMESTAMP 
				WHERE id = $2 AND version = $3
			`, params.Amount, u.accountID, u.version)
		}
		if err != nil {
			return nil, err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return nil, err
		}
		if rows == 0 {
			// Concurrent writer modified account version; abort & trigger retry
			return nil, ledger.ErrVersionMismatch
		}
	}

	// Step 4, 5, 6: Record entries, update idempotency key, and commit
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
