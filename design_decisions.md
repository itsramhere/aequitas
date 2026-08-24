# Architecture & Design Decisions Record (ADR)
## Project: Concurrency Control and Idempotency Tradeoffs in Skewed-Access Financial Ledgers

This document records all key architectural and design decisions, their rationale, trade-offs, and invariants for the double-entry ledger system.

### ADR-01: Double-Entry Invariants & Non-Negative Balance Check at Database Level
* **Decision**: Enforce double-entry debit/credit balance parity and a strict non-negative balance invariant (`balance >= 0`) at the database level, preceded by an application-level sufficiency read check in step 1 of the canonical transaction body.
* **Rationale**: 
  * Without an explicit read-check-then-write sequence for funds sufficiency, a transfer reduces to a blind atomic `UPDATE accounts SET balance = balance + delta`, which Postgres executes atomically even under `READ COMMITTED` isolation without any read-set contention.
  * The funds check creates the load-bearing read set (`SELECT balance`) that every Concurrency Control (CC) strategy must validate and protect.
  * Enforcing invariants in DB constraints prevents illegal state transitions even under edge-case concurrency failures.
* **Revision Note (Money Representation Caveat)**: Balances and amounts are stored as `NUMERIC(18,4)` but carried through the Go application as `float64`. This is acceptable for benchmark-scale amounts only because the workload now quantizes amounts to whole cents (see ADR-21 revision) so both sides of the auditor's reconciliation round identically. Production use requires integer minor units end-to-end; recorded here so the limitation is explicit rather than incidental.

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
* **Revision Note (Terminology Correction)**: The original text mis-stated the lock mode taken by foreign key checks: PostgreSQL FK validation acquires `FOR KEY SHARE` on the referenced row, not `FOR NO KEY UPDATE`. The deadlock-freedom conclusion is unaffected — the FK's `FOR KEY SHARE` locks land on account rows already locked by the *same transaction*, and all cross-transaction account access still follows the strict ascending order — but the documented mechanism was wrong and has been corrected here rather than silently rewritten.

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

**Revision Note (Correctness Fix)**: The in-flight retry feedback described above is now actually implemented: the retry controller exposes an `OnRetriableAttempt` callback invoked once per retriable attempt failure (even when a subsequent retry succeeds), and both `ledger_server` and `workload_gen` wire it to `AdaptiveStrategy.NoteRetry`. Previously the adaptive controller only observed errors returned after retry exhaustion, i.e. client-visible failures, so it would never switch under contention whose retries mostly succeeded.

**Revision Note (PRNG Isolation)**: The controller's backoff jitter was the last remaining caller of the global `math/rand` functions — the exact shared-mutex contention source ADR-21 eliminated in the workload generator. `UnifiedRetryController` now owns a dedicated `rand.Rand` guarded by a mutex (`jitterSleep`), so no benchmark hot path touches the global PRNG.

### ADR-06: Two-Stage Idempotency Execution Model (Stage A / Stage B)
**Decision**: Divide request processing into two explicit database transaction stages:

Stage A (Separate, Short Transaction): Insert idempotency key into idempotency_keys table with state pending. Rely on unique index (client_id, idempotency_key) to detect duplicate arrivals.

Stage B (Single Atomic Transaction): Execute canonical ledger transaction (debit, credit, entry log insert) AND update idempotency key state from pending to committed inside a single BEGIN ... COMMIT block.

**Rationale**: Eliminates split-brain between idempotency key state and ledger state under process crashes/kills. A process kill during Stage B can only result in either complete commit (key committed + ledger entries present) or full rollback (key pending + zero ledger entries).

**Revision Note (Correctness Fix)**: The serialized response payload (`response_payload` column) is now written inside the Stage B transaction, in the same statement that flips the key `pending -> committed`. Key state, ledger entries, and cached response therefore become visible atomically at commit; a duplicate arriving after commit can never observe a committed key with a missing/NULL payload (previously a post-commit UPDATE failure would yield a fabricated zero-value cached result). The duplicate read path additionally returns a minimal committed marker instead of a zero-value `TxResult` when the payload is unexpectedly empty.

**Revision Note (Duplicate Injection & Observability)**: The benchmark harness previously minted a fresh UUID key for every request, which made the entire duplicate-handling machinery (unique-index collision path, cached-response reads, 429 pending collisions, failed-key reclaim) unreachable by construction — the auditor's cardinality checks were therefore vacuously true. The workload generator now accepts `-duplicate-rate p`: with probability `p` a worker replays one of its *own* recent keys (worker-local ring buffer; client IDs are also worker-local, preserving the `(client_id, idempotency_key)` uniqueness domain). Duplicate requests served from cache are counted as a new first-class metric `cached_hits` and are excluded from `committed_txns`, TPS, and all latency percentile arrays — cached responses carry zero-valued durations by construction (the persisted payload contains IDs/status only), so mixing them into latency samples fabricated sub-microsecond percentiles.

### ADR-07: Pending-Collision Policy and Statement-Timeout vs TTL Safety Ordering
**Decision**:
When a retry arrives while Stage B is still pending, return an immediate 429/409 Processing, retry later response rather than blocking or racing.

Pending keys expire after a Configured Time-To-Live (TTL).

PostgreSQL statement_timeout for Stage B is enforced to be strictly less than the pending key TTL (statement_timeout < TTL).

**Rationale**: Rejecting pending retries avoids introducing secondary locking interactions between idempotency key lookups and underlying CC strategy benchmarking. Enforcing statement_timeout < TTL makes it provably impossible for a slow Stage B transaction to outlive its idempotency key. A live transaction will be aborted by Postgres before its key becomes eligible for TTL garbage collection, preventing double-execution races.

**Revision Note (Correctness Fix)**: The statement_timeout < TTL invariant is now enforced at runtime, not just documented: every CC strategy applies `SET LOCAL statement_timeout` per transaction (1500ms via `TransferOptions.StatementTimeout`), `ledger_server` fails fast unless `MaxAttempts × statement_timeout < pendingTTL` (5 × 1.5s = 7.5s < 10s), and `NewTTLCleaner` validates `pendingTTL > statementTimeout` on construction. Additionally, the TTL cleaner now **transitions** expired pending keys to the terminal `failed` state instead of hard-DELETEing the row. The DELETE variant had a double-apply race: a deleted in-flight key allowed a duplicate to insert a fresh `pending` row that the original (slow) Stage B commit could still match, executing the transfer twice. With a state transition, any late Stage B commit matches zero rows and aborts, preserving at-most-once by construction.

**Revision Note (Failed-Key Reclaim)**: Retries against a `failed` key are now re-executable: the manager atomically re-claims the key with a CAS-guarded `UPDATE ... WHERE state = 'failed'` (exactly one retryer wins; concurrent retryers receive the 429 collision response), then proceeds to Stage B. Previously the fall-through path never reset the state, so Stage B's `state = 'pending'` predicate always matched zero rows and every retry of a failed key aborted.

**Revision Note (Harness-Wide Enforcement)**: The runtime enforcement described above existed only in the `ledger_server` binary; the actual experiment harness (`workload_gen`) never set `TransferOptions.StatementTimeout` and never started a TTL cleaner, so no recorded benchmark run was actually covered by this invariant. Both are now wired into the harness: `workload_gen` sets `StatementTimeout` on every cell, starts a `TTLCleaner` whenever Stage A is enabled, and fail-fasts on `MaxAttempts × statement_timeout < pendingTTL` checked against the retry controller's *actual* configuration (previously the guard compared duplicated hardcoded constants, i.e. it could never fire). The original "provably impossible" wording is also corrected: `statement_timeout` bounds each statement, not the transaction, and the inequality does not account for time spent before Stage B begins (queueing, backoff, GC pauses). The invariant is a strong engineering margin, not a proof — the residual risk is covered by the TTL-cleaner state transition and by the TLA+ model.

**Revision Note (Insufficient-Funds Key Resolution)**: The unified retry controller reports business rejections as successful results (`status = insufficient_funds`, nil error), so `IdempotencyManager.ProcessTransfer` previously skipped its failure path and left the key `pending` indefinitely in any process without a TTL cleaner. The manager now transitions such keys to terminal `failed` immediately after Stage B returns that status: no ledger effect exists, duplicates stop receiving spurious 429s once the CAS reclaim runs, and at-most-once is unaffected because nothing was ever committed for the key.

**Revision Note (Reclaim Contention Classification)**: A transient failure of the failed-key reclaim `UPDATE` itself (lock timeout / deadlock on a contended key row) previously surfaced as a wrapped hard error and was counted as a client-visible failure. The claim statement is atomic — when it fails, the key remains `failed` and no effect exists — so such transients now map to `ErrProcessingRetryLater`, consistent with the `rows = 0` loser path directly beside it; permanent errors keep their identity and still fail loudly.

### ADR-08: Clock-Drift Isolation & Multi-Layer Telemetry Instrumentation
**Decision**: Client-side metrics use client clocks only; server-side metrics use server clocks only. In addition, capture per-second transient time-series data (`TimeSeriesPoint`) during the measured run window alongside overall cell percentiles.

**Rationale**: Clock separation prevents NTP drift from corrupting latency data. Per-second time-series logging enables fine-grained analysis of transient contention spikes, abort storms, and adaptive hot-swapping stabilization over time.

**Revision Note (Phase 2 Update)**: Added per-second time-series telemetry (`second`, `tps`, `retries`). Static aggregate cell reports (`p50`, `p95`, `p99`) obscured transient time-dependent phenomena; per-second logging provides fine-grained visibility into hot-swapping stabilization dynamics.

**Revision Note (Latency Population Policy)**: Latency percentile arrays are now populated exclusively by fresh committed executions. Cached duplicate responses (zero-valued durations by construction) and business rejections (no Stage B entry) no longer fabricate sub-microsecond samples. The previously dead `client_backoff_rejection_latency` metric — declared in every report but never written by any code path — is now measured as the client-observed end-to-end duration of `429` pending-collision rejections. Percentile computation was also corrected to the standard nearest-rank order statistic (`ceil(p*n)-1`); the previous `floor(p*n/100)` indexing was systematically off by one rank.

**Revision Note (Phase-Boundary Attribution Fix)**: The warm-up/measured split previously used a shared boolean flag plus sleeps, so requests *started* during warm-up could complete inside the measured window (inflating it) and in-flight requests cut off at window close vanished from all accounting. Phase membership is now decided by each request's start time against fixed deadlines (`warmupEnd`, `measuredEnd`): workers self-retire at `measuredEnd`, nothing drains into or out of the measured bucket, and TPS = committed / nominal window holds by construction.

**Revision Note (Time-Series Edge Semantics)**: The per-second sampler opens exactly at `warmupEnd` and drops any tick landing at or after `measuredEnd`, so the final second is usually absent from `time_series`. Recording a trailing partial-window delta as a full second would fabricate an end-of-run throughput dip; aggregate `throughput_tps` is counter-derived and unaffected. Requests starting inside the window but committing after `measuredEnd` are counted in aggregates (start-time bucketing) but may fall outside the last series point — series sums can therefore differ from aggregate TPS by design.

### ADR-09: Two-Check Independent Auditor for Correctness & Idempotency Verification
**Decision**: Implement an independent post-run auditor with two distinct checks:

**Balance Reconciliation**: Reconstruct account balances by summing the immutable append-only entries log and diffing against the accounts table.

**Idempotency Cardinality Audit**: For every committed idempotency key, count the number of distinct entry-sets/transactions created in the entries table (must equal exactly 1).

**Rationale**: Balance reconciliation catches lost writes or corrupted balances, but cannot detect duplicate execution of an idempotent transaction (since a duplicated transfer writes a second balanced pair of debit/credit entries, leaving net balances correct). The cardinality check explicitly verifies exactly-once semantic guarantees.

**Revision Note (Correctness Fix)**: The cardinality check is now scoped to `(client_id, idempotency_key)` — the actual uniqueness domain enforced by the schema — instead of `idempotency_key` alone, eliminating false positives when two different clients reuse the same key string. A converse check was added: every `committed` key must have exactly one matching transaction (catching committed keys whose ledger effect vanished), reported as `missing_transaction_count`.

**Revision Note (Empirical Exercisability)**: An audit can only catch violations the workload can produce. With fresh UUID keys per request the cardinality checks passed vacuously — duplicates were impossible by construction. Combined with the duplicate-injection flag (ADR-06 revision), committed-key replays now actually traverse the cached-read path and pending collisions traverse the 429 path, so a double-execution bug would surface as `duplicate_effect_count > 0` rather than as an untested invariant.

### ADR-10: Bare CC Path Execution Toggle (Experiment Set 1 Scoping)
**Decision**: Provide a configuration flag (--enable-stage-a=false) in the execution harness to bypass Stage A.

**Rationale**: Prevents Stage A unique index inserts from acting as a baseline bottleneck during pure CC strategy throughput comparisons.

### ADR-11: Database Cleaning Controls & Autovacuum Suppression
**Decision**: Programmatically disable PostgreSQL autovacuum on ledger tables during active benchmark execution windows, executing an explicit VACUUM ANALYZE hook between configuration cells.

**Rationale**: Retry-heavy strategies (OCC) generate significantly higher dead-tuple volumes. Background autovacuum triggering during execution windows would introduce uncommanded background disk and CPU I/O noise correlated with specific CC strategies.

**Revision Note (Correctness Fix)**: `RunCell` now restores default autovacuum via a deferred `RestoreAutovacuum` when the cell (and thus the run) ends, so suppression can no longer leak past the benchmark harness into subsequent uses of the database.

### ADR-12: Infrastructure Controls, Sized Connection Pools & UNLOGGED Diagnostic Variant
**Decision**: Size database connection pool strictly above the maximum tested concurrency level (250 clients + headroom) and include a benchmark harness toggle to run a secondary sweep on PostgreSQL UNLOGGED tables.

**Rationale**: Eliminates connection pool exhaustion as a binding constraint. The UNLOGGED variant serves as a diagnostic to confirm whether a throughput plateau is bound by logical application CC software or physical disk WAL fsync I/O.

**Revision Note (UNLOGGED Impossibility)**: The diagnostic variant cannot run on this schema and now fails fast with an explanatory error instead of a cryptic server error mid-cell. PostgreSQL's `ALTER TABLE ... SET LOGGED | UNLOGGED` refuses to change persistence for any table participating in a foreign key in *either* direction (permanent tables cannot reference unlogged ones and vice versa). Since `entries` references both `transactions` and `accounts`, no per-table toggle ordering can ever succeed — converting the whole set requires an FK-free schema (e.g. dropping the constraints for a dedicated diagnostic run). The harness inspects `pg_constraint` up front and rejects `-unlogged=true` with that explanation.

**Revision Note (Server-Side Connection Budget Prerequisite)**: Both binaries size their pools to 275 connections, but PostgreSQL's default `max_connections` is 100 — at concurrency 250 the pool must actually open ~250 simultaneous connections, so every recorded high-concurrency cell silently depends on the server having been raised accordingly (`max_connections >= 275`, e.g. via `ALTER SYSTEM SET max_connections`). This prerequisite is now recorded because its absence fails loudly but confusingly ("too many clients") far from its cause.

### ADR-13: Lexicographical Account ID Sorting for Optimistic Strategies (OCC/SSI)
Decision: Prior to issuing UPDATE statements, the application logic enforces lexicographical sorting (ascending order) of Account IDs for both OCC and SSI strategies, despite their optimistic nature.

**Rationale**: While OCC and SSI do not utilize application-level row locks, PostgreSQL's MVCC storage engine still acquires implicit row-level exclusive locks during the physical UPDATE execution phase. Updating accounts in debit/credit order under high contention causes true AB-BA physical lock deadlocks (SQLSTATE 40P01) at the storage layer, artificially inflating abort rates. Sorting IDs completely eliminates these structural deadlocks.

### ADR-14: Two-Phase Latency Instrumentation (CC Wait vs. Execution)
**Decision**: The instrumentation harness explicitly isolates concurrency control lock acquisition wait time (cc_wait_latency) from database logical execution time (db_latency).

**Rationale**: Logging the full Go database driver round-trip as "database execution" masks the true cost of pessimistic locking. By splitting SELECT ... FOR UPDATE (Wait) from the subsequent UPDATE and COMMIT (Execution), the telemetry accurately attributes high latencies to lock-queue depth rather than storage engine slowness.

**Revision Note (Comparability Fix)**: `db_latency` is now defined uniformly across all strategies as the full in-transaction window from `BEGIN` to `COMMIT`, lock waits included. Previously Pessimistic excluded its lock-acquisition wait from `db_latency` while OCC/SSI included theirs, making cross-strategy `db_latency` percentile comparisons invalid. `cc_wait_latency` remains the strategy-specific decomposition (lock-queue wait for Pessimistic; 0 for OCC/SSI), and `client_e2e_latency` remains the strategy-agnostic comparison metric.

**Revision Note (Window Alignment Fix)**: The "BEGIN to COMMIT" claim was still violated at the margins: Pessimistic started its `dbStart` clock immediately after `BeginTx` — including both of its `SET LOCAL` round-trips — while OCC and SSI started theirs *after* `applyStatementTimeout`, excluding one SET round-trip. Under contention a SET is a real database round-trip, so OCC/SSI `db_latency` was systematically flattered relative to Pessimistic. All three strategies now take `dbStart` at the same point (immediately after `BeginTx`), making the uniform-window claim true by construction.

### ADR-15: Cumulative Dead Tuple Accounting (HOT Pruning Bypass)
**Decision**: Database bloat (dead tuples generated) is calculated by capturing the delta of n_tup_upd + n_tup_del from pg_stat_user_tables before and after each benchmark cell, rather than sampling n_dead_tup.

**Rationale**: PostgreSQL's internal Heap-Only Tuples (HOT) pruning opportunistically cleans dead tuples during regular scans, even with Autovacuum disabled. Relying on the n_dead_tup gauge at the end of a run heavily undercounts true bloat generation. Tracking cumulative updates/deletes ensures mathematically exact bloat measurements.

### ADR-16: Strict Error Chain Traversal for Anomaly Detection
**Decision**: The unified retry controller utilizes errors.As() to traverse the entire error chain when evaluating retriable database failures, rather than relying on strict interface type assertions (err.(*pq.Error)).

**Rationale**: Serialization anomalies (SSI SQLSTATE 40001) are often raised by PostgreSQL during the COMMIT phase. When the application logic wraps these failures via fmt.Errorf("...: %w"), strict type assertions fail, causing the retry controller to misclassify recoverable CC conflicts as fatal business errors. Error chain traversal guarantees accurate anomaly interception.

**Revision Note (Compliance Completion)**: The retry controller's `IsRetriableError` used `errors.As`, but three sibling classifiers still used strict `err.(*pq.Error)` assertions — `IsInsufficientFundsError` (23514 check violations), `IsDeadlockOrLockTimeout`, and the Stage A unique-violation (23505) branch in the idempotency manager. They worked only because those particular errors happened to arrive unwrapped; any future `%w` wrapping would have silently misclassified them. All three now traverse the error chain via `errors.As`, making this ADR actually true of the whole codebase. The HTTP handlers likewise use `errors.Is` instead of `==` for sentinel mapping. Wrapped-error classification is pinned by regression tests (`cc_test.go`) so the strict-assertion pattern cannot quietly return.

**Revision Note (Telemetry Dead Code Removal)**: `GCMonitor`/`GCPauseSample`/`Sample`/`ParseEnvInt` were declared in the telemetry package, looked like part of the instrumentation story, and were referenced by nothing — dead surface area that invited false confidence about GC telemetry existing. Removed; `GetGCRuntimeSettings` (the only live export) remains.

### ADR-17: Segregation of Business Invariant Rejections
**Decision**: Application-level business logic failures (specifically ErrInsufficientFunds and ErrProcessingRetryLater) are explicitly intercepted, counted in dedicated telemetry metrics (insufficient_funds_count and collision_rejections), and strictly excluded from Concurrency Control failure rates.

**Rationale**: High-throughput strategies hit the mathematical zero-bound of account balances faster than slower strategies. Bounding CC abort rates requires filtering out valid business-rule rejections; otherwise, the fastest strategies are artificially penalized with inflated "failure" metrics.

### ADR-18: Adaptive Hybrid Concurrency Control (Sliding Window & Hot-Swapping)
**Decision**: Implement an adaptive hybrid CC strategy (`AdaptiveStrategy`) that encapsulates both `OCCStrategy` and `PessimisticStrategy`. The strategy routes requests to OCC by default and monitors execution over a 500ms sliding window using atomic request and retry counters (`interval_requests`, `interval_retries`). If the measured abort ratio R_abort = retries / requests exceeds 0.15, the controller safely hot-swaps active routing to Pessimistic locking under a `sync.RWMutex`. If R_abort drops below 0.05, active routing hot-swaps back to OCC. The strategy type reports `StrategyAdaptive`, while `IsolationLevel()` dynamically reflects the currently active underlying strategy.

**Rationale**: Low-skew/low-contention workloads incur unnecessary lock acquisition wait overhead under Pessimistic locking, whereas high-skew workloads experience severe abort storms under pure OCC. Dynamic threshold-based switching optimizes throughput and latency across shifting contention profiles without requiring application restarts or offline reconfiguration.

**Revision Note (Feedback Fix)**: The sliding-window abort ratio is now fed by *per-attempt* abort signals: the unified retry controller invokes `OnRetriableAttempt` (wired to `AdaptiveStrategy.NoteRetry`) on every retriable attempt failure, including aborts whose retry subsequently succeeds. Previously only errors returned after all 5 attempts failed incremented the counter, so the measured ratio was the client-visible-failure ratio; under contention where retries mostly succeeded the controller saw ~0 aborts and never hot-swapped, making ADAPTIVE benchmark results indistinguishable from OCC. Empirical results produced before this fix must be re-run before drawing conclusions about ADAPTIVE.

**Revision Note (Double-Counting Fix)**: `AdaptiveStrategy.ExecuteTransfer` additionally incremented `intervalRetries` whenever a retriable error surfaced to its caller. Since `OnRetriableAttempt` already fires for that same final attempt inside the retry loop, every fully-failed request was counted twice (5 attempt signals + 1 caller signal against 1 request — an abort ratio of 6.0), biasing the controller toward premature Pessimistic switchover under contention spikes. The caller-side increment is removed; `NoteRetry` is now the single, per-attempt source of abort signals.

**Revision Note (Denominator Semantics)**: The formula above reads "retries / requests", but the implementation counts the denominator per *attempt*: `AdaptiveStrategy.ExecuteTransfer` is invoked once per retry-loop attempt, so `intervalRequests` accumulates attempts and `R_abort = failed_attempts / attempts`. Both quantities live in [0,1] and the 0.15/0.05 thresholds (with their hysteresis gap) were reasoned against this per-attempt ratio — including the exhausted-failure case where 5 attempts yield a ratio of 1.0. The wording here is corrected to match the code; changing the code to a logical-request denominator was considered and rejected because it would silently re-tune threshold behavior for no correctness gain.

### ADR-19: Multi-Repetition Statistical Benchmark Harness
**Decision**: Automate empirical data collection across N=5 repetitions per experimental configuration cell (4 strategies x 6 Zipfian skews x 5 repetitions = 120 total cell runs) via `run_repetition_sweep.ps1`, persisting raw JSON metrics to `.\results\set1_ci\`.

**Rationale**: Single-execution benchmark samples are vulnerable to transient OS scheduling jitter, background disk I/O spikes, and Go GC pause variations. Running 5 repetitions per configuration cell yields statistically robust datasets for mean, median, and confidence interval calculations.

**Revision Note (Statistics Fix)**: `compute_confidence_intervals.py` now uses the two-sided 95% **Student-t** critical value matched to the sample size (t ≈ 2.78 at n=5, ≈ 4.30 at n=3) instead of the normal z = 1.96 — with n = 3–5 repetitions the z-based intervals were roughly 2× too narrow. The script also processes exactly one result directory per invocation (default `results/set1_ci`, overridable via CLI argument) instead of pooling `set1`, `set1_repetitions`, and `set1_ci` together, which mixed single-run, 3-rep, and 5-rep samples and corrupted both means and intervals.

**Revision Note (Plotting Compliance)**: The same one-directory-per-invocation rule is now applied to the plotting pipeline: `plot_heatmap.py` previously pooled `results/set1_ci` with `results/set1`, double-weighting every cell present in both directories. It now takes exactly one directory (default `results/set1_ci`). The byte-identical duplicate sweep script `run_final_sweep.ps1` was deleted in favor of `run_repetition_sweep.ps1` — two copies of the same harness invite divergent edits.

### ADR-20: Formal Specification of Two-Stage Idempotency under Process Crash Failures
**Decision**: Formally model the two-stage idempotency execution protocol in TLA+ (`Idempotency.tla`), capturing Stage A pending key registration, Stage B atomic transfer and state update, arbitrary worker process crashes, and TTL pending key cleanup.

**Rationale**: Empirical testing alone cannot prove safety across all possible interleavings of process crashes and retries. Formal model checking rigorously verifies the `Safety` invariant (at-most-once entry set creation per idempotency key) and `Liveness` invariant (eventual resolution of all pending keys to terminal committed or failed states) under arbitrary crash timings.

**Revision Note (Spec/Impl Alignment)**: The spec and implementation have been aligned. (a) `TTLCleanup` transitions `pending -> failed`, matching the implementation's state-transition (not DELETE) cleaner. (b) A new `Reclaim` action models the CAS-guarded `failed -> pending` re-claim by a retryer, matching the manager's reclaim path. (c) `Idempotency.cfg` was added so TLC actually checks `Safety` and `EntryConsistency` as invariants and `Liveness` as a temporal property — previously no model configuration existed and none of the properties were checkable as stated.

**Revision Note (Verified Model — Syntax, Liveness, and Honesty Fixes)**: Three defects were found by actually running TLC on the committed artifact:
1. **The module did not parse.** Line 44 contained `{<<c, k, "tx_entry_set">>>}` — three closing angle brackets instead of two — so every claim of TLC verification was false as committed; no tool could even read the file.
2. **`Liveness` failed after the syntax fix.** With only `WF_vars(Next)`, TLC produced a counterexample in which a key stayed `pending` with its worker in-flight forever while another key cycled `StageA -> Crash -> Reclaim` infinitely. Neither `StageB_Commit` nor `Crash` carries per-action fairness, so pending keys had no guaranteed route to a terminal state. The spec now adds per-key weak fairness on `TTLCleanup`.
3. **`TTLCleanup`'s `ck \notin inFlight` precondition assumed away the exact race the protocol defends against** (TTL expiry flipping a key while Stage B is mid-flight). The guard is removed to match the implementation's unconditional `UPDATE ... WHERE state = 'pending'`; a late commit now finds no pending row and aborts *within the model*, so at-most-once is established rather than presupposed.

With these fixes TLC reports: 625 distinct states, both invariants hold, and the temporal property verifies ("Model checking completed. No error has been found."). Residual modeling limitation, documented for honesty: `StageB_Commit` *replaces* the entry set rather than accumulating effects, so the `Safety` invariant cannot distinguish double execution from single execution at the effect level; the property is meaningful only because the `pending`-row precondition now genuinely blocks late commits in the model.

### ADR-21: Exact Zipfian Sampling via Cumulative Inversion & Per-Worker PRNG Isolation
* **Decision**: Replace the approximate theta < 1 power-law Zipf sampler with exact inversion of a precomputed cumulative distribution over the finite account population, and give each benchmark worker its own `rand.Rand` instance (`ZipfGenerator.NewSampler(seed)`) for both account selection and transfer amounts.
* **Rationale**: The approximate formula silently rewrote theta = 1.0 to 0.9999 to avoid a division-by-zero singularity, so benchmark cells labelled "skew 1.0" did not actually run at theta 1.0. Over a finite domain the Zipf distribution is well defined for any theta >= 0, so exact inversion (binary search over the cumulative mass table) removes the mislabeling. Per-worker PRNGs eliminate contention on the global `math/rand` mutex, which previously inflated client-side latency at high concurrency as a measurement artifact rather than a property of any CC strategy.
* **Revision Note (Transfer Amounts)**: Transfer amounts are now quantized to whole cents before insertion. Full-precision float64 amounts were stored as `round4(balance - x)` in the accounts table while the auditor reconstructed balances from `round4(x)` entries — two different roundings of the same operation whose half-cent drift accumulates ~√N·5·10⁻⁵ on hot accounts and can exceed the auditor's fixed 10⁻⁴ tolerance. Cent-quantized amounts round identically on both sides. A fuller fix would move the ledger to integer minor units; that remains future work (see ADR-01).

### ADR-22: Collision Rejection Accounting as a First-Class Metric
**Decision**: `collision_rejections` (429 pending-collision responses from ADR-07) is now actually counted by the benchmark runner — requests returning `ErrProcessingRetryLater` increment it and are excluded from `client_visible_failures`.
* **Rationale**: The metric was declared and reported but never incremented, always reading 0. Counting it separately keeps ADR-17's segregation (business outcomes vs CC failures) intact while making pending-collision rates observable for Experiment Set 2.

### ADR-23: No Synthetic Data in Plotting Artifacts
**Decision**: `plot_timeseries.py` no longer substitutes synthetic "fallback" time series when real `time_series` data is missing; it exits with an error instead, and the hardcoded hot-swap annotation derived from that synthetic data was removed.
* **Rationale**: Fabricated curves in a paper artifact misrepresent empirical results. A missing input must fail loudly so the underlying benchmark cell is re-run.

### ADR-24: Result Provenance & Experiment Isolation
**Decision**:
1. Diagnostic/investigation runs are written to a dedicated investigation directory (`results/set2_investigation`) and never into primary experiment directories (`results/set1`, `results/set2`). `investigate_anomalies.ps1` previously polluted the main Set 2 directory with diagnostic cells.
2. The consolidated archive generator (`consolidate_results.ps1`) stamps every archive with its generation timestamp (UTC) and an explicit notice that result files were captured by different instrument versions (e.g. legacy `abort_retry_rate_pct` vs. current `retries_per_request`/`client_failure_rate_pct`/`cached_hits`), so cross-file comparisons check field sets first.
3. Instrument-version changes are recorded as revision notes in this document, and results produced before a correctness-affecting revision are marked non-comparable until re-run.
* **Rationale**: Silent mixing of diagnostic and primary data, or of pre- and post-fix instrumentation, corrupts downstream statistics while looking perfectly routine. Provenance must be visible at the point of consumption (the archive header), not reconstructible only from git archaeology.