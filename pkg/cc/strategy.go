package cc

import (
	"context"
	"database/sql"

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
}

type Strategy interface {
	ExecuteTransfer(ctx context.Context, db *sql.DB, params ledger.TransferParams, opts TransferOptions) (*ledger.TxResult, error)
	Type() StrategyType
	IsolationLevel() sql.IsolationLevel
}
