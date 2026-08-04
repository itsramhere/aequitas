package cc

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type RetryStats struct {
	TotalAttempts         int
	InternalRetries       int
	DeadlockCount         int
	ClientVisibleFailures int
}

type UnifiedRetryController struct {
	MaxAttempts int
	BaseBackoff time.Duration
	MaxBackoff  time.Duration
}

func NewUnifiedRetryController(maxAttempts int, baseBackoff, maxBackoff time.Duration) *UnifiedRetryController {
	if maxAttempts <= 0 {
		maxAttempts = 5
	}
	if baseBackoff <= 0 {
		baseBackoff = 5 * time.Millisecond
	}
	if maxBackoff <= 0 {
		maxBackoff = 100 * time.Millisecond
	}
	return &UnifiedRetryController{
		MaxAttempts: maxAttempts,
		BaseBackoff: baseBackoff,
		MaxBackoff:  maxBackoff,
	}
}

// IsRetriableError checks if the error is a retriable CC conflict error
func IsRetriableError(err error) (bool, bool) { // (isRetriable, isDeadlockOrTimeout)
	if err == nil {
		return false, false
	}
	if errors.Is(err, ledger.ErrVersionMismatch) {
		return true, false
	}
	if pqErr, ok := err.(*pq.Error); ok {
		// 40001 = serialization_failure (SSI)
		if pqErr.Code == "40001" {
			return true, false
		}
		// 40P01 = deadlock_detected
		// 55P03 = lock_not_available
		// 57014 = statement_timeout (lock timeout)
		if pqErr.Code == "40P01" || pqErr.Code == "55P03" || pqErr.Code == "57014" {
			return true, true
		}
	}
	return false, false
}

func (c *UnifiedRetryController) ExecuteWithRetry(
	ctx context.Context,
	db *sql.DB,
	strategy Strategy,
	params ledger.TransferParams,
	opts TransferOptions,
) (*ledger.TxResult, *RetryStats, error) {
	stats := &RetryStats{}
	var lastErr error

	for attempt := 1; attempt <= c.MaxAttempts; attempt++ {
		stats.TotalAttempts = attempt
		result, err := strategy.ExecuteTransfer(ctx, db, params, opts)
		if err == nil {
			result.Attempts = attempt
			stats.InternalRetries = attempt - 1
			return result, stats, nil
		}

		lastErr = err
		retriable, isDeadlock := IsRetriableError(err)
		if isDeadlock {
			stats.DeadlockCount++
		}

		if !retriable {
			// Non-retriable business logic error (insufficient funds, account not found)
			return nil, stats, err
		}

		if attempt < c.MaxAttempts {
			// Compute exponential backoff with jitter
			backoff := c.BaseBackoff * time.Duration(1<<(attempt-1))
			if backoff > c.MaxBackoff {
				backoff = c.MaxBackoff
			}
			jitter := time.Duration(rand.Int63n(int64(backoff)/2 + 1))
			sleepTime := backoff/2 + jitter

			select {
			case <-ctx.Done():
				return nil, stats, ctx.Err()
			case <-time.After(sleepTime):
			}
		}
	}

	stats.ClientVisibleFailures = 1
	stats.InternalRetries = c.MaxAttempts - 1
	return nil, stats, lastErr
}
