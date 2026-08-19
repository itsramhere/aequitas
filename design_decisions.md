# Architecture & Design Decisions Record (ADR)
## Project: Concurrency Control and Idempotency Tradeoffs in Skewed-Access Financial Ledgers

This document records all key architectural and design decisions, their rationale, trade-offs, and invariants for the double-entry ledger system.

### ADR-01: Double-Entry Invariants & Non-Negative Balance Check at Database Level
* **Decision**: Enforce double-entry debit/credit balance parity and a strict non-negative balance invariant (`balance >= 0`) at the database level, preceded by an application-level sufficiency read check in step 1 of the canonical transaction body.
* **Rationale**: 
  * Without an explicit read-check-then-write sequence for funds sufficiency, a transfer reduces to a blind atomic `UPDATE accounts SET balance = balance + delta`, which Postgres executes atomically even under `READ COMMITTED` isolation without any read-set contention.
  * The funds check creates the load-bearing read set (`SELECT balance`) that every Concurrency Control (CC) strategy must validate and protect.
  * Enforcing invariants in DB constraints prevents illegal state transitions even under edge-case concurrency failures.

### ADR-02: Pluggable Concurrency Control Strategies & Database Isolation Level Pinning
* **Decision**: Implement four distinct, pluggable CC strategies with pinned PostgreSQL transaction isolation levels:
  1. **SSI (Serializable Snapshot Isolation)**: Pinned to `SERIALIZABLE`. Plain `SELECT` and `UPDATE` statements; PostgreSQL native MVCC SIREAD locks handle conflict detection.
  2. **Pessimistic Locking**: Pinned to `READ COMMITTED`. Explicit row locking via `SELECT ... FOR UPDATE` under strict total lock ordering.
  3. **OCC (Optimistic Concurrency Control)**: Pinned to `READ COMMITTED`. Version-stamped accounts with single-statement Compare-And-Swap (CAS) update.
  4. **Adaptive Hybrid CC (ADAPTIVE)**: Dynamically hot-swaps between OCC and Pessimistic locking at runtime while reflecting active isolation level (`READ COMMITTED`).
* **Rationale**: Pinning isolation levels to the minimum required for each strategy prevents double-enforcement confounds. 
* **Revision Note (Phase 1 Update)**: Expanded from 3 static strategies to include `ADAPTIVE`. Static strategy assignment could not optimize trade-offs when workload contention dynamically shifted. Adding `ADAPTIVE` enables autonomous runtime switching based on live abort feedback.

### ADR-03: Deterministic Total Lock Ordering for Pessimistic CC
* **Decision**: Enforce a global total lock acquisition order across all lockable resources in the pessimistic strategy:
  1. Idempotency Key Row (acquired in Stage A, separate transaction).
  2. Account Rows in strictly ascending numerical `account_id` order via `SELECT ... FOR UPDATE`.
  3. Ledger Entry Log Inserts (which take implicit `FOR NO KEY UPDATE` locks on referenced accounts via Foreign Key constraints).
  * In addition, Postgres `lock_timeout` is explicitly configured.
* **Rationale**: Sorting account IDs alone does not guarantee deadlock freedom when foreign key implicit locks on entry rows interleave with row locks. Using `FOR NO KEY UPDATE` for entry log foreign keys avoids self-conflicts with held `FOR UPDATE` account locks. Explicit global ordering guarantees deadlock-freedom by construction. Setting a finite `lock_timeout` ensures fast-failure and measurable metrics rather than indefinite blocking.

### ADR-04: Single-Statement Compare-And-Swap (CAS) for OCC
* **Decision**: Implement OCC writes as an atomic single-statement update:
  ```sql
  UPDATE accounts 
  SET balance = ?, version = version + 1 
  WHERE id = ? AND version = ?;
  ```

If zero rows are updated, the version predicate failed; the entire transaction rolls back and retries.

* **Rationale**: Performing a separate read-then-check-then-write loop in application logic reopens the lost-update window. The atomic CAS SQL statement guarantees that version increments and balance checks are atomically committed by Postgres.

### ADR-05: Unified Retry Controller across CC Strategies
**Decision**: Wrap all CC strategies (including `ADAPTIVE`) in a single, shared retry controller with identical maximum attempt count N and exponential backoff schedule with jitter.

**Rationale**: Each CC strategy fails under a different failure idiom: SSI raises SQLSTATE 40001 (serialization failure), Pessimistic raises lock_timeout, and OCC results in 0 rows updated (version mismatch). Without a unified controller, comparing strategy retry rates would compare disparate mechanisms. The shared controller normalizes in-transaction retries vs client-visible failures.

**Revision Note (Phase 1 Update)**: Integrated `AdaptiveStrategy` with `UnifiedRetryController`. In-flight retries identified during `ExecuteWithRetry` loops atomically increment the sliding window `interval_retries` counter, feeding real-time abort signals directly back to the adaptive controller.

### ADR-06: Two-Stage Idempotency Execution Model (Stage A / Stage B)
**Decision**: Divide request processing into two explicit database transaction stages:

Stage A (Separate, Short Transaction): Insert idempotency key into idempotency_keys table with state pending. Rely on unique index (client_id, idempotency_key) to detect duplicate arrivals.

Stage B (Single Atomic Transaction): Execute canonical ledger transaction (debit, credit, entry log insert) AND update idempotency key state from pending to committed inside a single BEGIN ... COMMIT block.

**Rationale**: Eliminates split-brain between idempotency key state and ledger state under process crashes/kills. A process kill during Stage B can only result in either complete commit (key committed + ledger entries present) or full rollback (key pending + zero ledger entries).

### ADR-07: Pending-Collision Policy and Statement-Timeout vs TTL Safety Ordering
**Decision**:
When a retry arrives while Stage B is still pending, return an immediate 429/409 Processing, retry later response rather than blocking or racing.

Pending keys expire after a Configured Time-To-Live (TTL).

PostgreSQL statement_timeout for Stage B is enforced to be strictly less than the pending key TTL (statement_timeout < TTL).

**Rationale**: Rejecting pending retries avoids introducing secondary locking interactions between idempotency key lookups and underlying CC strategy benchmarking. Enforcing statement_timeout < TTL makes it provably impossible for a slow Stage B transaction to outlive its idempotency key. A live transaction will be aborted by Postgres before its key becomes eligible for TTL garbage collection, preventing double-execution races.

### ADR-08: Clock-Drift Isolation & Multi-Layer Telemetry Instrumentation
**Decision**: Client-side metrics use client clocks only; server-side metrics use server clocks only. In addition, capture per-second transient time-series data (`TimeSeriesPoint`) during the measured run window alongside overall cell percentiles.

**Rationale**: Clock separation prevents NTP drift from corrupting latency data. Per-second time-series logging enables fine-grained analysis of transient contention spikes, abort storms, and adaptive hot-swapping stabilization over time.

**Revision Note (Phase 2 Update)**: Added per-second time-series telemetry (`second`, `tps`, `retries`). Static aggregate cell reports (`p50`, `p95`, `p99`) obscured transient time-dependent phenomena; per-second logging provides fine-grained visibility into hot-swapping stabilization dynamics.

### ADR-09: Two-Check Independent Auditor for Correctness & Idempotency Verification
**Decision**: Implement an independent post-run auditor with two distinct checks:

**Balance Reconciliation**: Reconstruct account balances by summing the immutable append-only entries log and diffing against the accounts table.

**Idempotency Cardinality Audit**: For every committed idempotency key, count the number of distinct entry-sets/transactions created in the entries table (must equal exactly 1).

**Rationale**: Balance reconciliation catches lost writes or corrupted balances, but cannot detect duplicate execution of an idempotent transaction (since a duplicated transfer writes a second balanced pair of debit/credit entries, leaving net balances correct). The cardinality check explicitly verifies exactly-once semantic guarantees.

### ADR-10: Bare CC Path Execution Toggle (Experiment Set 1 Scoping)
**Decision**: Provide a configuration flag (--enable-stage-a=false) in the execution harness to bypass Stage A.

**Rationale**: Prevents Stage A unique index inserts from acting as a baseline bottleneck during pure CC strategy throughput comparisons.

### ADR-11: Database Cleaning Controls & Autovacuum Suppression
**Decision**: Programmatically disable PostgreSQL autovacuum on ledger tables during active benchmark execution windows, executing an explicit VACUUM ANALYZE hook between configuration cells.

**Rationale**: Retry-heavy strategies (OCC) generate significantly higher dead-tuple volumes. Background autovacuum triggering during execution windows would introduce uncommanded background disk and CPU I/O noise correlated with specific CC strategies.

### ADR-12: Infrastructure Controls, Sized Connection Pools & UNLOGGED Diagnostic Variant
**Decision**: Size database connection pool strictly above the maximum tested concurrency level (250 clients + headroom) and include a benchmark harness toggle to run a secondary sweep on PostgreSQL UNLOGGED tables.

**Rationale**: Eliminates connection pool exhaustion as a binding constraint. The UNLOGGED variant serves as a diagnostic to confirm whether a throughput plateau is bound by logical application CC software or physical disk WAL fsync I/O.

### ADR-13: Lexicographical Account ID Sorting for Optimistic Strategies (OCC/SSI)
Decision: Prior to issuing UPDATE statements, the application logic enforces lexicographical sorting (ascending order) of Account IDs for both OCC and SSI strategies, despite their optimistic nature.

**Rationale**: While OCC and SSI do not utilize application-level row locks, PostgreSQL's MVCC storage engine still acquires implicit row-level exclusive locks during the physical UPDATE execution phase. Updating accounts in debit/credit order under high contention causes true AB-BA physical lock deadlocks (SQLSTATE 40P01) at the storage layer, artificially inflating abort rates. Sorting IDs completely eliminates these structural deadlocks.

### ADR-14: Two-Phase Latency Instrumentation (CC Wait vs. Execution)
**Decision**: The instrumentation harness explicitly isolates concurrency control lock acquisition wait time (cc_wait_latency) from database logical execution time (db_latency).

**Rationale**: Logging the full Go database driver round-trip as "database execution" masks the true cost of pessimistic locking. By splitting SELECT ... FOR UPDATE (Wait) from the subsequent UPDATE and COMMIT (Execution), the telemetry accurately attributes high latencies to lock-queue depth rather than storage engine slowness.

### ADR-15: Cumulative Dead Tuple Accounting (HOT Pruning Bypass)
**Decision**: Database bloat (dead tuples generated) is calculated by capturing the delta of n_tup_upd + n_tup_del from pg_stat_user_tables before and after each benchmark cell, rather than sampling n_dead_tup.

**Rationale**: PostgreSQL's internal Heap-Only Tuples (HOT) pruning opportunistically cleans dead tuples during regular scans, even with Autovacuum disabled. Relying on the n_dead_tup gauge at the end of a run heavily undercounts true bloat generation. Tracking cumulative updates/deletes ensures mathematically exact bloat measurements.

### ADR-16: Strict Error Chain Traversal for Anomaly Detection
**Decision**: The unified retry controller utilizes errors.As() to traverse the entire error chain when evaluating retriable database failures, rather than relying on strict interface type assertions (err.(*pq.Error)).

**Rationale**: Serialization anomalies (SSI SQLSTATE 40001) are often raised by PostgreSQL during the COMMIT phase. When the application logic wraps these failures via fmt.Errorf("...: %w"), strict type assertions fail, causing the retry controller to misclassify recoverable CC conflicts as fatal business errors. Error chain traversal guarantees accurate anomaly interception.

### ADR-17: Segregation of Business Invariant Rejections
**Decision**: Application-level business logic failures (specifically ErrInsufficientFunds and ErrProcessingRetryLater) are explicitly intercepted, counted in dedicated telemetry metrics (insufficient_funds_count and collision_rejections), and strictly excluded from Concurrency Control failure rates.

**Rationale**: High-throughput strategies hit the mathematical zero-bound of account balances faster than slower strategies. Bounding CC abort rates requires filtering out valid business-rule rejections; otherwise, the fastest strategies are artificially penalized with inflated "failure" metrics.

### ADR-18: Adaptive Hybrid Concurrency Control (Sliding Window & Hot-Swapping)
**Decision**: Implement an adaptive hybrid CC strategy (`AdaptiveStrategy`) that encapsulates both `OCCStrategy` and `PessimisticStrategy`. The strategy routes requests to OCC by default and monitors execution over a 500ms sliding window using atomic request and retry counters (`interval_requests`, `interval_retries`). If the measured abort ratio R_abort = retries / requests exceeds 0.15, the controller safely hot-swaps active routing to Pessimistic locking under a `sync.RWMutex`. If R_abort drops below 0.05, active routing hot-swaps back to OCC. The strategy type reports `StrategyAdaptive`, while `IsolationLevel()` dynamically reflects the currently active underlying strategy.

**Rationale**: Low-skew/low-contention workloads incur unnecessary lock acquisition wait overhead under Pessimistic locking, whereas high-skew workloads experience severe abort storms under pure OCC. Dynamic threshold-based switching optimizes throughput and latency across shifting contention profiles without requiring application restarts or offline reconfiguration.

### ADR-19: Multi-Repetition Statistical Benchmark Harness
**Decision**: Automate empirical data collection across N=5 repetitions per experimental configuration cell (4 strategies x 6 Zipfian skews x 5 repetitions = 120 total cell runs) via `run_repetition_sweep.ps1`, persisting raw JSON metrics to `.\results\set1_ci\`.

**Rationale**: Single-execution benchmark samples are vulnerable to transient OS scheduling jitter, background disk I/O spikes, and Go GC pause variations. Running 5 repetitions per configuration cell yields statistically robust datasets for mean, median, and confidence interval calculations.

### ADR-20: Formal Specification of Two-Stage Idempotency under Process Crash Failures
**Decision**: Formally model the two-stage idempotency execution protocol in TLA+ (`Idempotency.tla`), capturing Stage A pending key registration, Stage B atomic transfer and state update, arbitrary worker process crashes, and TTL pending key cleanup.

**Rationale**: Empirical testing alone cannot prove safety across all possible interleavings of process crashes and retries. Formal model checking rigorously verifies the `Safety` invariant (at-most-once entry set creation per idempotency key) and `Liveness` invariant (eventual resolution of all pending keys to terminal committed or failed states) under arbitrary crash timings.