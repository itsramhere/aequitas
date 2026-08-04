-- PostgreSQL Database Schema for Skewed-Access Double-Entry Ledger
-- Enforces DB-level invariants: balance >= 0, double-entry relations, and idempotency state machine.

CREATE TABLE IF NOT EXISTS accounts (
    id BIGINT PRIMARY KEY,
    balance NUMERIC(18, 4) NOT NULL CHECK (balance >= 0),
    version BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS idempotency_keys (
    client_id VARCHAR(128) NOT NULL,
    idempotency_key VARCHAR(128) NOT NULL,
    state VARCHAR(32) NOT NULL CHECK (state IN ('pending', 'committed', 'failed')),
    response_payload JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (client_id, idempotency_key)
);

CREATE INDEX IF NOT EXISTS idx_idempotency_state_created ON idempotency_keys (state, created_at);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY,
    client_id VARCHAR(128) NOT NULL,
    idempotency_key VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS entries (
    id BIGSERIAL PRIMARY KEY,
    transaction_id UUID NOT NULL REFERENCES transactions(id) ON DELETE CASCADE,
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount NUMERIC(18, 4) NOT NULL CHECK (amount > 0),
    entry_type VARCHAR(16) NOT NULL CHECK (entry_type IN ('debit', 'credit')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_entries_account_id ON entries (account_id);
CREATE INDEX IF NOT EXISTS idx_entries_transaction_id ON entries (transaction_id);

-- Helper procedure to initialize N balanced accounts for testing
-- Example: SELECT seed_accounts(10000, 1000.00);
CREATE OR REPLACE FUNCTION seed_accounts(num_accounts INT, initial_balance NUMERIC)
RETURNS VOID AS $$
BEGIN
    TRUNCATE TABLE entries, transactions, idempotency_keys, accounts RESTART IDENTITY CASCADE;
    INSERT INTO accounts (id, balance, version, updated_at)
    SELECT i, initial_balance, 0, CURRENT_TIMESTAMP
    FROM generate_series(1, num_accounts) AS i;
END;
$$ LANGUAGE plpgsql;
