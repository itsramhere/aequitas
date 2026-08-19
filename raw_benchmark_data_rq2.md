# Audit and Idempotency JSON Files

## Audit_OCC.json

``` json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 35755,
  "duplicate_effect_count": 0,
  "passed": true
}
```

## Audit_PESSIMISTIC.json

``` json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 33032,
  "duplicate_effect_count": 0,
  "passed": true
}
```

## Audit_SSI.json

``` json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 16186,
  "duplicate_effect_count": 0,
  "passed": true
}
```

## Set2_OCC_Idempotency.json

``` json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 33578,
  "committed_txns": 32321,
  "client_visible_failures": 1257,
  "internal_retries": 12531,
  "deadlock_count": 27,
  "throughput_tps": 1077.3666666666666,
  "abort_retry_rate_pct": 41.06260051224015,
  "realized_hot_row_hit_rate_pct": 3.672031516813898,
  "dead_tuples_generated": 45653,
  "dead_tuples_per_commit": 1.412487237399833,
  "app_latency": {
    "p50": 0,
    "p95": 1529300,
    "p99": 7266500
  },
  "db_latency": {
    "p50": 11451200,
    "p95": 28514400,
    "p99": 56697800
  },
  "wal_wait_proxy": {
    "p50": 1565100,
    "p95": 9703300,
    "p99": 17339000
  },
  "client_e2e_latency": {
    "p50": 20782000,
    "p95": 81381700,
    "p99": 217689500
  },
  "client_backoff_rejection_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "runtime_settings": {
    "gogc": "100 (default)",
    "gomaxprocs": 12,
    "num_cpu": 12
  }
}
```

## Set2_PESSIMISTIC_Idempotency.json

``` json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 27669,
  "committed_txns": 27669,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 922.3,
  "abort_retry_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.7701350902112485,
  "dead_tuples_generated": 71149,
  "dead_tuples_per_commit": 2.571433734504319,
  "app_latency": {
    "p50": 511600,
    "p95": 1541600,
    "p99": 3725200
  },
  "db_latency": {
    "p50": 12499600,
    "p95": 72517600,
    "p99": 1128200300
  },
  "wal_wait_proxy": {
    "p50": 1016100,
    "p95": 12583200,
    "p99": 32314300
  },
  "client_e2e_latency": {
    "p50": 17133100,
    "p95": 82802200,
    "p99": 1132455900
  },
  "client_backoff_rejection_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "runtime_settings": {
    "gogc": "100 (default)",
    "gomaxprocs": 12,
    "num_cpu": 12
  }
}
```

## Set2_SSI_Idempotency.json

``` json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 31930,
  "committed_txns": 11930,
  "client_visible_failures": 20000,
  "internal_retries": 4021,
  "deadlock_count": 50,
  "throughput_tps": 397.6666666666667,
  "abort_retry_rate_pct": 75.23019104290636,
  "realized_hot_row_hit_rate_pct": 3.5715218470056675,
  "dead_tuples_generated": 114448,
  "dead_tuples_per_commit": 9.593294216261526,
  "app_latency": {
    "p50": 0,
    "p95": 1528500,
    "p99": 7969300
  },
  "db_latency": {
    "p50": 10731200,
    "p95": 31613200,
    "p99": 53481700
  },
  "wal_wait_proxy": {
    "p50": 1620000,
    "p95": 13682000,
    "p99": 30091100
  },
  "client_e2e_latency": {
    "p50": 20204200,
    "p95": 68279500,
    "p99": 153239800
  },
  "client_backoff_rejection_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "runtime_settings": {
    "gogc": "100 (default)",
    "gomaxprocs": 12,
    "num_cpu": 12
  }
}
```
