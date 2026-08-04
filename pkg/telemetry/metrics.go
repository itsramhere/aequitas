package telemetry

import (
	"sort"
	"time"
)

type LatencyPercentiles struct {
	P50 time.Duration `json:"p50"`
	P95 time.Duration `json:"p95"`
	P99 time.Duration `json:"p99"`
}

type CellMetricsReport struct {
	Strategy              string             `json:"strategy"`
	SkewTheta             float64            `json:"skew_theta"`
	Concurrency           int                `json:"concurrency"`
	Duration              time.Duration      `json:"duration"`
	TotalRequests         int64              `json:"total_requests"`
	CommittedTxns         int64              `json:"committed_txns"`
	ClientVisibleFailures int64              `json:"client_visible_failures"`
	InternalRetries       int64              `json:"internal_retries"`
	DeadlockCount         int64              `json:"deadlock_count"`
	Throughput            float64            `json:"throughput_tps"`
	AbortRetryRate        float64            `json:"abort_retry_rate_pct"`
	RealizedHotRowHitRate float64            `json:"realized_hot_row_hit_rate_pct"`
	DeadTuplesGenerated   int64              `json:"dead_tuples_generated"`
	DeadTuplesPerCommit   float64            `json:"dead_tuples_per_commit"`
	AppLatency            LatencyPercentiles `json:"app_latency"`
	DBLatency             LatencyPercentiles `json:"db_latency"`
	WALWaitProxy          LatencyPercentiles `json:"wal_wait_proxy"`
	ClientEndToEnd        LatencyPercentiles `json:"client_e2e_latency"`
	ClientBackoffRejection LatencyPercentiles `json:"client_backoff_rejection_latency"`
	RuntimeSettings       GCRuntimeSettings  `json:"runtime_settings"`
}

func CalculatePercentiles(durations []time.Duration) LatencyPercentiles {
	if len(durations) == 0 {
		return LatencyPercentiles{}
	}

	sorted := make([]time.Duration, len(durations))
	copy(sorted, durations)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

	n := len(sorted)
	return LatencyPercentiles{
		P50: sorted[n*50/100],
		P95: sorted[n*95/100],
		P99: sorted[n*99/100],
	}
}
