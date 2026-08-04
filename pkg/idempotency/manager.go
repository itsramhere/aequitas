package idempotency

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type IdempotencyManager struct {
	db              *sql.DB
	retryController *cc.UnifiedRetryController
}

func NewIdempotencyManager(db *sql.DB, retryController *cc.UnifiedRetryController) *IdempotencyManager {
	return &IdempotencyManager{
		db:              db,
		retryController: retryController,
	}
}

func (m *IdempotencyManager) ProcessTransfer(
	ctx context.Context,
	strategy cc.Strategy,
	params ledger.TransferParams,
	opts cc.TransferOptions,
) (*ledger.TxResult, *cc.RetryStats, error) {
	// If Stage A is disabled (Bare CC Path for Experiment Set 1), bypass idempotency checks
	if !opts.EnableStageA {
		return m.retryController.ExecuteWithRetry(ctx, m.db, strategy, params, opts)
	}

	// Stage A: Separate short transaction to register pending key
	stageATx, err := m.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to begin Stage A transaction: %w", err)
	}

	_, err = stageATx.ExecContext(ctx, `
		INSERT INTO idempotency_keys (client_id, idempotency_key, state, created_at, updated_at)
		VALUES ($1, $2, 'pending', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`, params.ClientID, params.IdempotencyKey)

	if err != nil {
		stageATx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			// Idempotency key already exists -> check collision policy
			var state ledger.IdempotencyState
			var payload []byte
			lookupErr := m.db.QueryRowContext(ctx, `
				SELECT state, response_payload 
				FROM idempotency_keys 
				WHERE client_id = $1 AND idempotency_key = $2
			`, params.ClientID, params.IdempotencyKey).Scan(&state, &payload)

			if lookupErr != nil {
				return nil, nil, fmt.Errorf("idempotency key lookup failed: %w", lookupErr)
			}

			if state == ledger.StateCommitted {
				// Request already executed successfully -> return cached result
				var cachedResult ledger.TxResult
				if len(payload) > 0 {
					_ = json.Unmarshal(payload, &cachedResult)
				}
				cachedResult.IsCached = true
				return &cachedResult, &cc.RetryStats{}, nil
			}

			if state == ledger.StatePending {
				// Request currently in-flight in Stage B -> reject with 429/409 processing error
				return nil, nil, ledger.ErrProcessingRetryLater
			}

			// If state is failed, allow fresh re-execution by falling through
		} else {
			return nil, nil, fmt.Errorf("Stage A insert failed: %w", err)
		}
	} else {
		if err := stageATx.Commit(); err != nil {
			return nil, nil, fmt.Errorf("Stage A commit failed: %w", err)
		}
	}

	// Stage B: Atomic transaction with CC strategy transfer + commit key state update
	result, stats, err := m.retryController.ExecuteWithRetry(ctx, m.db, strategy, params, opts)
	if err != nil {
		// Mark key state as failed if Stage B aborted completely
		_, _ = m.db.ExecContext(ctx, `
			UPDATE idempotency_keys 
			SET state = 'failed', updated_at = CURRENT_TIMESTAMP 
			WHERE client_id = $1 AND idempotency_key = $2 AND state = 'pending'
		`, params.ClientID, params.IdempotencyKey)
		return nil, stats, err
	}

	// Store serialized cached result payload in Stage B completion
	if payload, marshalErr := json.Marshal(result); marshalErr == nil {
		_, _ = m.db.ExecContext(ctx, `
			UPDATE idempotency_keys 
			SET response_payload = $1 
			WHERE client_id = $2 AND idempotency_key = $3
		`, payload, params.ClientID, params.IdempotencyKey)
	}

	return result, stats, nil
}
