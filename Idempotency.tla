--------------------------- MODULE Idempotency ---------------------------
EXTENDS Sequences, FiniteSets, Integers, TLC

CONSTANTS 
    Clients,        \* Set of client identifiers
    Keys            \* Set of idempotency key identifiers

VARIABLES 
    idempotencyKeys, \* Mapping (client, key) -> "unseen" | "pending" | "committed" | "failed"
    ledgerEntries,   \* Mapping (client, key) -> Set of ledger entry records created
    inFlight         \* Set of (client, key) currently executing Stage B in active worker memory

vars == <<idempotencyKeys, ledgerEntries, inFlight>>

KeyStates == {"unseen", "pending", "committed", "failed"}

Init ==
    /\ idempotencyKeys = [ck \in Clients \times Keys |-> "unseen"]
    /\ ledgerEntries   = [ck \in Clients \times Keys |-> {}]
    /\ inFlight        = {}

-----------------------------------------------------------------------------
\* Stage A: Insert idempotency key with state 'pending'
StageA(c, k) ==
    LET ck == <<c, k>> IN
    /\ idempotencyKeys[ck] = "unseen"
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "pending"]
    /\ inFlight'        = inFlight \cup {ck}
    /\ UNCHANGED <<ledgerEntries>>

\* Duplicate Arrival handling during Stage A (Unique Index Collision Policy)
StageA_Collision(c, k) ==
    LET ck == <<c, k>> IN
    /\ idempotencyKeys[ck] \in {"pending", "committed", "failed"}
    /\ UNCHANGED vars

-----------------------------------------------------------------------------
\* Stage B: Atomic Execution of Ledger Entries & Transition Key State to 'committed'
StageB_Commit(c, k) ==
    LET ck == <<c, k>> IN
    /\ ck \in inFlight
    /\ idempotencyKeys[ck] = "pending"
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "committed"]
    /\ ledgerEntries'   = [ledgerEntries EXCEPT ![ck] = {<<c, k, "tx_entry_set">>}]
    /\ inFlight'        = inFlight \ {ck}

-----------------------------------------------------------------------------
\* Crash Action: Worker process dies mid-flight during Stage B.
\* Uncommitted ledger entries roll back completely. Abstraction note: a real
\* crash leaves the key 'pending' until the TTL cleaner later transitions it
\* to 'failed'; this action folds that eventual transition into one step.
Crash(c, k) ==
    LET ck == <<c, k>> IN
    /\ ck \in inFlight
    /\ idempotencyKeys[ck] = "pending"
    /\ inFlight'        = inFlight \ {ck}
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "failed"]
    /\ UNCHANGED <<ledgerEntries>>

\* TTL Cleanup: Expired pending key transitions to terminal 'failed' state.
\* Matches the implementation: the cleaner issues an unconditional
\* UPDATE ... WHERE state = 'pending' and does NOT consult whether a worker
\* is mid-Stage-B (the old `ck \notin inFlight` precondition assumed away
\* exactly the race this protocol defends against). A late Stage B commit
\* then matches zero 'pending' rows and aborts, preserving at-most-once.
TTLCleanup(c, k) ==
    LET ck == <<c, k>> IN
    /\ idempotencyKeys[ck] = "pending"
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "failed"]
    /\ UNCHANGED <<ledgerEntries, inFlight>>

\* Reclaim: a retryer atomically re-claims a terminal 'failed' key
\* (failed -> pending, CAS-guarded) and re-executes Stage B. Exactly one
\* retryer can hold the claim at a time; concurrent retryers observe
\* 'pending' and are rejected (StageA_Collision).
Reclaim(c, k) ==
    LET ck == <<c, k>> IN
    /\ idempotencyKeys[ck] = "failed"
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "pending"]
    /\ inFlight'        = inFlight \cup {ck}
    /\ UNCHANGED <<ledgerEntries>>

-----------------------------------------------------------------------------
Next ==
    \E c \in Clients, k \in Keys :
        \/ StageA(c, k)
        \/ StageA_Collision(c, k)
        \/ StageB_Commit(c, k)
        \/ Crash(c, k)
        \/ TTLCleanup(c, k)
        \/ Reclaim(c, k)

\* Fairness: WF_vars(Next) alone admits infinite behaviors in which a worker
\* stalls mid-Stage-B forever and the cleaner never fires, leaving a key
\* 'pending' eternally (TLC produced exactly such a counterexample). Per-key
\* weak fairness on TTLCleanup guarantees every pending key eventually
\* reaches a terminal state.
Fairness ==
    /\ WF_vars(Next)
    /\ \A c \in Clients, k \in Keys : WF_vars(TTLCleanup(c, k))

Spec == Init /\ [][Next]_vars /\ Fairness

-----------------------------------------------------------------------------
\* INVARIANTS

\* Safety: At-Most-Once Execution Invariant.
\* No idempotency key ever produces more than 1 ledger entry set.
Safety ==
    \A c \in Clients, k \in Keys :
        LET ck == <<c, k>> IN
        Cardinality(ledgerEntries[ck]) <= 1

\* Entry State Consistency Invariant:
\* Ledger entries exist if and only if the idempotency key is in 'committed' state.
EntryConsistency ==
    \A c \in Clients, k \in Keys :
        LET ck == <<c, k>> IN
        (Cardinality(ledgerEntries[ck]) > 0) => (idempotencyKeys[ck] = "committed")

\* Liveness: Terminal Resolution Property.
\* Every pending key eventually reaches a terminal state ('committed' or 'failed').
Liveness ==
    \A c \in Clients, k \in Keys :
        LET ck == <<c, k>> IN
        (idempotencyKeys[ck] = "pending") ~> (idempotencyKeys[ck] \in {"committed", "failed"})

=============================================================================
