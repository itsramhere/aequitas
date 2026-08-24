# Benchmark Raw Data Archive

_Generated: 2026-08-24 16:15:27 UTC by consolidate_results.ps1._

_Note: result files were captured by different instrument versions over the life of this project (e.g. older files carry `abort_retry_rate_pct` while current code emits `retries_per_request`/`client_failure_rate_pct`). Field sets are NOT uniform across files; check each JSON's fields before cross-file comparison._

### results\set1\OCC_skew_0.4_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 84788,
  "committed_txns": 84788,
  "client_visible_failures": 0,
  "internal_retries": 1994,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2826.266666666667,
  "retries_per_request": 0.023517478888521962,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.23711787052854336,
  "dead_tuples_generated": 194308,
  "dead_tuples_per_commit": 2.2916922205972545,
  "app_latency": {
    "p50": 0,
    "p95": 1762100,
    "p99": 8087600
  },
  "db_latency": {
    "p50": 12761300,
    "p95": 35228100,
    "p99": 78321700
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2010800,
    "p95": 12901100,
    "p99": 22854400
  },
  "client_e2e_latency": {
    "p50": 13569000,
    "p95": 39177100,
    "p99": 93070100
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



### results\set1\OCC_skew_0.6_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 79777,
  "committed_txns": 79465,
  "client_visible_failures": 312,
  "internal_retries": 7676,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2648.8333333333335,
  "retries_per_request": 0.09621820825551224,
  "client_failure_rate_pct": 0.39109016383168077,
  "realized_hot_row_hit_rate_pct": 1.0460095507706384,
  "dead_tuples_generated": 172647,
  "dead_tuples_per_commit": 2.1726168753539294,
  "app_latency": {
    "p50": 0,
    "p95": 1532500,
    "p99": 8090600
  },
  "db_latency": {
    "p50": 12308600,
    "p95": 32319200,
    "p99": 61894700
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2011100,
    "p95": 12284800,
    "p99": 21297400
  },
  "client_e2e_latency": {
    "p50": 13478600,
    "p95": 41334100,
    "p99": 101559800
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



### results\set1\OCC_skew_0.8_conc_10.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 68668,
  "committed_txns": 68568,
  "client_visible_failures": 100,
  "internal_retries": 7159,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2285.6,
  "retries_per_request": 0.10425525717947225,
  "client_failure_rate_pct": 0.14562824022834508,
  "realized_hot_row_hit_rate_pct": 3.7156796365316103,
  "dead_tuples_generated": 154629,
  "dead_tuples_per_commit": 2.2551190059502977,
  "app_latency": {
    "p50": 0,
    "p95": 578800,
    "p99": 1009700
  },
  "db_latency": {
    "p50": 2949300,
    "p95": 4293700,
    "p99": 7040900
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 537800,
    "p95": 1008000,
    "p99": 1584900
  },
  "client_e2e_latency": {
    "p50": 3275300,
    "p95": 9166200,
    "p99": 23770300
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



### results\set1\OCC_skew_0.8_conc_100.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 98969,
  "committed_txns": 94766,
  "client_visible_failures": 4203,
  "internal_retries": 37615,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3158.866666666667,
  "retries_per_request": 0.38006850629995254,
  "client_failure_rate_pct": 4.246784346613586,
  "realized_hot_row_hit_rate_pct": 3.673636073975345,
  "dead_tuples_generated": 211576,
  "dead_tuples_per_commit": 2.2326150729164467,
  "app_latency": {
    "p50": 0,
    "p95": 1013000,
    "p99": 2122200
  },
  "db_latency": {
    "p50": 5191900,
    "p95": 9954300,
    "p99": 15973500
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 567600,
    "p95": 3540000,
    "p99": 6690700
  },
  "client_e2e_latency": {
    "p50": 5735600,
    "p95": 32101800,
    "p99": 238166100
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



### results\set1\OCC_skew_0.8_conc_250.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 64499,
  "committed_txns": 60705,
  "client_visible_failures": 3794,
  "internal_retries": 19535,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2023.5,
  "retries_per_request": 0.30287291275833733,
  "client_failure_rate_pct": 5.882261740492101,
  "realized_hot_row_hit_rate_pct": 3.7202089910871448,
  "dead_tuples_generated": 132823,
  "dead_tuples_per_commit": 2.18800757762952,
  "app_latency": {
    "p50": 0,
    "p95": 1014400,
    "p99": 2998000
  },
  "db_latency": {
    "p50": 5500600,
    "p95": 20548900,
    "p99": 73203200
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 689100,
    "p95": 4059400,
    "p99": 12056200
  },
  "client_e2e_latency": {
    "p50": 6111500,
    "p95": 52541000,
    "p99": 411747400
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



### results\set1\OCC_skew_0.8_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 69862,
  "committed_txns": 67012,
  "client_visible_failures": 2850,
  "internal_retries": 26009,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2233.733333333333,
  "retries_per_request": 0.3722910881452005,
  "client_failure_rate_pct": 4.079470957029573,
  "realized_hot_row_hit_rate_pct": 3.727893955270415,
  "dead_tuples_generated": 148052,
  "dead_tuples_per_commit": 2.2093356413776637,
  "app_latency": {
    "p50": 0,
    "p95": 1507700,
    "p99": 3823100
  },
  "db_latency": {
    "p50": 7239000,
    "p95": 16843900,
    "p99": 29955400
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1052400,
    "p95": 5450000,
    "p99": 10250200
  },
  "client_e2e_latency": {
    "p50": 8233100,
    "p95": 40290400,
    "p99": 131036700
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



### results\set1\OCC_skew_0_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 92322,
  "committed_txns": 92322,
  "client_visible_failures": 0,
  "internal_retries": 1405,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3077.4,
  "retries_per_request": 0.015218474469790516,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.00907874617737003,
  "dead_tuples_generated": 206777,
  "dead_tuples_per_commit": 2.2397370074305147,
  "app_latency": {
    "p50": 0,
    "p95": 1511900,
    "p99": 7151500
  },
  "db_latency": {
    "p50": 12929400,
    "p95": 30144400,
    "p99": 56339500
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2002200,
    "p95": 12554900,
    "p99": 20838600
  },
  "client_e2e_latency": {
    "p50": 13555800,
    "p95": 32902600,
    "p99": 63464700
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



### results\set1\OCC_skew_1.2_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 41269,
  "committed_txns": 29932,
  "client_visible_failures": 11337,
  "internal_retries": 64348,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 997.7333333333333,
  "retries_per_request": 1.5592333228331192,
  "client_failure_rate_pct": 27.470983062347038,
  "realized_hot_row_hit_rate_pct": 20.78250982013645,
  "dead_tuples_generated": 69883,
  "dead_tuples_per_commit": 2.334725377522384,
  "app_latency": {
    "p50": 0,
    "p95": 571200,
    "p99": 1116800
  },
  "db_latency": {
    "p50": 3922100,
    "p95": 6734200,
    "p99": 10017400
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 530200,
    "p95": 1523300,
    "p99": 2320900
  },
  "client_e2e_latency": {
    "p50": 4726800,
    "p95": 61838600,
    "p99": 98875000
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



### results\set1\OCC_skew_1_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 59978,
  "committed_txns": 51364,
  "client_visible_failures": 8614,
  "internal_retries": 57051,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 1712.1333333333334,
  "retries_per_request": 0.9511987728833906,
  "client_failure_rate_pct": 14.361932708659841,
  "realized_hot_row_hit_rate_pct": 10.251294336957233,
  "dead_tuples_generated": 116993,
  "dead_tuples_per_commit": 2.277723697531345,
  "app_latency": {
    "p50": 0,
    "p95": 1011500,
    "p99": 1584800
  },
  "db_latency": {
    "p50": 4736800,
    "p95": 8748400,
    "p99": 13227800
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 535400,
    "p95": 2059900,
    "p99": 3644600
  },
  "client_e2e_latency": {
    "p50": 5589700,
    "p95": 47029500,
    "p99": 93744900
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



### results\set1\PESSIMISTIC_skew_0.4_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 96363,
  "committed_txns": 96363,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3212.1,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.2327910688723125,
  "dead_tuples_generated": 216016,
  "dead_tuples_per_commit": 2.2416902753131387,
  "app_latency": {
    "p50": 993600,
    "p95": 4520400,
    "p99": 9829200
  },
  "db_latency": {
    "p50": 9919600,
    "p95": 22578100,
    "p99": 40366600
  },
  "cc_wait_latency": {
    "p50": 1908700,
    "p95": 9283300,
    "p99": 20442500
  },
  "wal_wait_proxy": {
    "p50": 2000300,
    "p95": 11524800,
    "p99": 19413800
  },
  "client_e2e_latency": {
    "p50": 13434600,
    "p95": 30769500,
    "p99": 56739000
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



### results\set1\PESSIMISTIC_skew_0.6_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 103573,
  "committed_txns": 103573,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3452.4333333333334,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 1.0007668381282258,
  "dead_tuples_generated": 228492,
  "dead_tuples_per_commit": 2.206096183368252,
  "app_latency": {
    "p50": 615400,
    "p95": 3619900,
    "p99": 7579700
  },
  "db_latency": {
    "p50": 8230900,
    "p95": 18059600,
    "p99": 31402800
  },
  "cc_wait_latency": {
    "p50": 1522800,
    "p95": 11212400,
    "p99": 39077200
  },
  "wal_wait_proxy": {
    "p50": 1654500,
    "p95": 9522800,
    "p99": 15053900
  },
  "client_e2e_latency": {
    "p50": 11492100,
    "p95": 28802500,
    "p99": 68529200
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



### results\set1\PESSIMISTIC_skew_0.8_conc_10.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 76040,
  "committed_txns": 76031,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 9,
  "throughput_tps": 2534.366666666667,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.635706872925136,
  "dead_tuples_generated": 170210,
  "dead_tuples_per_commit": 2.2386921124278256,
  "app_latency": {
    "p50": 0,
    "p95": 666700,
    "p99": 1009100
  },
  "db_latency": {
    "p50": 2724400,
    "p95": 3338800,
    "p99": 4009300
  },
  "cc_wait_latency": {
    "p50": 629600,
    "p95": 2793700,
    "p99": 5583100
  },
  "wal_wait_proxy": {
    "p50": 559800,
    "p95": 1022700,
    "p99": 1369200
  },
  "client_e2e_latency": {
    "p50": 3567500,
    "p95": 5744800,
    "p99": 8822900
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



### results\set1\PESSIMISTIC_skew_0.8_conc_100.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 88556,
  "committed_txns": 88546,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 10,
  "throughput_tps": 2951.5333333333333,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.7418343871486104,
  "dead_tuples_generated": 196344,
  "dead_tuples_per_commit": 2.21742371196892,
  "app_latency": {
    "p50": 0,
    "p95": 1011700,
    "p99": 1660400
  },
  "db_latency": {
    "p50": 3012400,
    "p95": 5003700,
    "p99": 7511100
  },
  "cc_wait_latency": {
    "p50": 1001200,
    "p95": 67118600,
    "p99": 908889800
  },
  "wal_wait_proxy": {
    "p50": 996300,
    "p95": 1997800,
    "p99": 2999300
  },
  "client_e2e_latency": {
    "p50": 4534600,
    "p95": 71720100,
    "p99": 916277600
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



### results\set1\PESSIMISTIC_skew_0.8_conc_250.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 47284,
  "committed_txns": 45000,
  "client_visible_failures": 2281,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 3,
  "throughput_tps": 1500,
  "retries_per_request": 0,
  "client_failure_rate_pct": 4.8240419592251085,
  "realized_hot_row_hit_rate_pct": 3.6857339765678843,
  "dead_tuples_generated": 102526,
  "dead_tuples_per_commit": 2.2783555555555557,
  "app_latency": {
    "p50": 509400,
    "p95": 1999300,
    "p99": 5997700
  },
  "db_latency": {
    "p50": 4270900,
    "p95": 14977400,
    "p99": 35562100
  },
  "cc_wait_latency": {
    "p50": 1149800,
    "p95": 57402600,
    "p99": 2045281200
  },
  "wal_wait_proxy": {
    "p50": 999200,
    "p95": 3509000,
    "p99": 10518700
  },
  "client_e2e_latency": {
    "p50": 6526500,
    "p95": 78536400,
    "p99": 2064673100
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



### results\set1\PESSIMISTIC_skew_0.8_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 82207,
  "committed_txns": 82207,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2740.233333333333,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.6810985067042155,
  "dead_tuples_generated": 185790,
  "dead_tuples_per_commit": 2.2600265184230053,
  "app_latency": {
    "p50": 0,
    "p95": 1020000,
    "p99": 2012000
  },
  "db_latency": {
    "p50": 3535600,
    "p95": 6190300,
    "p99": 10519400
  },
  "cc_wait_latency": {
    "p50": 1003100,
    "p95": 22046300,
    "p99": 391621800
  },
  "wal_wait_proxy": {
    "p50": 998500,
    "p95": 2002100,
    "p99": 3565300
  },
  "client_e2e_latency": {
    "p50": 5512300,
    "p95": 27635500,
    "p99": 396192200
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



### results\set1\PESSIMISTIC_skew_0_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 96291,
  "committed_txns": 96291,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3209.7,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.013078910932616549,
  "dead_tuples_generated": 218836,
  "dead_tuples_per_commit": 2.2726526882055436,
  "app_latency": {
    "p50": 996800,
    "p95": 4518400,
    "p99": 10517100
  },
  "db_latency": {
    "p50": 9975400,
    "p95": 22778100,
    "p99": 42548000
  },
  "cc_wait_latency": {
    "p50": 1995200,
    "p95": 8693500,
    "p99": 18568200
  },
  "wal_wait_proxy": {
    "p50": 2001300,
    "p95": 11585800,
    "p99": 19178000
  },
  "client_e2e_latency": {
    "p50": 13287400,
    "p95": 30573900,
    "p99": 59376500
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



### results\set1\PESSIMISTIC_skew_1.2_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 27159,
  "committed_txns": 26319,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 840,
  "throughput_tps": 877.3,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 20.831781502172564,
  "dead_tuples_generated": 59470,
  "dead_tuples_per_commit": 2.259584330711653,
  "app_latency": {
    "p50": 0,
    "p95": 680000,
    "p99": 780500
  },
  "db_latency": {
    "p50": 2336900,
    "p95": 3323700,
    "p99": 4593400
  },
  "cc_wait_latency": {
    "p50": 1190500,
    "p95": 281763700,
    "p99": 595807600
  },
  "wal_wait_proxy": {
    "p50": 557900,
    "p95": 741200,
    "p99": 1297000
  },
  "client_e2e_latency": {
    "p50": 4362500,
    "p95": 297989500,
    "p99": 613347900
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



### results\set1\PESSIMISTIC_skew_1_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 44032,
  "committed_txns": 43988,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 44,
  "throughput_tps": 1466.2666666666667,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 10.244236747747838,
  "dead_tuples_generated": 98048,
  "dead_tuples_per_commit": 2.228971537692098,
  "app_latency": {
    "p50": 0,
    "p95": 659900,
    "p99": 851500
  },
  "db_latency": {
    "p50": 2763400,
    "p95": 3606800,
    "p99": 4923200
  },
  "cc_wait_latency": {
    "p50": 1109500,
    "p95": 207679900,
    "p99": 554588500
  },
  "wal_wait_proxy": {
    "p50": 562700,
    "p95": 1116200,
    "p99": 1775300
  },
  "client_e2e_latency": {
    "p50": 4024000,
    "p95": 211748500,
    "p99": 559014300
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



### results\set1\SSI_skew_0.4_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 83171,
  "committed_txns": 82705,
  "client_visible_failures": 466,
  "internal_retries": 1989,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2756.8333333333335,
  "retries_per_request": 0.023914585612773685,
  "client_failure_rate_pct": 0.5602914477401979,
  "realized_hot_row_hit_rate_pct": 0.2288575004469873,
  "dead_tuples_generated": 194525,
  "dead_tuples_per_commit": 2.352034338915422,
  "app_latency": {
    "p50": 0,
    "p95": 1777300,
    "p99": 8077300
  },
  "db_latency": {
    "p50": 12605100,
    "p95": 39268100,
    "p99": 77641500
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2518000,
    "p95": 15687900,
    "p99": 35168600
  },
  "client_e2e_latency": {
    "p50": 13526200,
    "p95": 43556200,
    "p99": 85987300
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



### results\set1\SSI_skew_0.6_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 114437,
  "committed_txns": 113174,
  "client_visible_failures": 1263,
  "internal_retries": 11779,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3772.4666666666667,
  "retries_per_request": 0.10292999641724268,
  "client_failure_rate_pct": 1.103664024747241,
  "realized_hot_row_hit_rate_pct": 1.0796581414427833,
  "dead_tuples_generated": 247391,
  "dead_tuples_per_commit": 2.185934932051531,
  "app_latency": {
    "p50": 0,
    "p95": 1510300,
    "p99": 5121600
  },
  "db_latency": {
    "p50": 8886000,
    "p95": 21253300,
    "p99": 37602100
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1605600,
    "p95": 10733500,
    "p99": 18834600
  },
  "client_e2e_latency": {
    "p50": 9693400,
    "p95": 28109000,
    "p99": 66646500
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



### results\set1\SSI_skew_0.8_conc_10.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 73108,
  "committed_txns": 72951,
  "client_visible_failures": 157,
  "internal_retries": 8536,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2431.7,
  "retries_per_request": 0.11675876785030366,
  "client_failure_rate_pct": 0.21475077966843573,
  "realized_hot_row_hit_rate_pct": 3.697886921758995,
  "dead_tuples_generated": 164969,
  "dead_tuples_per_commit": 2.2613672190922673,
  "app_latency": {
    "p50": 0,
    "p95": 586900,
    "p99": 671800
  },
  "db_latency": {
    "p50": 2845300,
    "p95": 3497700,
    "p99": 4118600
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 552300,
    "p95": 991500,
    "p99": 1215300
  },
  "client_e2e_latency": {
    "p50": 2961700,
    "p95": 8898300,
    "p99": 21301000
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



### results\set1\SSI_skew_0.8_conc_100.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 101303,
  "committed_txns": 92017,
  "client_visible_failures": 9286,
  "internal_retries": 56975,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3067.233333333333,
  "retries_per_request": 0.5624216459532294,
  "client_failure_rate_pct": 9.166559726760314,
  "realized_hot_row_hit_rate_pct": 3.6391535461632145,
  "dead_tuples_generated": 208374,
  "dead_tuples_per_commit": 2.2645163393720726,
  "app_latency": {
    "p50": 0,
    "p95": 1509700,
    "p99": 6225600
  },
  "db_latency": {
    "p50": 8932500,
    "p95": 31457900,
    "p99": 61601700
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1592000,
    "p95": 20638200,
    "p99": 51159800
  },
  "client_e2e_latency": {
    "p50": 10925900,
    "p95": 66244300,
    "p99": 136182400
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



### results\set1\SSI_skew_0.8_conc_250.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 72100,
  "committed_txns": 64535,
  "client_visible_failures": 7565,
  "internal_retries": 36568,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2151.1666666666665,
  "retries_per_request": 0.5071844660194175,
  "client_failure_rate_pct": 10.492371705963938,
  "realized_hot_row_hit_rate_pct": 3.748303622418589,
  "dead_tuples_generated": 148185,
  "dead_tuples_per_commit": 2.296195862710157,
  "app_latency": {
    "p50": 0,
    "p95": 1537300,
    "p99": 8244200
  },
  "db_latency": {
    "p50": 9309300,
    "p95": 57404600,
    "p99": 121719200
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1612900,
    "p95": 23125700,
    "p99": 86560900
  },
  "client_e2e_latency": {
    "p50": 11226700,
    "p95": 105723000,
    "p99": 235838700
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



### results\set1\SSI_skew_0.8_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 84283,
  "committed_txns": 79656,
  "client_visible_failures": 4627,
  "internal_retries": 36127,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2655.2,
  "retries_per_request": 0.428639227364949,
  "client_failure_rate_pct": 5.489837808336201,
  "realized_hot_row_hit_rate_pct": 3.6798958591085627,
  "dead_tuples_generated": 173695,
  "dead_tuples_per_commit": 2.180563924876971,
  "app_latency": {
    "p50": 0,
    "p95": 1509300,
    "p99": 4959700
  },
  "db_latency": {
    "p50": 6640200,
    "p95": 19162500,
    "p99": 35700000
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1065600,
    "p95": 7114700,
    "p99": 15093200
  },
  "client_e2e_latency": {
    "p50": 7734200,
    "p95": 42251100,
    "p99": 96905900
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



### results\set1\SSI_skew_0_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 94204,
  "committed_txns": 93622,
  "client_visible_failures": 582,
  "internal_retries": 1332,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3120.733333333333,
  "retries_per_request": 0.014139526983992188,
  "client_failure_rate_pct": 0.6178081610122712,
  "realized_hot_row_hit_rate_pct": 0.01270546382001534,
  "dead_tuples_generated": 210296,
  "dead_tuples_per_commit": 2.246224178077802,
  "app_latency": {
    "p50": 0,
    "p95": 1633900,
    "p99": 8064900
  },
  "db_latency": {
    "p50": 12081300,
    "p95": 31975500,
    "p99": 58920600
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2203100,
    "p95": 14269500,
    "p99": 26916200
  },
  "client_e2e_latency": {
    "p50": 12684600,
    "p95": 34886200,
    "p99": 66501900
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



### results\set1\SSI_skew_1.2_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 55959,
  "committed_txns": 39351,
  "client_visible_failures": 13467,
  "internal_retries": 79538,
  "deadlock_count": 0,
  "insufficient_funds_count": 3141,
  "throughput_tps": 1311.7,
  "retries_per_request": 1.42136206865741,
  "client_failure_rate_pct": 24.065833914115693,
  "realized_hot_row_hit_rate_pct": 20.940799656720763,
  "dead_tuples_generated": 92578,
  "dead_tuples_per_commit": 2.35262128027242,
  "app_latency": {
    "p50": 0,
    "p95": 562300,
    "p99": 1133600
  },
  "db_latency": {
    "p50": 3232400,
    "p95": 5500500,
    "p99": 8638800
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 525500,
    "p95": 1518800,
    "p99": 2608500
  },
  "client_e2e_latency": {
    "p50": 4140300,
    "p95": 59048600,
    "p99": 79932800
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



### results\set1\SSI_skew_1_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 75548,
  "committed_txns": 63859,
  "client_visible_failures": 10954,
  "internal_retries": 72907,
  "deadlock_count": 0,
  "insufficient_funds_count": 735,
  "throughput_tps": 2128.633333333333,
  "retries_per_request": 0.9650420924445385,
  "client_failure_rate_pct": 14.499391115582148,
  "realized_hot_row_hit_rate_pct": 10.311132633472269,
  "dead_tuples_generated": 147699,
  "dead_tuples_per_commit": 2.3128924662146293,
  "app_latency": {
    "p50": 0,
    "p95": 1004600,
    "p99": 1573800
  },
  "db_latency": {
    "p50": 4104800,
    "p95": 7600700,
    "p99": 12390600
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 532900,
    "p95": 2058700,
    "p99": 3883100
  },
  "client_e2e_latency": {
    "p50": 4778900,
    "p95": 42340900,
    "p99": 80735400
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



### results\set1_repetitions\OCC_skew_0.6_run_1.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 138868,
  "committed_txns": 138378,
  "client_visible_failures": 490,
  "collision_rejections": 0,
  "internal_retries": 13431,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4612.6,
  "retries_per_request": 0.09671774634905089,
  "client_failure_rate_pct": 0.35285306910159286,
  "realized_hot_row_hit_rate_pct": 1.0323015618092946,
  "dead_tuples_generated": 311236,
  "dead_tuples_per_commit": 2.249172556331209,
  "app_latency": {
    "p50": 0,
    "p95": 1195200,
    "p99": 5234700
  },
  "db_latency": {
    "p50": 7596000,
    "p95": 16915400,
    "p99": 28605800
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1058400,
    "p95": 8087600,
    "p99": 14141200
  },
  "client_e2e_latency": {
    "p50": 8166000,
    "p95": 22594600,
    "p99": 56925000
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



### results\set1_repetitions\OCC_skew_0.6_run_2.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 141438,
  "committed_txns": 140893,
  "client_visible_failures": 545,
  "collision_rejections": 0,
  "internal_retries": 13906,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4696.433333333333,
  "retries_per_request": 0.0983186979453895,
  "client_failure_rate_pct": 0.3853278468304133,
  "realized_hot_row_hit_rate_pct": 1.0433767236589753,
  "dead_tuples_generated": 314170,
  "dead_tuples_per_commit": 2.2298481826634395,
  "app_latency": {
    "p50": 0,
    "p95": 1199600,
    "p99": 5153100
  },
  "db_latency": {
    "p50": 7550700,
    "p95": 16334600,
    "p99": 27177400
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1060900,
    "p95": 7649900,
    "p99": 14072200
  },
  "client_e2e_latency": {
    "p50": 8128700,
    "p95": 21823200,
    "p99": 55788700
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



### results\set1_repetitions\OCC_skew_0.6_run_3.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 140985,
  "committed_txns": 140477,
  "client_visible_failures": 508,
  "collision_rejections": 0,
  "internal_retries": 13514,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4682.566666666667,
  "retries_per_request": 0.09585416888321453,
  "client_failure_rate_pct": 0.3603220200730574,
  "realized_hot_row_hit_rate_pct": 1.0224910520310717,
  "dead_tuples_generated": 316512,
  "dead_tuples_per_commit": 2.2531232870861424,
  "app_latency": {
    "p50": 0,
    "p95": 1072300,
    "p99": 5066600
  },
  "db_latency": {
    "p50": 7612700,
    "p95": 16293500,
    "p99": 26617900
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1085400,
    "p95": 7908600,
    "p99": 13874200
  },
  "client_e2e_latency": {
    "p50": 8175000,
    "p95": 21706500,
    "p99": 54629100
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



### results\set1_repetitions\OCC_skew_0.8_run_1.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 110501,
  "committed_txns": 105983,
  "client_visible_failures": 4518,
  "collision_rejections": 0,
  "internal_retries": 41698,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3532.766666666667,
  "retries_per_request": 0.3773540510945602,
  "client_failure_rate_pct": 4.088650781440892,
  "realized_hot_row_hit_rate_pct": 3.717837736784566,
  "dead_tuples_generated": 234148,
  "dead_tuples_per_commit": 2.2092977175584765,
  "app_latency": {
    "p50": 0,
    "p95": 1009400,
    "p99": 3291400
  },
  "db_latency": {
    "p50": 4682500,
    "p95": 10896700,
    "p99": 17975400
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 557800,
    "p95": 4025900,
    "p99": 8463000
  },
  "client_e2e_latency": {
    "p50": 5357500,
    "p95": 28204400,
    "p99": 86640700
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



### results\set1_repetitions\OCC_skew_0.8_run_2.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 109833,
  "committed_txns": 105444,
  "client_visible_failures": 4389,
  "collision_rejections": 0,
  "internal_retries": 41455,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3514.8,
  "retries_per_request": 0.37743665382899494,
  "client_failure_rate_pct": 3.996066755893038,
  "realized_hot_row_hit_rate_pct": 3.657164441801217,
  "dead_tuples_generated": 237529,
  "dead_tuples_per_commit": 2.2526554379575887,
  "app_latency": {
    "p50": 0,
    "p95": 1009900,
    "p99": 3377100
  },
  "db_latency": {
    "p50": 4702800,
    "p95": 10938500,
    "p99": 18583200
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 550200,
    "p95": 4077200,
    "p99": 8366600
  },
  "client_e2e_latency": {
    "p50": 5437700,
    "p95": 29521900,
    "p99": 86957500
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



### results\set1_repetitions\OCC_skew_0.8_run_3.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 110236,
  "committed_txns": 105724,
  "client_visible_failures": 4512,
  "collision_rejections": 0,
  "internal_retries": 41664,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3524.133333333333,
  "retries_per_request": 0.3779527559055118,
  "client_failure_rate_pct": 4.0930367575020865,
  "realized_hot_row_hit_rate_pct": 3.689787089150159,
  "dead_tuples_generated": 237015,
  "dead_tuples_per_commit": 2.2418277779879685,
  "app_latency": {
    "p50": 0,
    "p95": 1010300,
    "p99": 3324800
  },
  "db_latency": {
    "p50": 4641300,
    "p95": 10729300,
    "p99": 18205200
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 553200,
    "p95": 3840900,
    "p99": 8023000
  },
  "client_e2e_latency": {
    "p50": 5296700,
    "p95": 28482400,
    "p99": 85602300
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



### results\set1_repetitions\OCC_skew_1_run_1.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 95044,
  "committed_txns": 83387,
  "client_visible_failures": 11550,
  "collision_rejections": 0,
  "internal_retries": 82082,
  "deadlock_count": 0,
  "insufficient_funds_count": 107,
  "throughput_tps": 2779.5666666666666,
  "retries_per_request": 0.8636210597197088,
  "client_failure_rate_pct": 12.152266318757627,
  "realized_hot_row_hit_rate_pct": 10.135741632929712,
  "dead_tuples_generated": 189790,
  "dead_tuples_per_commit": 2.276014246825045,
  "app_latency": {
    "p50": 0,
    "p95": 531200,
    "p99": 1014900
  },
  "db_latency": {
    "p50": 2782100,
    "p95": 4324300,
    "p99": 5701000
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 516900,
    "p95": 1080200,
    "p99": 2077100
  },
  "client_e2e_latency": {
    "p50": 3172300,
    "p95": 36210100,
    "p99": 74067000
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



### results\set1_repetitions\OCC_skew_1_run_2.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 95040,
  "committed_txns": 83464,
  "client_visible_failures": 11576,
  "collision_rejections": 0,
  "internal_retries": 82446,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 2782.133333333333,
  "retries_per_request": 0.8674873737373737,
  "client_failure_rate_pct": 12.18013468013468,
  "realized_hot_row_hit_rate_pct": 10.239377508754927,
  "dead_tuples_generated": 190154,
  "dead_tuples_per_commit": 2.2782756637592256,
  "app_latency": {
    "p50": 0,
    "p95": 530800,
    "p99": 1014100
  },
  "db_latency": {
    "p50": 2752700,
    "p95": 4292600,
    "p99": 5680200
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 516800,
    "p95": 1089100,
    "p99": 2101500
  },
  "client_e2e_latency": {
    "p50": 3169900,
    "p95": 35885200,
    "p99": 73306000
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



### results\set1_repetitions\OCC_skew_1_run_3.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 94257,
  "committed_txns": 82666,
  "client_visible_failures": 11530,
  "collision_rejections": 0,
  "internal_retries": 81787,
  "deadlock_count": 0,
  "insufficient_funds_count": 61,
  "throughput_tps": 2755.5333333333333,
  "retries_per_request": 0.8677021335285443,
  "client_failure_rate_pct": 12.232513235091293,
  "realized_hot_row_hit_rate_pct": 10.195599585301233,
  "dead_tuples_generated": 189048,
  "dead_tuples_per_commit": 2.286889410398471,
  "app_latency": {
    "p50": 0,
    "p95": 531900,
    "p99": 1023600
  },
  "db_latency": {
    "p50": 2843900,
    "p95": 4428500,
    "p99": 5874300
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 517100,
    "p95": 1104200,
    "p99": 2110000
  },
  "client_e2e_latency": {
    "p50": 3187400,
    "p95": 35971300,
    "p99": 74186000
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



### results\set1_repetitions\PESSIMISTIC_skew_0.6_run_1.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 136938,
  "committed_txns": 136938,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4564.6,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 1.0226902262502444,
  "dead_tuples_generated": 302028,
  "dead_tuples_per_commit": 2.205582088244315,
  "app_latency": {
    "p50": 539500,
    "p95": 2524200,
    "p99": 6522700
  },
  "db_latency": {
    "p50": 5998000,
    "p95": 14517800,
    "p99": 28374800
  },
  "cc_wait_latency": {
    "p50": 1110000,
    "p95": 9040400,
    "p99": 30552700
  },
  "wal_wait_proxy": {
    "p50": 1095200,
    "p95": 8000400,
    "p99": 14711900
  },
  "client_e2e_latency": {
    "p50": 8522800,
    "p95": 22284600,
    "p99": 59044600
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



### results\set1_repetitions\PESSIMISTIC_skew_0.6_run_2.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 137970,
  "committed_txns": 137970,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4599,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 1.0187618271781818,
  "dead_tuples_generated": 311162,
  "dead_tuples_per_commit": 2.2552873813147785,
  "app_latency": {
    "p50": 537000,
    "p95": 2507300,
    "p99": 6522500
  },
  "db_latency": {
    "p50": 5983900,
    "p95": 14882000,
    "p99": 24163500
  },
  "cc_wait_latency": {
    "p50": 1103200,
    "p95": 9138100,
    "p99": 30068400
  },
  "wal_wait_proxy": {
    "p50": 1084900,
    "p95": 8117600,
    "p99": 14076600
  },
  "client_e2e_latency": {
    "p50": 8580900,
    "p95": 22119900,
    "p99": 50639500
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



### results\set1_repetitions\PESSIMISTIC_skew_0.6_run_3.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 145258,
  "committed_txns": 145258,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4841.933333333333,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 1.0200923568147624,
  "dead_tuples_generated": 322100,
  "dead_tuples_per_commit": 2.217433807432293,
  "app_latency": {
    "p50": 536500,
    "p95": 2083100,
    "p99": 6407000
  },
  "db_latency": {
    "p50": 5559500,
    "p95": 14026400,
    "p99": 22938800
  },
  "cc_wait_latency": {
    "p50": 1081100,
    "p95": 8732500,
    "p99": 28072400
  },
  "wal_wait_proxy": {
    "p50": 1009900,
    "p95": 7577500,
    "p99": 13248100
  },
  "client_e2e_latency": {
    "p50": 8348500,
    "p95": 20732400,
    "p99": 48098100
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



### results\set1_repetitions\PESSIMISTIC_skew_0.8_run_1.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 89232,
  "committed_txns": 89215,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 17,
  "throughput_tps": 2973.8333333333335,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.6923496815007235,
  "dead_tuples_generated": 201238,
  "dead_tuples_per_commit": 2.255652076444544,
  "app_latency": {
    "p50": 0,
    "p95": 1369100,
    "p99": 2452100
  },
  "db_latency": {
    "p50": 3010300,
    "p95": 6024400,
    "p99": 12350800
  },
  "cc_wait_latency": {
    "p50": 1008700,
    "p95": 20014900,
    "p99": 343823400
  },
  "wal_wait_proxy": {
    "p50": 618700,
    "p95": 2209200,
    "p99": 4199200
  },
  "client_e2e_latency": {
    "p50": 4527300,
    "p95": 26855800,
    "p99": 349184300
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



### results\set1_repetitions\PESSIMISTIC_skew_0.8_run_2.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 96219,
  "committed_txns": 96219,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3207.3,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.6837967787345294,
  "dead_tuples_generated": 214206,
  "dead_tuples_per_commit": 2.2262339039067127,
  "app_latency": {
    "p50": 0,
    "p95": 1328400,
    "p99": 2275600
  },
  "db_latency": {
    "p50": 2999600,
    "p95": 5634100,
    "p99": 10547000
  },
  "cc_wait_latency": {
    "p50": 1007100,
    "p95": 18964200,
    "p99": 296211100
  },
  "wal_wait_proxy": {
    "p50": 612100,
    "p95": 2031000,
    "p99": 3796900
  },
  "client_e2e_latency": {
    "p50": 4368400,
    "p95": 24793000,
    "p99": 300846000
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



### results\set1_repetitions\PESSIMISTIC_skew_0.8_run_3.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 103138,
  "committed_txns": 103036,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 102,
  "throughput_tps": 3434.5333333333333,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.721062771934451,
  "dead_tuples_generated": 230220,
  "dead_tuples_per_commit": 2.234364688070189,
  "app_latency": {
    "p50": 0,
    "p95": 1257500,
    "p99": 2001300
  },
  "db_latency": {
    "p50": 2884500,
    "p95": 4896400,
    "p99": 8229400
  },
  "cc_wait_latency": {
    "p50": 1002300,
    "p95": 18948500,
    "p99": 288626700
  },
  "wal_wait_proxy": {
    "p50": 573900,
    "p95": 1886700,
    "p99": 3206400
  },
  "client_e2e_latency": {
    "p50": 4094100,
    "p95": 23830900,
    "p99": 294650300
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



### results\set1_repetitions\PESSIMISTIC_skew_1_run_1.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 66254,
  "committed_txns": 65980,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 274,
  "throughput_tps": 2199.3333333333335,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 10.265068488642273,
  "dead_tuples_generated": 147858,
  "dead_tuples_per_commit": 2.2409518035768414,
  "app_latency": {
    "p50": 0,
    "p95": 591100,
    "p99": 646200
  },
  "db_latency": {
    "p50": 1698700,
    "p95": 2307100,
    "p99": 2861900
  },
  "cc_wait_latency": {
    "p50": 569200,
    "p95": 132382800,
    "p99": 357077100
  },
  "wal_wait_proxy": {
    "p50": 535700,
    "p95": 625000,
    "p99": 1173300
  },
  "client_e2e_latency": {
    "p50": 2768000,
    "p95": 136505700,
    "p99": 362505800
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



### results\set1_repetitions\PESSIMISTIC_skew_1_run_2.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 65559,
  "committed_txns": 65018,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 541,
  "throughput_tps": 2167.266666666667,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 10.14563042293945,
  "dead_tuples_generated": 145508,
  "dead_tuples_per_commit": 2.2379648712664184,
  "app_latency": {
    "p50": 0,
    "p95": 592500,
    "p99": 646700
  },
  "db_latency": {
    "p50": 1703900,
    "p95": 2319200,
    "p99": 2894200
  },
  "cc_wait_latency": {
    "p50": 569700,
    "p95": 128837200,
    "p99": 351374800
  },
  "wal_wait_proxy": {
    "p50": 536200,
    "p95": 623800,
    "p99": 1187800
  },
  "client_e2e_latency": {
    "p50": 2773300,
    "p95": 135615800,
    "p99": 360263700
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



### results\set1_repetitions\PESSIMISTIC_skew_1_run_3.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 65981,
  "committed_txns": 65540,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 441,
  "throughput_tps": 2184.6666666666665,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 10.22138052635016,
  "dead_tuples_generated": 148340,
  "dead_tuples_per_commit": 2.2633506255721696,
  "app_latency": {
    "p50": 0,
    "p95": 591400,
    "p99": 651800
  },
  "db_latency": {
    "p50": 1695300,
    "p95": 2311000,
    "p99": 2888900
  },
  "cc_wait_latency": {
    "p50": 567100,
    "p95": 126917300,
    "p99": 351823700
  },
  "wal_wait_proxy": {
    "p50": 535200,
    "p95": 623100,
    "p99": 1236600
  },
  "client_e2e_latency": {
    "p50": 2763600,
    "p95": 132628600,
    "p99": 361257600
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



### results\set1_repetitions\SSI_skew_0.6_run_1.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 110550,
  "committed_txns": 110122,
  "client_visible_failures": 428,
  "collision_rejections": 0,
  "internal_retries": 11940,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3670.733333333333,
  "retries_per_request": 0.10800542740841249,
  "client_failure_rate_pct": 0.3871551334237901,
  "realized_hot_row_hit_rate_pct": 1.0202952709928825,
  "dead_tuples_generated": 244218,
  "dead_tuples_per_commit": 2.217704001017054,
  "app_latency": {
    "p50": 0,
    "p95": 1511000,
    "p99": 5739700
  },
  "db_latency": {
    "p50": 9278600,
    "p95": 21911500,
    "p99": 36516800
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 2005300,
    "p95": 11095000,
    "p99": 19105400
  },
  "client_e2e_latency": {
    "p50": 10182800,
    "p95": 29375000,
    "p99": 68162500
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



### results\set1_repetitions\SSI_skew_0.6_run_2.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 117285,
  "committed_txns": 116846,
  "client_visible_failures": 439,
  "collision_rejections": 0,
  "internal_retries": 12343,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 3894.866666666667,
  "retries_per_request": 0.10523937417402054,
  "client_failure_rate_pct": 0.37430191414076824,
  "realized_hot_row_hit_rate_pct": 1.0059459398317643,
  "dead_tuples_generated": 266620,
  "dead_tuples_per_commit": 2.281806822655461,
  "app_latency": {
    "p50": 0,
    "p95": 1509400,
    "p99": 5159900
  },
  "db_latency": {
    "p50": 9139700,
    "p95": 20225600,
    "p99": 32062700
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1760000,
    "p95": 10726600,
    "p99": 18479100
  },
  "client_e2e_latency": {
    "p50": 9827700,
    "p95": 26972700,
    "p99": 61761600
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



### results\set1_repetitions\SSI_skew_0.6_run_3.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 135356,
  "committed_txns": 134863,
  "client_visible_failures": 493,
  "collision_rejections": 0,
  "internal_retries": 13989,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4495.433333333333,
  "retries_per_request": 0.10334968527438754,
  "client_failure_rate_pct": 0.3642247111321256,
  "realized_hot_row_hit_rate_pct": 1.0250603613354425,
  "dead_tuples_generated": 298612,
  "dead_tuples_per_commit": 2.214187731253198,
  "app_latency": {
    "p50": 0,
    "p95": 1383400,
    "p99": 5421300
  },
  "db_latency": {
    "p50": 7124600,
    "p95": 18923000,
    "p99": 36092100
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1236800,
    "p95": 9639200,
    "p99": 19872600
  },
  "client_e2e_latency": {
    "p50": 7794900,
    "p95": 25517100,
    "p99": 64739900
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



### results\set1_repetitions\SSI_skew_0.8_run_1.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 147580,
  "committed_txns": 141488,
  "client_visible_failures": 6092,
  "collision_rejections": 0,
  "internal_retries": 59073,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4716.266666666666,
  "retries_per_request": 0.4002778154221439,
  "client_failure_rate_pct": 4.127930613904323,
  "realized_hot_row_hit_rate_pct": 3.6923913536049304,
  "dead_tuples_generated": 316334,
  "dead_tuples_per_commit": 2.235765577292774,
  "app_latency": {
    "p50": 0,
    "p95": 1008000,
    "p99": 2066400
  },
  "db_latency": {
    "p50": 3662800,
    "p95": 7702600,
    "p99": 13260100
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 532300,
    "p95": 2655700,
    "p99": 5695000
  },
  "client_e2e_latency": {
    "p50": 4195300,
    "p95": 22880500,
    "p99": 69918600
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



### results\set1_repetitions\SSI_skew_0.8_run_2.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 132757,
  "committed_txns": 127024,
  "client_visible_failures": 5733,
  "collision_rejections": 0,
  "internal_retries": 54052,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 4234.133333333333,
  "retries_per_request": 0.4071499054663784,
  "client_failure_rate_pct": 4.318416354693161,
  "realized_hot_row_hit_rate_pct": 3.6151412318234724,
  "dead_tuples_generated": 294390,
  "dead_tuples_per_commit": 2.3175935256329514,
  "app_latency": {
    "p50": 0,
    "p95": 1008900,
    "p99": 2598100
  },
  "db_latency": {
    "p50": 4106000,
    "p95": 9273900,
    "p99": 15918700
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 557400,
    "p95": 3598000,
    "p99": 7630300
  },
  "client_e2e_latency": {
    "p50": 4690500,
    "p95": 25267700,
    "p99": 73771100
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



### results\set1_repetitions\SSI_skew_0.8_run_3.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 113749,
  "committed_txns": 108541,
  "client_visible_failures": 5148,
  "collision_rejections": 0,
  "internal_retries": 47166,
  "deadlock_count": 0,
  "insufficient_funds_count": 60,
  "throughput_tps": 3618.0333333333333,
  "retries_per_request": 0.41464979911911315,
  "client_failure_rate_pct": 4.525754072563275,
  "realized_hot_row_hit_rate_pct": 3.685543154818821,
  "dead_tuples_generated": 250931,
  "dead_tuples_per_commit": 2.311854506591979,
  "app_latency": {
    "p50": 0,
    "p95": 1011300,
    "p99": 3779700
  },
  "db_latency": {
    "p50": 4631100,
    "p95": 11986000,
    "p99": 20844300
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 998200,
    "p95": 4639200,
    "p99": 9929000
  },
  "client_e2e_latency": {
    "p50": 5449300,
    "p95": 30265500,
    "p99": 81065200
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



### results\set1_repetitions\SSI_skew_1_run_1.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 96592,
  "committed_txns": 83841,
  "client_visible_failures": 11749,
  "collision_rejections": 0,
  "internal_retries": 84055,
  "deadlock_count": 0,
  "insufficient_funds_count": 1002,
  "throughput_tps": 2794.7,
  "retries_per_request": 0.870206642372039,
  "client_failure_rate_pct": 12.163533211860196,
  "realized_hot_row_hit_rate_pct": 10.241899539431921,
  "dead_tuples_generated": 192174,
  "dead_tuples_per_commit": 2.2921243782874727,
  "app_latency": {
    "p50": 0,
    "p95": 550300,
    "p99": 1078400
  },
  "db_latency": {
    "p50": 2675300,
    "p95": 4773100,
    "p99": 7466600
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 518800,
    "p95": 1519300,
    "p99": 2736900
  },
  "client_e2e_latency": {
    "p50": 3177300,
    "p95": 36057900,
    "p99": 72581500
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



### results\set1_repetitions\SSI_skew_1_run_2.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 96470,
  "committed_txns": 83842,
  "client_visible_failures": 11597,
  "collision_rejections": 0,
  "internal_retries": 83839,
  "deadlock_count": 0,
  "insufficient_funds_count": 1031,
  "throughput_tps": 2794.733333333333,
  "retries_per_request": 0.8690681040738053,
  "client_failure_rate_pct": 12.021353788742614,
  "realized_hot_row_hit_rate_pct": 10.069555286786294,
  "dead_tuples_generated": 192125,
  "dead_tuples_per_commit": 2.2915126070465877,
  "app_latency": {
    "p50": 0,
    "p95": 553700,
    "p99": 1180600
  },
  "db_latency": {
    "p50": 2649100,
    "p95": 5080200,
    "p99": 8546300
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 518500,
    "p95": 1518600,
    "p99": 2713400
  },
  "client_e2e_latency": {
    "p50": 3159700,
    "p95": 36261300,
    "p99": 73538600
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



### results\set1_repetitions\SSI_skew_1_run_3.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 104245,
  "committed_txns": 91013,
  "client_visible_failures": 12191,
  "collision_rejections": 0,
  "internal_retries": 88859,
  "deadlock_count": 0,
  "insufficient_funds_count": 1041,
  "throughput_tps": 3033.766666666667,
  "retries_per_request": 0.8524053911458583,
  "client_failure_rate_pct": 11.694565686603674,
  "realized_hot_row_hit_rate_pct": 10.247981642156041,
  "dead_tuples_generated": 207376,
  "dead_tuples_per_commit": 2.2785316383373804,
  "app_latency": {
    "p50": 0,
    "p95": 531700,
    "p99": 1014000
  },
  "db_latency": {
    "p50": 2592300,
    "p95": 3924600,
    "p99": 5190400
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 517400,
    "p95": 1110000,
    "p99": 2103800
  },
  "client_e2e_latency": {
    "p50": 2951200,
    "p95": 34470300,
    "p99": 70108800
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



### results\set2\Anomaly_OCC.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 22631,
  "committed_txns": 19162,
  "client_visible_failures": 3469,
  "internal_retries": 8978,
  "deadlock_count": 123,
  "throughput_tps": 638.7333333333333,
  "abort_retry_rate_pct": 54.99977906411559,
  "realized_hot_row_hit_rate_pct": 3.6310763547765945,
  "dead_tuples_generated": 11711,
  "dead_tuples_per_commit": 0.6111574992172008,
  "app_latency": {
    "p50": 0,
    "p95": 1534800,
    "p99": 7273200
  },
  "db_latency": {
    "p50": 8223600,
    "p95": 74690300,
    "p99": 512721900
  },
  "wal_wait_proxy": {
    "p50": 1039300,
    "p95": 7606900,
    "p99": 28588100
  },
  "client_e2e_latency": {
    "p50": 9775200,
    "p95": 377381500,
    "p99": 1463646800
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



### results\set2\Anomaly_PESSIMISTIC.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 32898,
  "committed_txns": 30689,
  "client_visible_failures": 2209,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 1022.9666666666667,
  "abort_retry_rate_pct": 6.714693902364885,
  "realized_hot_row_hit_rate_pct": 3.588198812793479,
  "dead_tuples_generated": 53230,
  "dead_tuples_per_commit": 1.7344977027599466,
  "app_latency": {
    "p50": 557900,
    "p95": 2716400,
    "p99": 9527100
  },
  "db_latency": {
    "p50": 7533700,
    "p95": 132858100,
    "p99": 3042338400
  },
  "wal_wait_proxy": {
    "p50": 1001600,
    "p95": 5048200,
    "p99": 16539500
  },
  "client_e2e_latency": {
    "p50": 8278500,
    "p95": 135671500,
    "p99": 3043337500
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



### results\set2\Audit_OCC.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 35755,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2\Audit_PESSIMISTIC.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 33032,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2\Audit_SSI.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 16186,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2\Set2_OCC_Idempotency.json

```json
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



### results\set2\Set2_PESSIMISTIC_Idempotency.json

```json
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



### results\set2\Set2_SSI_Idempotency.json

```json
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



### results\set2_investigation\Investigate_SSI_StageA.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 15036,
  "committed_txns": 4594,
  "client_visible_failures": 10442,
  "collision_rejections": 0,
  "internal_retries": 50317,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 153.13333333333333,
  "retries_per_request": 3.3464352221335463,
  "client_failure_rate_pct": 69.44666134610269,
  "realized_hot_row_hit_rate_pct": 3.3854243596939533,
  "dead_tuples_generated": 192391,
  "dead_tuples_per_commit": 41.878754897692644,
  "app_latency": {
    "p50": 0,
    "p95": 1092700,
    "p99": 6327300
  },
  "db_latency": {
    "p50": 5710500,
    "p95": 13062500,
    "p99": 20920100
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1058800,
    "p95": 6602600,
    "p99": 13155900
  },
  "client_e2e_latency": {
    "p50": 36883300,
    "p95": 115858200,
    "p99": 148200700
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



### results\set2AfterInvestigation\Audit_OCC.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 52720,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2AfterInvestigation\Audit_PESSIMISTIC.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 38916,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2AfterInvestigation\Audit_SSI.json

```json
{
  "initial_account_balance": 1000,
  "total_accounts_checked": 10000,
  "balance_discrepancy_count": 0,
  "total_committed_keys": 5683,
  "duplicate_effect_count": 0,
  "passed": true
}
```



### results\set2AfterInvestigation\Set2_OCC_Idempotency.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 47730,
  "committed_txns": 45904,
  "client_visible_failures": 1826,
  "collision_rejections": 0,
  "internal_retries": 17308,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 1530.1333333333334,
  "retries_per_request": 0.3626230882044836,
  "client_failure_rate_pct": 3.8256861512675466,
  "realized_hot_row_hit_rate_pct": 3.6949585194639436,
  "dead_tuples_generated": 210997,
  "dead_tuples_per_commit": 4.596483966538864,
  "app_latency": {
    "p50": 0,
    "p95": 1515200,
    "p99": 5690000
  },
  "db_latency": {
    "p50": 9607000,
    "p95": 22798800,
    "p99": 40731900
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1518200,
    "p95": 7780800,
    "p99": 12868100
  },
  "client_e2e_latency": {
    "p50": 17040400,
    "p95": 61656700,
    "p99": 154968500
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



### results\set2AfterInvestigation\Set2_PESSIMISTIC_Idempotency.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 33191,
  "committed_txns": 33191,
  "client_visible_failures": 0,
  "collision_rejections": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 1106.3666666666666,
  "retries_per_request": 0,
  "client_failure_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.7864600702906546,
  "dead_tuples_generated": 153126,
  "dead_tuples_per_commit": 4.613479557711428,
  "app_latency": {
    "p50": 0,
    "p95": 1513700,
    "p99": 2839300
  },
  "db_latency": {
    "p50": 9592000,
    "p95": 19174800,
    "p99": 27276400
  },
  "cc_wait_latency": {
    "p50": 1178500,
    "p95": 36178900,
    "p99": 1010480200
  },
  "wal_wait_proxy": {
    "p50": 1012000,
    "p95": 7042200,
    "p99": 14170700
  },
  "client_e2e_latency": {
    "p50": 15243400,
    "p95": 52171900,
    "p99": 1027715400
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



### results\set2AfterInvestigation\Set2_SSI_Idempotency.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 16222,
  "committed_txns": 5055,
  "client_visible_failures": 11167,
  "collision_rejections": 0,
  "internal_retries": 54009,
  "deadlock_count": 0,
  "insufficient_funds_count": 0,
  "throughput_tps": 168.5,
  "retries_per_request": 3.3293675255825423,
  "client_failure_rate_pct": 68.83861422759216,
  "realized_hot_row_hit_rate_pct": 3.648091934084996,
  "dead_tuples_generated": 208649,
  "dead_tuples_per_commit": 41.2757665677547,
  "app_latency": {
    "p50": 0,
    "p95": 665700,
    "p99": 4308400
  },
  "db_latency": {
    "p50": 5214900,
    "p95": 11527700,
    "p99": 16615800
  },
  "cc_wait_latency": {
    "p50": 0,
    "p95": 0,
    "p99": 0
  },
  "wal_wait_proxy": {
    "p50": 1038900,
    "p95": 5770500,
    "p99": 10670700
  },
  "client_e2e_latency": {
    "p50": 34501800,
    "p95": 108793200,
    "p99": 138388900
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



