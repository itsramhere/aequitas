package ledger

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Standard Ledger Errors
var (
	ErrInsufficientFunds    = errors.New("insufficient funds for transfer")
	ErrAccountNotFound      = errors.New("one or both accounts not found")
	ErrVersionMismatch      = errors.New("optimistic concurrency version mismatch")
	ErrProcessingRetryLater = errors.New("request currently processing, retry later")
	ErrSameAccountTransfer  = errors.New("debited and credited accounts must be distinct")
	ErrInvalidAmount        = errors.New("transfer amount must be positive")
)

type IdempotencyState string

const (
	StatePending   IdempotencyState = "pending"
	StateCommitted IdempotencyState = "committed"
	StateFailed    IdempotencyState = "failed"
)

type EntryType string

const (
	EntryDebit  EntryType = "debit"
	EntryCredit EntryType = "credit"
)

type Account struct {
	ID        int64     `json:"id"`
	Balance   float64   `json:"balance"`
	Version   int64     `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IdempotencyKey struct {
	ClientID        string           `json:"client_id"`
	IdempotencyKey  string           `json:"idempotency_key"`
	State           IdempotencyState `json:"state"`
	ResponsePayload []byte           `json:"response_payload"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
}

type Transaction struct {
	ID             uuid.UUID `json:"id"`
	ClientID       string    `json:"client_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	CreatedAt      time.Time `json:"created_at"`
}

type Entry struct {
	ID            int64     `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	AccountID     int64     `json:"account_id"`
	Amount        float64   `json:"amount"`
	EntryType     EntryType `json:"entry_type"`
	CreatedAt     time.Time `json:"created_at"`
}

type TransferParams struct {
	ClientID          string  `json:"client_id"`
	IdempotencyKey    string  `json:"idempotency_key"`
	DebitedAccountID  int64   `json:"debited_account_id"`
	CreditedAccountID int64   `json:"credited_account_id"`
	Amount            float64 `json:"amount"`
}

func (p TransferParams) Validate() error {
	if p.DebitedAccountID == p.CreditedAccountID {
		return ErrSameAccountTransfer
	}
	if p.Amount <= 0 {
		return ErrInvalidAmount
	}
	return nil
}

type TxResult struct {
	TransactionID  uuid.UUID     `json:"transaction_id"`
	ClientID       string        `json:"client_id"`
	IdempotencyKey string        `json:"idempotency_key"`
	Status         string        `json:"status"`
	Attempts       int           `json:"attempts"`
	AppLatency     time.Duration `json:"app_latency"`
	DBLatency      time.Duration `json:"db_latency"`
	CCWaitLatency  time.Duration `json:"cc_wait_latency"`
	WALWaitProxy   time.Duration `json:"wal_wait_proxy"`
	IsCached       bool          `json:"is_cached"`
}
