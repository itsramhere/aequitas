package idempotency

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type TTLCleaner struct {
	db               *sql.DB
	pendingTTL       time.Duration
	statementTimeout time.Duration
	interval         time.Duration
}

func NewTTLCleaner(db *sql.DB, pendingTTL, statementTimeout, interval time.Duration) (*TTLCleaner, error) {
	if pendingTTL <= statementTimeout {
		return nil, fmt.Errorf("safety constraint violation: pendingTTL (%v) must be strictly greater than statementTimeout (%v)", pendingTTL, statementTimeout)
	}
	if interval <= 0 {
		interval = 5 * time.Second
	}
	return &TTLCleaner{
		db:               db,
		pendingTTL:       pendingTTL,
		statementTimeout: statementTimeout,
		interval:         interval,
	}, nil
}

func (c *TTLCleaner) Start(ctx context.Context) {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.cleanStalePendingKeys(ctx)
		}
	}
}

func (c *TTLCleaner) cleanStalePendingKeys(ctx context.Context) {
	cutoff := time.Now().Add(-c.pendingTTL)
	// Transition expired pending keys to the terminal 'failed' state instead of
	// deleting them. A hard DELETE would allow a duplicate request to insert a
	// fresh 'pending' row that a still-in-flight Stage B could then match,
	// double-applying the transfer. With a state transition, any late Stage B
	// commit attempt matches 0 rows and aborts, preserving at-most-once.
	res, err := c.db.ExecContext(ctx, `
		UPDATE idempotency_keys
		SET state = 'failed', updated_at = CURRENT_TIMESTAMP
		WHERE state = 'pending' AND GREATEST(created_at, updated_at) < $1
	`, cutoff)
	if err != nil {
		log.Printf("[TTL Cleaner Error] Failed to expire stale pending keys: %v", err)
		return
	}
	rows, err := res.RowsAffected()
	if err == nil && rows > 0 {
		log.Printf("[TTL Cleaner] Transitioned %d stale pending idempotency key(s) to 'failed'", rows)
	}
}
