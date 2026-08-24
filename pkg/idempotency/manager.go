package idempotency

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ledger/skewed-ledger/pkg/cc"
	"github.com/ledger/skewed-ledger/pkg/ledger"
	"github.com/lib/pq"
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
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation (ADR-16: chain traversal)
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
				// Request already executed successfully -> return cached result.
				// The payload is persisted inside Stage B (same transaction as the
				// transfer commit), so a committed key always carries a payload.
				// If it is somehow empty, return a minimal committed marker rather
				// than a fabricated zero-value result.
				var cachedResult ledger.TxResult
				if len(payload) > 0 {
					_ = json.Unmarshal(payload, &cachedResult)
				} else {
					cachedResult = ledger.TxResult{
						ClientID:       params.ClientID,
						IdempotencyKey: params.IdempotencyKey,
						Status:         "committed",
					}
				}
				cachedResult.IsCached = true
				return &cachedResult, &cc.RetryStats{}, nil
			}

			if state == ledger.StatePending {
				// Request currently in-flight in Stage B -> reject with 429/409 processing error
				return nil, nil, ledger.ErrProcessingRetryLater
			}

			// state == 'failed': allow re-execution by atomically claiming the key
			// (failed -> pending). The CAS guard ensures exactly one retryer wins;
			// concurrent retryers lose the race and are told to retry later.
			res, claimErr := m.db.ExecContext(ctx, `
				UPDATE idempotency_keys
				SET state = 'pending', updated_at = CURRENT_TIMESTAMP
				WHERE client_id = $1 AND idempotency_key = $2 AND state = 'failed'
			`, params.ClientID, params.IdempotencyKey)
			if claimErr != nil {
				// Transient contention on the key row (lock timeout / deadlock)
				// is a retry-later condition, not a client-visible failure: the
				// claim is atomic, so no state changed and no effect exists.
				// Permanent errors keep their identity and fail loudly.
				if retriable, _ := cc.IsRetriableError(claimErr); retriable {
					return nil, nil, ledger.ErrProcessingRetryLater
				}
				return nil, nil, fmt.Errorf("failed to reclaim failed idempotency key: %w", claimErr)
			}
			if rows, rowsErr := res.RowsAffected(); rowsErr != nil || rows == 0 {
				return nil, nil, ledger.ErrProcessingRetryLater
			}
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

	// Business rejection (ADR-07/ADR-17 revision): the retry controller reports
	// insufficient funds as a successful result (nil error), so without this
	// transition the key would linger in 'pending' — blocking duplicates with
	// 429s until TTL expiry. No ledger effect exists, so the key moves to the
	// terminal 'failed' state and later retries may re-claim it.
	if result != nil && result.Status == "insufficient_funds" {
		_, _ = m.db.ExecContext(ctx, `
			UPDATE idempotency_keys 
			SET state = 'failed', updated_at = CURRENT_TIMESTAMP 
			WHERE client_id = $1 AND idempotency_key = $2 AND state = 'pending'
		`, params.ClientID, params.IdempotencyKey)
	}

	// The cached response payload is persisted inside the Stage B transaction
	// itself (see ledger.RecordEntriesAndCommit), so key state, ledger entries,
	// and response payload all flip atomically at commit.
	return result, stats, nil
}
