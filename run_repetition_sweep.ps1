# run_repetition_sweep.ps1
# Automates 5-repetition empirical benchmark sweep across strategies and skew levels

$OutputDir = ".\results\set1_ci"
if (-not (Test-Path -Path $OutputDir)) {
    New-Item -ItemType Directory -Force -Path $OutputDir | Out-Null
}

$strategies = @("SSI", "PESSIMISTIC", "OCC", "ADAPTIVE")
$skews = @(0.0, 0.4, 0.6, 0.8, 1.0, 1.2)
$repetitions = 1..5

Write-Host "=== STARTING 5-REPETITION EXPERIMENT SWEEP ===" -ForegroundColor Green

foreach ($strat in $strategies) {
    foreach ($skew in $skews) {
        foreach ($rep in $repetitions) {
            $outFile = "$OutputDir\${strat}_skew_${skew}_rep_${rep}.json"
            Write-Host "Running: Strategy=$strat, Skew=$skew, Repetition=$rep..." -ForegroundColor Cyan
            
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

Write-Host "`n=== 5-REPETITION EXPERIMENT SWEEP COMPLETE! ===" -ForegroundColor Green
Write-Host "All reports saved to $OutputDir"
