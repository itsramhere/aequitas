package cc

import (
	"context"
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ledger/skewed-ledger/pkg/ledger"
)

type AdaptiveStrategy struct {
	occStrategy         *OCCStrategy
	pessimisticStrategy *PessimisticStrategy
	activeStrategy      Strategy
	mu                  sync.RWMutex

	intervalRequests int64
	intervalRetries  int64

	windowInterval time.Duration
	stopCh         chan struct{}
	running        int32
}

func NewAdaptiveStrategy(lockTimeout string) *AdaptiveStrategy {
	occ := NewOCCStrategy()
	pess := NewPessimisticStrategy(lockTimeout)
	return &AdaptiveStrategy{
		occStrategy:         occ,
		pessimisticStrategy: pess,
		activeStrategy:      occ,
		windowInterval:      500 * time.Millisecond,
		stopCh:              make(chan struct{}),
	}
}

func (s *AdaptiveStrategy) Init(ctx context.Context) {
	if atomic.CompareAndSwapInt32(&s.running, 0, 1) {
		go s.monitorLoop(ctx)
	}
}

func (s *AdaptiveStrategy) Start(ctx context.Context) {
	s.Init(ctx)
}

func (s *AdaptiveStrategy) Stop() {
	if atomic.CompareAndSwapInt32(&s.running, 1, 0) {
		close(s.stopCh)
	}
}

func (s *AdaptiveStrategy) monitorLoop(ctx context.Context) {
	ticker := time.NewTicker(s.windowInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-s.stopCh:
			return
		case <-ticker.C:
			s.evaluateWindow()
		}
	}
}

func (s *AdaptiveStrategy) evaluateWindow() {
	reqs := atomic.SwapInt64(&s.intervalRequests, 0)
	retries := atomic.SwapInt64(&s.intervalRetries, 0)

	if reqs <= 0 {
		return
	}

	abortRatio := float64(retries) / float64(reqs)

	s.mu.RLock()
	currentType := s.activeStrategy.Type()
	s.mu.RUnlock()

	if abortRatio > 0.15 && currentType != StrategyPessimistic {
		s.mu.Lock()
		s.activeStrategy = s.pessimisticStrategy
		s.mu.Unlock()
	} else if abortRatio < 0.05 && currentType != StrategyOCC {
		s.mu.Lock()
		s.activeStrategy = s.occStrategy
		s.mu.Unlock()
	}
}

func (s *AdaptiveStrategy) Type() StrategyType {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return StrategyAdaptive
}

func (s *AdaptiveStrategy) IsolationLevel() sql.IsolationLevel {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.activeStrategy.IsolationLevel()
}

func (s *AdaptiveStrategy) ActiveStrategyType() StrategyType {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.activeStrategy.Type()
}

func (s *AdaptiveStrategy) ExecuteTransfer(
	ctx context.Context,
	db *sql.DB,
	params ledger.TransferParams,
	opts TransferOptions,
) (*ledger.TxResult, error) {
	atomic.AddInt64(&s.intervalRequests, 1)

	s.mu.RLock()
	active := s.activeStrategy
	s.mu.RUnlock()

	result, err := active.ExecuteTransfer(ctx, db, params, opts)
	// Retry accounting note (ADR-18 revision): per-attempt aborts are counted
	// exclusively via NoteRetry, which the unified retry controller invokes
	// from inside its retry loop — including the final exhausted attempt.
	// Incrementing intervalRetries here as well would double-count the last
	// attempt of every fully-failed request and inflate the abort ratio.
	return result, err
}

// NoteRetry records a single in-flight retriable attempt failure. It is wired
// as UnifiedRetryController.OnRetriableAttempt so the sliding window measures
// the true per-attempt abort ratio, not just client-visible exhaustion.
func (s *AdaptiveStrategy) NoteRetry() {
	atomic.AddInt64(&s.intervalRetries, 1)
}
