package cc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type PessimisticStrategy struct {
	LockTimeout string // e.g. "2000ms"
}

func NewPessimisticStrategy(lockTimeout string) *PessimisticStrategy {
	if lockTimeout == "" {
		lockTimeout = "2000ms"
	}
	return &PessimisticStrategy{LockTimeout: lockTimeout}
}

func (s *PessimisticStrategy) Type() StrategyType {
	return StrategyPessimistic
}

func (s *PessimisticStrategy) IsolationLevel() sql.IsolationLevel {
	return sql.LevelReadCommitted
}

// IsDeadlockOrLockTimeout returns true if the error is a Postgres deadlock (40P01) or lock timeout (55P03/57014)
func IsDeadlockOrLockTimeout(err error) bool {
	if err == nil {
		return false
	}
	if pqErr, ok := err.(*pq.Error); ok {
		// 40P01 = deadlock_detected
		// 55P03 = lock_not_available
		// 57014 = statement_timeout (which triggers when lock wait times out under lock_timeout)
		return pqErr.Code == "40P01" || pqErr.Code == "55P03" || pqErr.Code == "57014"
	}
	return false
}

func (s *PessimisticStrategy) ExecuteTransfer(ctx context.Context, db *sql.DB, params ledger.TransferParams, opts TransferOptions) (*ledger.TxResult, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	start := time.Now()
	txOpts := &sql.TxOptions{Isolation: sql.LevelReadCommitted}
	tx, err := db.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to begin Pessimistic transaction: %w", err)
	}
	defer tx.Rollback()

	// Set session lock timeout to avoid indefinite blocking
	_, err = tx.ExecContext(ctx, fmt.Sprintf("SET LOCAL lock_timeout = '%s'", s.LockTimeout))
	if err != nil {
		return nil, fmt.Errorf("failed to set lock_timeout: %w", err)
	}

	// Step 1: Execute and time the SELECT ... FOR UPDATE statements (CC Wait phase)
	ccWaitStart := time.Now()

	firstID := params.DebitedAccountID
	secondID := params.CreditedAccountID
	if firstID > secondID {
		firstID, secondID = secondID, firstID
	}

	// Execute lock acquisitions in deterministic order
	var firstBalance, secondBalance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", firstID).Scan(&firstBalance)
	if err == sql.ErrNoRows {
		return nil, ledger.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", secondID).Scan(&secondBalance)
	if err == sql.ErrNoRows {
		return nil, ledger.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	ccWaitDuration := time.Since(ccWaitStart)

	// Step 2: Time the actual UPDATE and COMMIT statements (DB Execution phase)
	dbExecStart := time.Now()

	// Identify debited account balance for sufficiency check
	debitedBalance := firstBalance
	if params.DebitedAccountID == secondID {
		debitedBalance = secondBalance
	}

	// Check funds sufficiency
	if debitedBalance-params.Amount < 0 {
		return nil, ledger.ErrInsufficientFunds
	}

	// Apply debit and credit (safe because row locks are held)
	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.DebitedAccountID)
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", params.Amount, params.CreditedAccountID)
	if err != nil {
		return nil, err
	}

	// Step 4, 5, 6: Write entry log, update idempotency state, and commit
	txID, walWaitProxy, err := ledger.RecordEntriesAndCommit(tx, params, opts.EnableStageA)
	if err != nil {
		return nil, err
	}

	dbExecDuration := time.Since(dbExecStart)
	totalDuration := time.Since(start)
	appDuration := totalDuration - dbExecDuration - ccWaitDuration
	if appDuration < 0 {
		appDuration = 0
	}

	return &ledger.TxResult{
		TransactionID:  txID,
		ClientID:       params.ClientID,
		IdempotencyKey: params.IdempotencyKey,
		Status:         "committed",
		Attempts:       1,
		AppLatency:     appDuration,
		DBLatency:      dbExecDuration,
		CCWaitLatency:  ccWaitDuration,
		WALWaitProxy:   walWaitProxy,
		IsCached:       false,
	}, nil
}
