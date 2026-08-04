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
	res, err := c.db.ExecContext(ctx, `
		DELETE FROM idempotency_keys
		WHERE state = 'pending' AND created_at < $1
	`, cutoff)
	if err != nil {
		log.Printf("[TTL Cleaner Error] Failed to delete stale pending keys: %v", err)
		return
	}
	rows, err := res.RowsAffected()
	if err == nil && rows > 0 {
		log.Printf("[TTL Cleaner] Cleaned %d stale pending idempotency key(s)", rows)
	}
}
