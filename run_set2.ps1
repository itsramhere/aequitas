# run_set2.ps1

$OutputDir = ".\results\set2AfterInvestigation"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Force -Path $OutputDir | Out-Null
}

Write-Host "=== STARTING EXPERIMENT SET 2 (IDEMPOTENCY SWEEP) ===" -ForegroundColor Green

$strategies = @("SSI", "PESSIMISTIC", "OCC")

foreach ($strat in $strategies) {
    $outFile = "$OutputDir\Set2_${strat}_Idempotency.json"
    Write-Host "Running Stage A/B Idempotency Sweep: Strategy=$strat..." -ForegroundColor Cyan
    
    # Run workload generator with Stage A ENABLED
    go run .\cmd\workload_gen\main.go `
        -strategy $strat `
        -skew 0.8 `
        -concurrency 50 `
        -duration 30s `
        -warmup 5s `
        -enable-stage-a=true `
        -output $outFile

    # Run Independent Auditor Check
    $auditFile = "$OutputDir\Audit_${strat}.json"
    Write-Host "Running Independent Cardinality Audit for $strat..." -ForegroundColor Yellow
    go run .\cmd\auditor\main.go `
        -initial-balance 1000.00 `
        -output $auditFile
}

Write-Host "`n=== SET 2 AND AUDIT COMPLETE ===" -ForegroundColor Green