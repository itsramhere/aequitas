# Architecture & Design Decisions Record (ADR)

## Project: Concurrency Control and Idempotency Tradeoffs in Skewed-Access Financial Ledgers

This document records all key architectural and design decisions, their rationale, trade-offs, and invariants for the double-entry ledger system.

---

### ADR-01: Double-Entry Invariants & Non-Negative Balance Check at Database Level

* **Decision**: Enforce double-entry debit/credit balance parity and a strict non-negative balance invariant (`balance >= 0`) at the database level, preceded by an application-level sufficiency read check in step 1 of the canonical transaction body.
* **Rationale**: 
  - Without an explicit read-check-then-write sequence for funds sufficiency, a transfer reduces to a blind atomic `UPDATE accounts SET balance = balance + delta`, which Postgres executes atomically even under `READ COMMITTED` isolation without any read-set contention.
  - The funds check creates the load-bearing read set (`SELECT balance`) that every Concurrency Control (CC) strategy must validate and protect.
  - Enforcing invariants in DB constraints prevents illegal state transitions even under edge-case concurrency failures.

---

### ADR-02: Pluggable Concurrency Control Strategies & Database Isolation Level Pinning

* **Decision**: Implement three distinct, pluggable CC strategies with pinned PostgreSQL transaction isolation levels:
  1. **SSI (Serializable Snapshot Isolation)**: Pinned to `SERIALIZABLE`. Plain `SELECT` and `UPDATE` statements; PostgreSQL native MVCC SIREAD locks handle conflict detection.
  2. **Pessimistic Locking**: Pinned to `READ COMMITTED`. Explicit row locking via `SELECT ... FOR UPDATE` under strict total lock ordering.
  3. **OCC (Optimistic Concurrency Control)**: Pinned to `READ COMMITTED`. Version-stamped accounts with single-statement Compare-And-Swap (CAS) update.
* **Rationale**: 
  - Pinning isolation levels to the minimum required for each strategy prevents double-enforcement confounds (e.g., running pessimistic locking inside `SERIALIZABLE` isolation would double-enforce ordering via both Postgres SSI conflict detection and application locks, contaminating benchmark data).

---

### ADR-03: Deterministic Total Lock Ordering for Pessimistic CC

* **Decision**: Enforce a global total lock acquisition order across all lockable resources in the pessimistic strategy:
  1. Idempotency Key Row (acquired in Stage A, separate transaction).
  2. Account Rows in strictly ascending numerical `account_id` order via `SELECT ... FOR UPDATE`.
  3. Ledger Entry Log Inserts (which take implicit `FOR NO KEY UPDATE` locks on referenced accounts via Foreign Key constraints).
  * In addition, Postgres `lock_timeout` is explicitly configured.
* **Rationale**:
  - Sorting account IDs alone does not guarantee deadlock freedom when foreign key implicit locks on entry rows interleave with row locks. Using `FOR NO KEY UPDATE` for entry log foreign keys avoids self-conflicts with held `FOR UPDATE` account locks.
  - Explicit global ordering guarantees deadlock-freedom by construction. Setting a finite `lock_timeout` ensures fast-failure and measurable metrics rather than indefinite blocking.

---

### ADR-04: Single-Statement Compare-And-Swap (CAS) for OCC

* **Decision**: Implement OCC writes as an atomic single-statement update:
  ```sql
  UPDATE accounts 
  SET balance = ?, version = version + 1 
  WHERE id = ? AND version = ?;
  ```
  If zero rows are updated, the version predicate failed; the entire transaction rolls back and retries.
* **Rationale**:
  - Performing a separate read-then-check-then-write loop in application logic reopens the lost-update window.
  - The atomic CAS SQL statement guarantees that version increments and balance checks are atomically committed by Postgres.

---

### ADR-05: Unified Retry Controller across CC Strategies

* **Decision**: Wrap all three CC strategies in a single, shared retry controller with identical maximum attempt count $N$ and exponential backoff schedule with jitter.
* **Rationale**:
  - Each CC strategy fails under a different failure idiom: SSI raises SQLSTATE `40001` (serialization failure), Pessimistic raises `lock_timeout`, and OCC results in 0 rows updated (version mismatch).
  - Without a unified controller, comparing strategy retry rates would compare disparate mechanisms. The shared controller normalizes in-transaction retries vs client-visible failures.

---

### ADR-06: Two-Stage Idempotency Execution Model (Stage A / Stage B)

* **Decision**: Divide request processing into two explicit database transaction stages:
  - **Stage A (Separate, Short Transaction)**: Insert idempotency key into `idempotency_keys` table with state `pending`. Rely on unique index `(client_id, idempotency_key)` to detect duplicate arrivals.
  - **Stage B (Single Atomic Transaction)**: Execute canonical ledger transaction (debit, credit, entry log insert) AND update idempotency key state from `pending` to `committed` inside a single `BEGIN ... COMMIT` block.
* **Rationale**:
  - Eliminates split-brain between idempotency key state and ledger state under process crashes/kills. A process kill during Stage B can only result in either complete commit (key `committed` + ledger entries present) or full rollback (key `pending` + zero ledger entries).

---

### ADR-07: Pending-Collision Policy and Statement-Timeout vs TTL Safety Ordering

* **Decision**: 
  - When a retry arrives while Stage B is still `pending`, return an immediate `429/409 Processing, retry later` response rather than blocking or racing.
  - Pending keys expire after a Configured Time-To-Live (TTL).
  - PostgreSQL `statement_timeout` for Stage B is enforced to be strictly less than the pending key TTL (`statement_timeout < TTL`).
* **Rationale**:
  - Rejecting pending retries avoids introducing secondary locking interactions between idempotency key lookups and underlying CC strategy benchmarking.
  - Enforcing `statement_timeout < TTL` makes it provably impossible for a slow Stage B transaction to outlive its idempotency key. A live transaction will be aborted by Postgres before its key becomes eligible for TTL garbage collection, preventing double-execution races.

---

### ADR-08: Clock-Drift Isolation & Multi-Layer Telemetry Instrumentation

* **Decision**:
  - Never calculate latency by subtracting client timestamps from server timestamps. Client-side metrics (end-to-end latency, rejection backoff) use the client clock only. Server-side metrics (app CC time, DB execution, pool wait, WAL wait proxy) use the server clock only.
  - Server latency is instrumented as distinct non-overlapping or explicitly tagged spans: Application-level CC wait vs PostgreSQL query execution/commit vs Connection Pool wait vs WAL fsync proxy vs Go GC pause correlation.
* **Rationale**:
  - Prevents NTP clock drift from contaminating latency measurements across distributed client-server setups.
  - Decomposing latency allows exact attribution of bottlenecks (e.g. distinguishing application lock acquisition delay from storage engine WAL disk fsync bounds).

---

### ADR-09: Two-Check Independent Auditor for Correctness & Idempotency Verification

* **Decision**: Implement an independent post-run auditor with two distinct checks:
  1. **Balance Reconciliation**: Reconstruct account balances by summing the immutable append-only `entries` log and diffing against the `accounts` table.
  2. **Idempotency Cardinality Audit**: For every committed idempotency key, count the number of distinct entry-sets/transactions created in the `entries` table (must equal exactly 1).
* **Rationale**:
  - Balance reconciliation catches lost writes or corrupted balances, but **cannot detect duplicate execution** of an idempotent transaction (since a duplicated transfer writes a second balanced pair of debit/credit entries, leaving net balances correct).
  - The cardinality check explicitly verifies exactly-once semantic guarantees by catching duplicate entry creation.

---

### ADR-10: Bare CC Path Execution Toggle (Experiment Set 1 Scoping)

* **Decision**: Provide a configuration flag (`--enable-stage-a=false`) in the execution harness to bypass Stage A (idempotency-key insert) during Experiment Set 1 runs, executing only steps 1–4 and 6 of the canonical transaction body.
* **Rationale**:
  - If Experiment Set 1 always traversed Stage A, the idempotency key table's unique index insert would act as a second, strategy-independent write bottleneck on every transaction, adding WAL and locking overhead that confounds pure CC strategy throughput comparisons (RQ1).

---

### ADR-11: Database Cleaning Controls & Autovacuum Suppression

* **Decision**: 
  - Programmatically disable PostgreSQL autovacuum on ledger tables during active 5-minute benchmark execution windows (`ALTER TABLE ... SET (autovacuum_enabled = false)`).
  - Execute an explicit `VACUUM ANALYZE` hook between individual configuration cells in the benchmark harness.
* **Rationale**:
  - Retry-heavy strategies (such as OCC under heavy contention) generate significantly higher dead-tuple volume per committed transaction than low-retry strategies (pessimistic). Background autovacuum triggering during execution windows would introduce uncommanded background disk and CPU noise correlated with specific CC strategies.
  - Running explicit `VACUUM` between cells ensures clean storage baselines and prevents cross-cell performance bleed.

---

### ADR-12: Enhanced Telemetry (Dead-Tuple Rate, Realized Hot-Row Hit Rate & Deadlock Count)

* **Decision**:
  - **Dead-tuple Generation Rate**: Sample `pg_stat_user_tables.n_dead_tup` before and after each cell run to record exact dead tuples generated per committed transaction.
  - **Realized Hot-Row Hit Rate**: Track the exact percentage of total requests that hit the single most contended account ID under the configured Zipfian skew coefficient ($\theta$).
  - **Deadlock & Lock Timeout Metrics**: Instrument the pessimistic locking strategy to catch SQLSTATE `40P01` (deadlock detected) and SQLSTATE `55P03`/`57014` (lock timeout) explicitly, reporting them under a dedicated `deadlock_count` metric rather than generic errors.
* **Rationale**:
  - Dead-tuple rate quantifies the operational GC/bloat overhead of OCC retry loops.
  - Realized hot-row hit rate verifies whether observed throughput plateaus at high skew are caused by single-row physical serialization limits vs hardware limits.
  - Deadlock counting treats lock-wait/order anomalies as explicit experimental findings rather than unhandled failures.

---

### ADR-13: Infrastructure Controls, Sized Connection Pools & UNLOGGED Diagnostic Variant

* **Decision**:
  - **Environment Pinning**: Explicitly record `GOGC` (e.g. `100`) and `GOMAXPROCS` settings in benchmark output JSON metadata.
  - **Connection Sizing**: Size database connection pool (`sql.DB.SetMaxOpenConns(275)`) and PostgreSQL `max_connections` (`300`) strictly above the maximum tested concurrency level (250 clients + headroom).
  - **UNLOGGED Diagnostic Variant**: Include a benchmark harness toggle (`--use-unlogged-tables`) to run a secondary sweep on PostgreSQL `UNLOGGED` tables at heavy skew ($\theta = 1.2$).
* **Rationale**:
  - Sizing connection pools and Postgres limits above client concurrency eliminates pool exhaustion as a binding constraint.
  - Pinning runtime environment variables ensures run reproducibility.
  - The `UNLOGGED` variant skips WAL fsync entirely, serving as a diagnostic to confirm whether a throughput plateau at $\theta = 1.2$ is bound by application CC software logic or disk WAL fsync I/O.
