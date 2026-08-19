package cc

import (
	"sync/atomic"
	"testing"
	"errors"

	"github.com/lib/pq"
	"github.com/ledger/skewed-ledger/pkg/ledger"
)

func TestIsInsufficientFundsError(t *testing.T) {
	if !IsInsufficientFundsError(ledger.ErrInsufficientFunds) {
		t.Errorf("Expected ErrInsufficientFunds to be recognized as insufficient funds error")
	}

	pgCheckViolationErr := &pq.Error{Code: "23514"}
	if !IsInsufficientFundsError(pgCheckViolationErr) {
		t.Errorf("Expected PG code 23514 (check_violation) to be recognized as insufficient funds error")
	}

	otherErr := errors.New("some random error")
	if IsInsufficientFundsError(otherErr) {
		t.Errorf("Did not expect random error to be recognized as insufficient funds error")
	}

	if IsInsufficientFundsError(nil) {
		t.Errorf("Did not expect nil to be recognized as insufficient funds error")
	}
}

func TestIsRetriableError(t *testing.T) {
	retriable, isDeadlock := IsRetriableError(ledger.ErrVersionMismatch)
	if !retriable || isDeadlock {
		t.Errorf("Expected ErrVersionMismatch to be retriable, not deadlock: got retriable=%v, isDeadlock=%v", retriable, isDeadlock)
	}

	pgDeadlockErr := &pq.Error{Code: "40P01"}
	retriable, isDeadlock = IsRetriableError(pgDeadlockErr)
	if !retriable || !isDeadlock {
		t.Errorf("Expected 40P01 to be retriable deadlock: got retriable=%v, isDeadlock=%v", retriable, isDeadlock)
	}
}

func TestAdaptiveStrategyThresholds(t *testing.T) {
	strat := NewAdaptiveStrategy("2000ms")
	if strat.ActiveStrategyType() != StrategyOCC {
		t.Errorf("Expected initial strategy to be OCC, got %s", strat.ActiveStrategyType())
	}

	// Simulate high abort ratio (2 / 10 = 0.20 > 0.15)
	atomic.StoreInt64(&strat.intervalRequests, 10)
	atomic.StoreInt64(&strat.intervalRetries, 2)
	strat.evaluateWindow()

	if strat.ActiveStrategyType() != StrategyPessimistic {
		t.Errorf("Expected strategy to switch to PESSIMISTIC, got %s", strat.ActiveStrategyType())
	}

	// Simulate low abort ratio (0 / 10 = 0.00 < 0.05)
	atomic.StoreInt64(&strat.intervalRequests, 10)
	atomic.StoreInt64(&strat.intervalRetries, 0)
	strat.evaluateWindow()

	if strat.ActiveStrategyType() != StrategyOCC {
		t.Errorf("Expected strategy to switch back to OCC, got %s", strat.ActiveStrategyType())
	}
}
