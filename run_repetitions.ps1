$OutputDir = ".\results\set1_repetitions"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Force -Path $OutputDir | Out-Null
}

$strategies = @("SSI", "PESSIMISTIC", "OCC")
$skews = @(0.6, 0.8, 1.0)
$runs = 3

Write-Host "=== STARTING TARGETED VARIANCE SWEEP ===" -ForegroundColor Green

foreach ($skew in $skews) {
    foreach ($strat in $strategies) {
        for ($i = 1; $i -le $runs; $i++) {
            $outFile = "$OutputDir\${strat}_skew_${skew}_run_${i}.json"
            Write-Host "Running: Strategy=$strat, Skew=$skew, Run=$i/3..." -ForegroundColor Cyan
            
            go run .\cmd\workload_gen\main.go `
                -strategy $strat `
                -skew $skew `
                -concurrency 50 `
                -duration 30s `
                -warmup 5s `
                -enable-stage-a=false `
                -output $outFile
        }
    }
}

Write-Host "=== VARIANCE SWEEP COMPLETE ===" -ForegroundColor Green