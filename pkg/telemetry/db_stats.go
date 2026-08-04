package telemetry

import (
	"context"
	"database/sql"
	"fmt"
)

type DBStatsCollector struct {
	db *sql.DB
}

func NewDBStatsCollector(db *sql.DB) *DBStatsCollector {
	return &DBStatsCollector{db: db}
}

func (c *DBStatsCollector) GetDeadTupleCount(ctx context.Context) (int64, error) {
	var totalDeadTuples int64
	err := c.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(n_dead_tup), 0) 
		FROM pg_stat_user_tables 
		WHERE relname IN ('accounts', 'entries', 'transactions', 'idempotency_keys')
	`).Scan(&totalDeadTuples)
	if err != nil {
		return 0, fmt.Errorf("failed to query pg_stat_user_tables: %w", err)
	}
	return totalDeadTuples, nil
}
