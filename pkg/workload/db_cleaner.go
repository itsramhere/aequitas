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

// RestoreAutovacuum re-enables default autovacuum on ledger tables after the
// benchmark finishes, so suppression never leaks into subsequent runs
func (c *DBCleaner) RestoreAutovacuum(ctx context.Context) error {
	tables := []string{"accounts", "entries", "transactions", "idempotency_keys"}
	for _, table := range tables {
		query := fmt.Sprintf("ALTER TABLE %s SET (autovacuum_enabled = true);", table)
		if _, err := c.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to restore autovacuum on %s: %w", table, err)
		}
	}
	log.Println("[DBCleaner] Autovacuum restored for all ledger tables.")
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

// ToggleUnloggedVariant sets tables to UNLOGGED (skipping WAL fsync) or LOGGED.
//
// ADR-12 revision: PostgreSQL's ALTER TABLE ... SET LOGGED/UNLOGGED refuses to
// change the persistence of any table that participates in a foreign key in
// either direction (a permanent table cannot reference an unlogged one and
// vice versa). The ledger schema links entries -> transactions and
// entries -> accounts, so no per-table toggle order can ever succeed. Rather
// than failing mid-cell with a cryptic server error, the request is rejected
// up front with an explanation; the diagnostic requires an FK-free schema
// (e.g. dropping the constraints for a dedicated diagnostic run).
func (c *DBCleaner) ToggleUnloggedVariant(ctx context.Context, unlogged bool) error {
	if unlogged {
		var fkCount int
		err := c.db.QueryRowContext(ctx, `
			SELECT COUNT(DISTINCT con.oid)
			FROM pg_constraint con
			JOIN pg_class cls ON cls.oid = con.conrelid OR cls.oid = con.confrelid
			WHERE con.contype = 'f'
			  AND cls.relname IN ('accounts', 'entries', 'transactions', 'idempotency_keys')
		`).Scan(&fkCount)
		if err != nil {
			return fmt.Errorf("failed to inspect foreign keys for UNLOGGED variant: %w", err)
		}
		if fkCount > 0 {
			return fmt.Errorf("UNLOGGED diagnostic variant is unsupported on this schema: %d foreign key constraint(s) touch the ledger tables and PostgreSQL cannot ALTER TABLE ... SET UNLOGGED across FK relationships; use an FK-free schema for this diagnostic", fkCount)
		}
	}

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
