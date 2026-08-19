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
    /\ ledgerEntries'   = [ledgerEntries EXCEPT ![ck] = {<<c, k, "tx_entry_set">>>}]
    /\ inFlight'        = inFlight \ {ck}

-----------------------------------------------------------------------------
\* Crash Action: Worker process dies mid-flight during Stage B
\* Uncommitted ledger entries roll back completely, key remains in 'pending' or transitions to 'failed'
Crash(c, k) ==
    LET ck == <<c, k>> IN
    /\ ck \in inFlight
    /\ idempotencyKeys[ck] = "pending"
    /\ inFlight'        = inFlight \ {ck}
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "failed"]
    /\ UNCHANGED <<ledgerEntries>>

\* TTL Cleanup: Expired pending key transitions to terminal 'failed' state
TTLCleanup(c, k) ==
    LET ck == <<c, k>> IN
    /\ ck \notin inFlight
    /\ idempotencyKeys[ck] = "pending"
    /\ idempotencyKeys' = [idempotencyKeys EXCEPT ![ck] = "failed"]
    /\ UNCHANGED <<ledgerEntries, inFlight>>

-----------------------------------------------------------------------------
Next ==
    \E c \in Clients, k \in Keys :
        \/ StageA(c, k)
        \/ StageA_Collision(c, k)
        \/ StageB_Commit(c, k)
        \/ Crash(c, k)
        \/ TTLCleanup(c, k)

Spec == Init /\ [][Next]_vars /\ WF_vars(Next)

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
