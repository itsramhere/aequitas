# investigate_anomalies.ps1

# Isolation runs go to the investigation directory, never into results\set2,
# so primary experiment data is not contaminated with diagnostic runs.
$OutputDir = ".\results\set2_investigation"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Force -Path $OutputDir | Out-Null
}

Write-Host "=== Starting Unexpected Findings Investigation ===" -ForegroundColor Green
Write-Host "Targeting Skew 0.8 with 250 Concurrent Clients (30s each)" -ForegroundColor Yellow
Write-Host "------------------------------------------------"

# 1. Isolate OCC Deadlocks & Verify Zero-Bloat Retries
$occOut = "$OutputDir\Anomaly_OCC.json"
Write-Host "[1/2] Running OCC..." -ForegroundColor Cyan
Write-Host "Expectation: deadlock_count > 0, dead_tuples_per_commit near 0"
go run .\cmd\workload_gen\main.go `
    -strategy OCC `
    -skew 0.8 `
    -concurrency 250 `
    -duration 30s `
    -warmup 5s `
    -enable-stage-a=false `
    -output $occOut

# 2. Isolate Pessimistic Dead Tuple Generation
$pessOut = "$OutputDir\Anomaly_PESSIMISTIC.json"
Write-Host "`n[2/2] Running PESSIMISTIC..." -ForegroundColor Cyan
Write-Host "Expectation: deadlock_count == 0, dead_tuples_per_commit > OCC"
go run .\cmd\workload_gen\main.go `
    -strategy PESSIMISTIC `
    -skew 0.8 `
    -concurrency 250 `
    -duration 30s `
    -warmup 5s `
    -enable-stage-a=false `
    -output $pessOut

Write-Host "`n=== Investigation Complete ===" -ForegroundColor Green
Write-Host "Review the JSON files in $OutputDir"