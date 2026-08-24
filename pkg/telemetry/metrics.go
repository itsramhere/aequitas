package telemetry

import (
	"math"
	"sort"
	"time"
)

type LatencyPercentiles struct {
	P50 time.Duration `json:"p50"`
	P95 time.Duration `json:"p95"`
	P99 time.Duration `json:"p99"`
}

type TimeSeriesPoint struct {
	Second  int     `json:"second"`
	TPS     float64 `json:"tps"`
	Retries int64   `json:"retries"`
}

type CellMetricsReport struct {
	Strategy               string             `json:"strategy"`
	SkewTheta              float64            `json:"skew_theta"`
	Concurrency            int                `json:"concurrency"`
	Duration               time.Duration      `json:"duration"`
	TotalRequests          int64              `json:"total_requests"`
	CommittedTxns          int64              `json:"committed_txns"`
	ClientVisibleFailures  int64              `json:"client_visible_failures"`
	CollisionRejections    int64              `json:"collision_rejections"`
	CachedHits             int64              `json:"cached_hits"`
	InternalRetries        int64              `json:"internal_retries"`
	DeadlockCount          int64              `json:"deadlock_count"`
	InsufficientFundsCount int64              `json:"insufficient_funds_count"`
	Throughput             float64            `json:"throughput_tps"`
	RetriesPerRequest      float64            `json:"retries_per_request"`
	ClientFailureRate      float64            `json:"client_failure_rate_pct"`
	RealizedHotRowHitRate  float64            `json:"realized_hot_row_hit_rate_pct"`
	DeadTuplesGenerated    int64              `json:"dead_tuples_generated"`
	DeadTuplesPerCommit    float64            `json:"dead_tuples_per_commit"`
	AppLatency             LatencyPercentiles `json:"app_latency"`
	DBLatency              LatencyPercentiles `json:"db_latency"`
	CCWaitLatency          LatencyPercentiles `json:"cc_wait_latency"`
	WALWaitProxy           LatencyPercentiles `json:"wal_wait_proxy"`
	ClientEndToEnd         LatencyPercentiles `json:"client_e2e_latency"`
	ClientBackoffRejection LatencyPercentiles `json:"client_backoff_rejection_latency"`
	RuntimeSettings        GCRuntimeSettings  `json:"runtime_settings"`
	TimeSeries             []TimeSeriesPoint  `json:"time_series"`
}

// CalculatePercentiles returns nearest-rank percentiles (ceil(p*n)-1), the
// standard order-statistic definition. The previous floor(p*n/100) indexing
// was systematically off by one rank.
func CalculatePercentiles(durations []time.Duration) LatencyPercentiles {
	if len(durations) == 0 {
		return LatencyPercentiles{}
	}

	sorted := make([]time.Duration, len(durations))
	copy(sorted, durations)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	n := len(sorted)
	rank := func(p float64) time.Duration {
		idx := int(math.Ceil(p*float64(n))) - 1
		if idx < 0 {
			idx = 0
		}
		if idx >= n {
			idx = n - 1
		}
		return sorted[idx]
	}

	return LatencyPercentiles{
		P50: rank(0.50),
		P95: rank(0.95),
		P99: rank(0.99),
	}
}
