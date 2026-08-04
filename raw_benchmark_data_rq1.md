# Combined JSON Files

## OCC_skew_0.4_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 119240,
  "committed_txns": 119239,
  "client_visible_failures": 1,
  "internal_retries": 2794,
  "deadlock_count": 0,
  "throughput_tps": 3974.633333333333,
  "abort_retry_rate_pct": 2.3440120764844012,
  "realized_hot_row_hit_rate_pct": 0.2312610025016214,
  "dead_tuples_generated": 6097,
  "dead_tuples_per_commit": 0.05113259923347227,
  "app_latency": {
    "p50": 0,
    "p95": 1508700,
    "p99": 7104000
  },
  "db_latency": {
    "p50": 9622500,
    "p95": 23711800,
    "p99": 47806700
  },
  "wal_wait_proxy": {
    "p50": 1512800,
    "p95": 10198900,
    "p99": 18099200
  },
  "client_e2e_latency": {
    "p50": 10192200,
    "p95": 26893700,
    "p99": 55502100
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

## OCC_skew_0.6_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 117986,
  "committed_txns": 117587,
  "client_visible_failures": 399,
  "internal_retries": 11249,
  "deadlock_count": 1,
  "throughput_tps": 3919.5666666666666,
  "abort_retry_rate_pct": 9.87235773735867,
  "realized_hot_row_hit_rate_pct": 1.006711409395973,
  "dead_tuples_generated": 6353,
  "dead_tuples_per_commit": 0.05402808133552178,
  "app_latency": {
    "p50": 0,
    "p95": 1508500,
    "p99": 6383000
  },
  "db_latency": {
    "p50": 8566400,
    "p95": 20145800,
    "p99": 34844600
  },
  "wal_wait_proxy": {
    "p50": 1359600,
    "p95": 9028400,
    "p99": 15677400
  },
  "client_e2e_latency": {
    "p50": 9275700,
    "p95": 26910900,
    "p99": 65700600
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

## OCC_skew_0.8_conc_10.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 75562,
  "committed_txns": 75427,
  "client_visible_failures": 135,
  "internal_retries": 7585,
  "deadlock_count": 2,
  "throughput_tps": 2514.233333333333,
  "abort_retry_rate_pct": 10.216775627961145,
  "realized_hot_row_hit_rate_pct": 3.651587260679576,
  "dead_tuples_generated": 1354,
  "dead_tuples_per_commit": 0.017951131557665027,
  "app_latency": {
    "p50": 0,
    "p95": 564100,
    "p99": 666100
  },
  "db_latency": {
    "p50": 2714800,
    "p95": 3533000,
    "p99": 4550800
  },
  "wal_wait_proxy": {
    "p50": 531700,
    "p95": 678100,
    "p99": 1359000
  },
  "client_e2e_latency": {
    "p50": 2764000,
    "p95": 8277200,
    "p99": 19835800
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

## OCC_skew_0.8_conc_100.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 34042,
  "committed_txns": 31991,
  "client_visible_failures": 2051,
  "internal_retries": 16768,
  "deadlock_count": 131,
  "throughput_tps": 1066.3666666666666,
  "abort_retry_rate_pct": 55.28171082780096,
  "realized_hot_row_hit_rate_pct": 3.571823509897158,
  "dead_tuples_generated": 14917,
  "dead_tuples_per_commit": 0.4662873933293739,
  "app_latency": {
    "p50": 0,
    "p95": 1026300,
    "p99": 3738500
  },
  "db_latency": {
    "p50": 5677700,
    "p95": 16041000,
    "p99": 104133000
  },
  "wal_wait_proxy": {
    "p50": 1009100,
    "p95": 5577300,
    "p99": 11023300
  },
  "client_e2e_latency": {
    "p50": 6598800,
    "p95": 113750800,
    "p99": 1171587300
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

## OCC_skew_0.8_conc_250.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 25432,
  "committed_txns": 21419,
  "client_visible_failures": 4013,
  "internal_retries": 9813,
  "deadlock_count": 168,
  "throughput_tps": 713.9666666666667,
  "abort_retry_rate_pct": 54.364580056621584,
  "realized_hot_row_hit_rate_pct": 3.757114536036324,
  "dead_tuples_generated": 12747,
  "dead_tuples_per_commit": 0.5951258228675474,
  "app_latency": {
    "p50": 0,
    "p95": 1295200,
    "p99": 4041600
  },
  "db_latency": {
    "p50": 5729200,
    "p95": 25841500,
    "p99": 342461700
  },
  "wal_wait_proxy": {
    "p50": 1008900,
    "p95": 6047300,
    "p99": 17594400
  },
  "client_e2e_latency": {
    "p50": 6594800,
    "p95": 220987600,
    "p99": 1608465000
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

## OCC_skew_0.8_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 66430,
  "committed_txns": 63574,
  "client_visible_failures": 2856,
  "internal_retries": 25862,
  "deadlock_count": 15,
  "throughput_tps": 2119.133333333333,
  "abort_retry_rate_pct": 43.23046816197501,
  "realized_hot_row_hit_rate_pct": 3.6228890151319524,
  "dead_tuples_generated": 5601,
  "dead_tuples_per_commit": 0.08810205429892723,
  "app_latency": {
    "p50": 0,
    "p95": 1013900,
    "p99": 3404400
  },
  "db_latency": {
    "p50": 5134400,
    "p95": 12797600,
    "p99": 22674000
  },
  "wal_wait_proxy": {
    "p50": 704400,
    "p95": 4242600,
    "p99": 9184400
  },
  "client_e2e_latency": {
    "p50": 5864100,
    "p95": 33206900,
    "p99": 99060500
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

## OCC_skew_0_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 125265,
  "committed_txns": 125265,
  "client_visible_failures": 0,
  "internal_retries": 1726,
  "deadlock_count": 0,
  "throughput_tps": 4175.5,
  "abort_retry_rate_pct": 1.3778788967389135,
  "realized_hot_row_hit_rate_pct": 0.008621522077993738,
  "dead_tuples_generated": 5141,
  "dead_tuples_per_commit": 0.041040993094639364,
  "app_latency": {
    "p50": 0,
    "p95": 1506800,
    "p99": 6439200
  },
  "db_latency": {
    "p50": 9527200,
    "p95": 22704500,
    "p99": 45339000
  },
  "wal_wait_proxy": {
    "p50": 1512400,
    "p95": 10115800,
    "p99": 17857500
  },
  "client_e2e_latency": {
    "p50": 10041400,
    "p95": 24715100,
    "p99": 49767900
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

## OCC_skew_1.2_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 2346,
  "committed_txns": 1611,
  "client_visible_failures": 735,
  "internal_retries": 4069,
  "deadlock_count": 68,
  "throughput_tps": 53.7,
  "abort_retry_rate_pct": 204.77408354646207,
  "realized_hot_row_hit_rate_pct": 21.3528618230873,
  "dead_tuples_generated": 746,
  "dead_tuples_per_commit": 0.46306641837368095,
  "app_latency": {
    "p50": 0,
    "p95": 1007800,
    "p99": 2007900
  },
  "db_latency": {
    "p50": 4189800,
    "p95": 25995100,
    "p99": 2022755600
  },
  "wal_wait_proxy": {
    "p50": 523800,
    "p95": 2000300,
    "p99": 5121000
  },
  "client_e2e_latency": {
    "p50": 6446400,
    "p95": 1132283800,
    "p99": 3093342400
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

## OCC_skew_1_conc_50.json

```json
{
  "strategy": "OCC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 19895,
  "committed_txns": 17252,
  "client_visible_failures": 2643,
  "internal_retries": 18571,
  "deadlock_count": 69,
  "throughput_tps": 575.0666666666667,
  "abort_retry_rate_pct": 106.62980648404121,
  "realized_hot_row_hit_rate_pct": 10.215579556084702,
  "dead_tuples_generated": 3279,
  "dead_tuples_per_commit": 0.19006492000927427,
  "app_latency": {
    "p50": 0,
    "p95": 822100,
    "p99": 1646500
  },
  "db_latency": {
    "p50": 3711600,
    "p95": 9617300,
    "p99": 33286700
  },
  "wal_wait_proxy": {
    "p50": 521800,
    "p95": 2079500,
    "p99": 5946600
  },
  "client_e2e_latency": {
    "p50": 4629500,
    "p95": 60007000,
    "p99": 1063231900
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

## PESSIMISTIC_skew_0.4_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 142298,
  "committed_txns": 142298,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 4743.266666666666,
  "abort_retry_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.237134022136599,
  "dead_tuples_generated": 4273,
  "dead_tuples_per_commit": 0.030028531672967998,
  "app_latency": {
    "p50": 540300,
    "p95": 2516400,
    "p99": 7477200
  },
  "db_latency": {
    "p50": 8268600,
    "p95": 19693700,
    "p99": 36325400
  },
  "wal_wait_proxy": {
    "p50": 1114300,
    "p95": 8502900,
    "p99": 15398800
  },
  "client_e2e_latency": {
    "p50": 9001100,
    "p95": 21031500,
    "p99": 40866100
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

## PESSIMISTIC_skew_0.6_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 138865,
  "committed_txns": 138865,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 4628.833333333333,
  "abort_retry_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 1.0383257865649618,
  "dead_tuples_generated": 15168,
  "dead_tuples_per_commit": 0.10922838728261261,
  "app_latency": {
    "p50": 535800,
    "p95": 2409400,
    "p99": 6563700
  },
  "db_latency": {
    "p50": 7865100,
    "p95": 20624800,
    "p99": 52327200
  },
  "wal_wait_proxy": {
    "p50": 1104600,
    "p95": 7882400,
    "p99": 13935700
  },
  "client_e2e_latency": {
    "p50": 8512000,
    "p95": 21845200,
    "p99": 55674000
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

## PESSIMISTIC_skew_0.8_conc_10.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 88174,
  "committed_txns": 88174,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 2939.133333333333,
  "abort_retry_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 3.573084921232552,
  "dead_tuples_generated": 1179,
  "dead_tuples_per_commit": 0.013371288588472792,
  "app_latency": {
    "p50": 0,
    "p95": 1007800,
    "p99": 1469700
  },
  "db_latency": {
    "p50": 2832800,
    "p95": 5077600,
    "p99": 8212600
  },
  "wal_wait_proxy": {
    "p50": 535000,
    "p95": 1223600,
    "p99": 1540300
  },
  "client_e2e_latency": {
    "p50": 3104100,
    "p95": 5438000,
    "p99": 8553900
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

## PESSIMISTIC_skew_0.8_conc_100.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 101048,
  "committed_txns": 100905,
  "client_visible_failures": 143,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 3363.5,
  "abort_retry_rate_pct": 0.14151690285804766,
  "realized_hot_row_hit_rate_pct": 3.6608311168506846,
  "dead_tuples_generated": 123798,
  "dead_tuples_per_commit": 1.2268767652742678,
  "app_latency": {
    "p50": 0,
    "p95": 1335700,
    "p99": 1607700
  },
  "db_latency": {
    "p50": 3681600,
    "p95": 56663400,
    "p99": 803500600
  },
  "wal_wait_proxy": {
    "p50": 551400,
    "p95": 1998000,
    "p99": 3457500
  },
  "client_e2e_latency": {
    "p50": 4015500,
    "p95": 56888400,
    "p99": 804384100
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

## PESSIMISTIC_skew_0.8_conc_250.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 65847,
  "committed_txns": 63908,
  "client_visible_failures": 1939,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 2130.266666666667,
  "abort_retry_rate_pct": 2.944705149817,
  "realized_hot_row_hit_rate_pct": 3.6778739788269634,
  "dead_tuples_generated": 90334,
  "dead_tuples_per_commit": 1.4135006571947175,
  "app_latency": {
    "p50": 0,
    "p95": 1507300,
    "p99": 3436900
  },
  "db_latency": {
    "p50": 4976000,
    "p95": 54402200,
    "p99": 1471308400
  },
  "wal_wait_proxy": {
    "p50": 995800,
    "p95": 3011800,
    "p99": 7148000
  },
  "client_e2e_latency": {
    "p50": 5395600,
    "p95": 56111300,
    "p99": 1472315100
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

## PESSIMISTIC_skew_0.8_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 122938,
  "committed_txns": 122897,
  "client_visible_failures": 41,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 4096.566666666667,
  "abort_retry_rate_pct": 0.03335014397501179,
  "realized_hot_row_hit_rate_pct": 3.702829682706541,
  "dead_tuples_generated": 86481,
  "dead_tuples_per_commit": 0.7036868271804845,
  "app_latency": {
    "p50": 0,
    "p95": 1319300,
    "p99": 1508200
  },
  "db_latency": {
    "p50": 3377600,
    "p95": 21127500,
    "p99": 233301400
  },
  "wal_wait_proxy": {
    "p50": 0,
    "p95": 1480100,
    "p99": 2320800
  },
  "client_e2e_latency": {
    "p50": 3548600,
    "p95": 21650300,
    "p99": 233536000
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

## PESSIMISTIC_skew_0_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 140013,
  "committed_txns": 140013,
  "client_visible_failures": 0,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 4667.1,
  "abort_retry_rate_pct": 0,
  "realized_hot_row_hit_rate_pct": 0.006852984786373775,
  "dead_tuples_generated": 5262,
  "dead_tuples_per_commit": 0.037582224507724284,
  "app_latency": {
    "p50": 552000,
    "p95": 2999200,
    "p99": 7315300
  },
  "db_latency": {
    "p50": 8515300,
    "p95": 19599800,
    "p99": 35394600
  },
  "wal_wait_proxy": {
    "p50": 1163200,
    "p95": 8951700,
    "p99": 15309600
  },
  "client_e2e_latency": {
    "p50": 9233300,
    "p95": 20853500,
    "p99": 39420300
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

## PESSIMISTIC_skew_1.2_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 31279,
  "committed_txns": 30101,
  "client_visible_failures": 1178,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 1003.3666666666667,
  "abort_retry_rate_pct": 3.7661050545094152,
  "realized_hot_row_hit_rate_pct": 20.79856812187113,
  "dead_tuples_generated": 11242,
  "dead_tuples_per_commit": 0.3734759642536793,
  "app_latency": {
    "p50": 0,
    "p95": 649300,
    "p99": 779000
  },
  "db_latency": {
    "p50": 3368000,
    "p95": 245590200,
    "p99": 507304500
  },
  "wal_wait_proxy": {
    "p50": 543700,
    "p95": 720700,
    "p99": 1353800
  },
  "client_e2e_latency": {
    "p50": 3534900,
    "p95": 245617600,
    "p99": 507367600
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

## PESSIMISTIC_skew_1_conc_50.json

```json
{
  "strategy": "PESSIMISTIC",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 60653,
  "committed_txns": 60163,
  "client_visible_failures": 490,
  "internal_retries": 0,
  "deadlock_count": 0,
  "throughput_tps": 2005.4333333333334,
  "abort_retry_rate_pct": 0.8078743013536017,
  "realized_hot_row_hit_rate_pct": 10.214541205817628,
  "dead_tuples_generated": 21537,
  "dead_tuples_per_commit": 0.3579774944733474,
  "app_latency": {
    "p50": 0,
    "p95": 607400,
    "p99": 684500
  },
  "db_latency": {
    "p50": 2806000,
    "p95": 142663100,
    "p99": 390265100
  },
  "wal_wait_proxy": {
    "p50": 543700,
    "p95": 653200,
    "p99": 1218100
  },
  "client_e2e_latency": {
    "p50": 2876800,
    "p95": 142748200,
    "p99": 390265100
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

## SSI_skew_0.4_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.4,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 166334,
  "committed_txns": 165916,
  "client_visible_failures": 418,
  "internal_retries": 4033,
  "deadlock_count": 0,
  "throughput_tps": 5530.533333333334,
  "abort_retry_rate_pct": 2.6759411785924705,
  "realized_hot_row_hit_rate_pct": 0.24514892597620377,
  "dead_tuples_generated": 19286,
  "dead_tuples_per_commit": 0.11623954290122712,
  "app_latency": {
    "p50": 0,
    "p95": 1069100,
    "p99": 5000700
  },
  "db_latency": {
    "p50": 6999400,
    "p95": 17109700,
    "p99": 29446800
  },
  "wal_wait_proxy": {
    "p50": 1124900,
    "p95": 8581600,
    "p99": 16028400
  },
  "client_e2e_latency": {
    "p50": 7426900,
    "p95": 19151400,
    "p99": 35398700
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

## SSI_skew_0.6_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.6,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 155207,
  "committed_txns": 154038,
  "client_visible_failures": 1169,
  "internal_retries": 15171,
  "deadlock_count": 1,
  "throughput_tps": 5134.6,
  "abort_retry_rate_pct": 10.527875675710503,
  "realized_hot_row_hit_rate_pct": 1.0584059006555968,
  "dead_tuples_generated": 12052,
  "dead_tuples_per_commit": 0.07824043417857932,
  "app_latency": {
    "p50": 0,
    "p95": 1054200,
    "p99": 4544500
  },
  "db_latency": {
    "p50": 6136100,
    "p95": 15635300,
    "p99": 25994400
  },
  "wal_wait_proxy": {
    "p50": 1055600,
    "p95": 8081000,
    "p99": 14815300
  },
  "client_e2e_latency": {
    "p50": 6787000,
    "p95": 20752700,
    "p99": 51641800
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

## SSI_skew_0.8_conc_10.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 10,
  "duration": 30000000000,
  "total_requests": 76069,
  "committed_txns": 75877,
  "client_visible_failures": 192,
  "internal_retries": 8495,
  "deadlock_count": 0,
  "throughput_tps": 2529.233333333333,
  "abort_retry_rate_pct": 11.419895095242477,
  "realized_hot_row_hit_rate_pct": 3.6841072779755066,
  "dead_tuples_generated": 1479,
  "dead_tuples_per_commit": 0.019492072696601078,
  "app_latency": {
    "p50": 0,
    "p95": 588600,
    "p99": 1009700
  },
  "db_latency": {
    "p50": 2696900,
    "p95": 4336600,
    "p99": 6351300
  },
  "wal_wait_proxy": {
    "p50": 537200,
    "p95": 1088800,
    "p99": 1680200
  },
  "client_e2e_latency": {
    "p50": 2782400,
    "p95": 8746000,
    "p99": 21241200
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

## SSI_skew_0.8_conc_100.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 100,
  "duration": 30000000000,
  "total_requests": 91880,
  "committed_txns": 83133,
  "client_visible_failures": 8747,
  "internal_retries": 47844,
  "deadlock_count": 6,
  "throughput_tps": 2771.1,
  "abort_retry_rate_pct": 61.59229429690901,
  "realized_hot_row_hit_rate_pct": 3.6987271177379974,
  "dead_tuples_generated": 14706,
  "dead_tuples_per_commit": 0.17689726101548123,
  "app_latency": {
    "p50": 0,
    "p95": 1522900,
    "p99": 6997200
  },
  "db_latency": {
    "p50": 7947500,
    "p95": 31087700,
    "p99": 63280800
  },
  "wal_wait_proxy": {
    "p50": 1528600,
    "p95": 17421600,
    "p99": 47364000
  },
  "client_e2e_latency": {
    "p50": 10130500,
    "p95": 68461400,
    "p99": 142365900
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

## SSI_skew_0.8_conc_250.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 250,
  "duration": 30000000000,
  "total_requests": 61722,
  "committed_txns": 54409,
  "client_visible_failures": 7313,
  "internal_retries": 31103,
  "deadlock_count": 56,
  "throughput_tps": 1813.6333333333334,
  "abort_retry_rate_pct": 62.24036810213538,
  "realized_hot_row_hit_rate_pct": 3.7167715241172137,
  "dead_tuples_generated": 13763,
  "dead_tuples_per_commit": 0.25295447444356634,
  "app_latency": {
    "p50": 0,
    "p95": 1586900,
    "p99": 11220700
  },
  "db_latency": {
    "p50": 7876100,
    "p95": 50899800,
    "p99": 104778200
  },
  "wal_wait_proxy": {
    "p50": 1519100,
    "p95": 19767400,
    "p99": 71645800
  },
  "client_e2e_latency": {
    "p50": 9836200,
    "p95": 104004000,
    "p99": 268938100
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

## SSI_skew_0.8_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0.8,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 101521,
  "committed_txns": 96295,
  "client_visible_failures": 5226,
  "internal_retries": 39123,
  "deadlock_count": 9,
  "throughput_tps": 3209.8333333333335,
  "abort_retry_rate_pct": 43.68455787472543,
  "realized_hot_row_hit_rate_pct": 3.7031381547452003,
  "dead_tuples_generated": 8585,
  "dead_tuples_per_commit": 0.0891531232151202,
  "app_latency": {
    "p50": 0,
    "p95": 1010100,
    "p99": 3158400
  },
  "db_latency": {
    "p50": 4286300,
    "p95": 11279700,
    "p99": 20901400
  },
  "wal_wait_proxy": {
    "p50": 608800,
    "p95": 4167800,
    "p99": 9345500
  },
  "client_e2e_latency": {
    "p50": 5134800,
    "p95": 28846700,
    "p99": 80599800
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

## SSI_skew_0_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 0,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 152636,
  "committed_txns": 152255,
  "client_visible_failures": 381,
  "internal_retries": 2246,
  "deadlock_count": 0,
  "throughput_tps": 5075.166666666667,
  "abort_retry_rate_pct": 1.7210880788280616,
  "realized_hot_row_hit_rate_pct": 0.01155971580015789,
  "dead_tuples_generated": 10177,
  "dead_tuples_per_commit": 0.06684181143476405,
  "app_latency": {
    "p50": 0,
    "p95": 1251000,
    "p99": 5332300
  },
  "db_latency": {
    "p50": 7541300,
    "p95": 19275100,
    "p99": 37264200
  },
  "wal_wait_proxy": {
    "p50": 1298800,
    "p95": 9545000,
    "p99": 18723500
  },
  "client_e2e_latency": {
    "p50": 8017100,
    "p95": 20913000,
    "p99": 42634500
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

## SSI_skew_1.2_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1.2,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 11026,
  "committed_txns": 7624,
  "client_visible_failures": 3402,
  "internal_retries": 14771,
  "deadlock_count": 44,
  "throughput_tps": 254.13333333333333,
  "abort_retry_rate_pct": 164.81951750408126,
  "realized_hot_row_hit_rate_pct": 20.63436602220669,
  "dead_tuples_generated": 1351,
  "dead_tuples_per_commit": 0.17720356768100734,
  "app_latency": {
    "p50": 0,
    "p95": 1007900,
    "p99": 2249900
  },
  "db_latency": {
    "p50": 2843700,
    "p95": 9988200,
    "p99": 35799000
  },
  "wal_wait_proxy": {
    "p50": 517900,
    "p95": 1530300,
    "p99": 4172300
  },
  "client_e2e_latency": {
    "p50": 4468700,
    "p95": 71422900,
    "p99": 1070740900
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

## SSI_skew_1_conc_50.json

```json
{
  "strategy": "SSI",
  "skew_theta": 1,
  "concurrency": 50,
  "duration": 30000000000,
  "total_requests": 20965,
  "committed_txns": 17545,
  "client_visible_failures": 3420,
  "internal_retries": 18551,
  "deadlock_count": 27,
  "throughput_tps": 584.8333333333334,
  "abort_retry_rate_pct": 104.79847364655377,
  "realized_hot_row_hit_rate_pct": 10.553548963110265,
  "dead_tuples_generated": 3515,
  "dead_tuples_per_commit": 0.20034197777144486,
  "app_latency": {
    "p50": 0,
    "p95": 1010200,
    "p99": 2346800
  },
  "db_latency": {
    "p50": 3143800,
    "p95": 10978900,
    "p99": 31053500
  },
  "wal_wait_proxy": {
    "p50": 521000,
    "p95": 2000500,
    "p99": 4819500
  },
  "client_e2e_latency": {
    "p50": 3989800,
    "p95": 54460500,
    "p99": 1048637800
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

