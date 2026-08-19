# Benchmark Raw Data Archive

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



