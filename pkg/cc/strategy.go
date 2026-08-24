package cc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type StrategyType string

const (
	StrategySSI         StrategyType = "SSI"
	StrategyPessimistic StrategyType = "PESSIMISTIC"
	StrategyOCC         StrategyType = "OCC"
	StrategyAdaptive    StrategyType = "ADAPTIVE"
)

type TransferOptions struct {
	EnableStageA bool // If false, skips Stage A (Experiment Set 1 bare CC path)
	// StatementTimeout, when non-empty (e.g. "1500ms"), is applied as a per-
	// transaction SET LOCAL statement_timeout. It must satisfy
	// MaxAttempts * statement_timeout + total backoff < pending key TTL, so a
	// Stage B transaction can never outlive its idempotency key (ADR-07).
	StatementTimeout string
}

type Strategy interface {
	ExecuteTransfer(ctx context.Context, db *sql.DB, params ledger.TransferParams, opts TransferOptions) (*ledger.TxResult, error)
	Type() StrategyType
	IsolationLevel() sql.IsolationLevel
}

// applyStatementTimeout enforces the per-transaction statement_timeout guard
// required by ADR-07 (statement_timeout < pending key TTL), so a live Stage B
// transaction is always aborted by Postgres before its idempotency key can
// expire. The value comes from trusted harness flags, not user input.
func applyStatementTimeout(ctx context.Context, tx *sql.Tx, opts TransferOptions) error {
	if opts.StatementTimeout == "" {
		return nil
	}
	if _, err := tx.ExecContext(ctx, fmt.Sprintf("SET LOCAL statement_timeout = '%s'", opts.StatementTimeout)); err != nil {
		return fmt.Errorf("failed to set statement_timeout: %w", err)
	}
	return nil
}
