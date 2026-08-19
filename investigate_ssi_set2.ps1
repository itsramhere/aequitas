# investigate_ssi_set2.ps1

$OutputDir = ".\results\set2_investigation"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Force -Path $OutputDir | Out-Null
}

Write-Host "=== Isolating SSI Idempotency Anomaly ===" -ForegroundColor Green
Write-Host "Targeting Skew 0.8 with 50 Concurrent Clients (Stage A Enabled)" -ForegroundColor Yellow
Write-Host "------------------------------------------------"

$outFile = "$OutputDir\Investigate_SSI_StageA.json"

go run .\cmd\workload_gen\main.go `
    -strategy SSI `
    -skew 0.8 `
    -concurrency 50 `
    -duration 30s `
    -warmup 5s `
    -enable-stage-a=true `
    -output $outFile

Write-Host "`n=== Investigation Complete ===" -ForegroundColor Green
Write-Host "Review $outFile to verify if internal_retries remain disproportionately low compared to client_visible_failures."