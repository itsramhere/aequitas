# investigate_anomalies.ps1

$OutputDir = ".\results\set2"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Path $OutputDir | Out-Null
}

Write-Host "=== Starting Unexpected Findings Investigation ==="
Write-Host "Targeting Skew 0.8 with 250 Concurrent Clients"
Write-Host "------------------------------------------------"

# 1. Isolate OCC Deadlocks & Verify Zero-Bloat Retries
Write-Host "[1/2] Running OCC..."
Write-Host "Expectation: deadlock_count > 0, dead_tuples_per_commit near 0"
go run .\cmd\workload_gen\main.go -strategy="OCC" -skew=0.8 -concurrency=250 -duration=30s -enable-stage-a=false -output="$OutputDir\Anomaly_OCC.json"

# 2. Isolate Pessimistic Dead Tuple Generation
Write-Host "`n[2/2] Running PESSIMISTIC..."
Write-Host "Expectation: deadlock_count == 0, dead_tuples_per_commit > OCC"
go run .\cmd\workload_gen\main.go -strategy="PESSIMISTIC" -skew=0.8 -concurrency=250 -duration=30s -enable-stage-a=false -output="$OutputDir\Anomaly_PESSIMISTIC.json"

Write-Host "`n=== Investigation Complete ==="
Write-Host "Review the JSON files in $OutputDir"