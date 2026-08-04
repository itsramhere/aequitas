package workload

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type DBCleaner struct {
	db *sql.DB
}

func NewDBCleaner(db *sql.DB) *DBCleaner {
	return &DBCleaner{db: db}
}

// DisableAutovacuum suppresses autovacuum on ledger tables during the measured cell execution window
func (c *DBCleaner) DisableAutovacuum(ctx context.Context) error {
	tables := []string{"accounts", "entries", "transactions", "idempotency_keys"}
	for _, table := range tables {
		query := fmt.Sprintf("ALTER TABLE %s SET (autovacuum_enabled = false);", table)
		if _, err := c.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to disable autovacuum on %s: %w", table, err)
		}
	}
	log.Println("[DBCleaner] Autovacuum successfully disabled for all ledger tables.")
	return nil
}

// ExplicitVacuum runs a full VACUUM ANALYZE between cells to clean dead tuples and reset storage state
func (c *DBCleaner) ExplicitVacuum(ctx context.Context) error {
	log.Println("[DBCleaner] Executing inter-cell VACUUM ANALYZE...")
	tables := []string{"accounts", "entries", "transactions", "idempotency_keys"}
	for _, table := range tables {
		query := fmt.Sprintf("VACUUM ANALYZE %s;", table)
		if _, err := c.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to run VACUUM ANALYZE on %s: %w", table, err)
		}
	}
	log.Println("[DBCleaner] Inter-cell VACUUM ANALYZE complete.")
	return nil
}

// ToggleUnloggedVariant sets tables to UNLOGGED (skipping WAL fsync) or LOGGED
func (c *DBCleaner) ToggleUnloggedVariant(ctx context.Context, unlogged bool) error {
	mode := "LOGGED"
	if unlogged {
		mode = "UNLOGGED"
	}
	tables := []string{"accounts", "entries", "transactions", "idempotency_keys"}
	for _, table := range tables {
		query := fmt.Sprintf("ALTER TABLE %s SET %s;", table, mode)
		if _, err := c.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to set %s table mode on %s: %w", mode, table, err)
		}
	}
	log.Printf("[DBCleaner] Storage mode set to %s for all tables.", mode)
	return nil
}
