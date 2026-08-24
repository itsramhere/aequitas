# Create output directory for results
New-Item -ItemType Directory -Force -Path ".\results\set1" | Out-Null

$strategies = @("SSI", "PESSIMISTIC", "OCC")
$skews = @(0.0, 0.4, 0.6, 0.8, 1.0, 1.2)
$concurrenctLevels = @(10, 50, 100, 250)

Write-Host "=== STARTING EXPERIMENT SET 1 SWEEP ===" -ForegroundColor Green

# 1. Skew Sweep (Fixed Concurrency = 50)
foreach ($strat in $strategies) {
    foreach ($skew in $skews) {
        # Format filename (e.g., SSI_skew_0.8_conc_50.json)
        $outFile = ".\results\set1\${strat}_skew_${skew}_conc_50.json"
        Write-Host "Running: Strategy=$strat, Skew=$skew, Concurrency=50..." -ForegroundColor Cyan
        
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

# 2. Concurrency Sweep (Fixed Skew = 0.8, excluding concurrency=50 which was run above)
foreach ($strat in $strategies) {
    foreach ($conc in $concurrenctLevels) {
        if ($conc -eq 50) { continue } # Skip reference cell already run
        
        $outFile = ".\results\set1\${strat}_skew_0.8_conc_${conc}.json"
        Write-Host "Running: Strategy=$strat, Skew=0.8, Concurrency=$conc..." -ForegroundColor Cyan
        
        go run .\cmd\workload_gen\main.go `
            -strategy $strat `
            -skew 0.8 `
            -concurrency $conc `
            -duration 30s `
            -warmup 5s `
            -enable-stage-a=false `
            -output $outFile
    }
}

Write-Host "=== EXPERIMENT SET 1 COMPLETE! All JSON reports saved to .\results\set1 ===" -ForegroundColor Green