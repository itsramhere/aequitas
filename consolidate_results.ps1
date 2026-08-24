# consolidate_results.ps1
# Consolidates all benchmark JSON results into a single Markdown file for the Detailed Project Report.

$OutputFile = ".\all_benchmark_results.md"
$ResultsDir = ".\results"

Write-Host "Searching for JSON benchmark results in '$ResultsDir'..." -ForegroundColor Cyan

if (-not (Test-Path -Path $ResultsDir)) {
    Write-Error "Results directory '$ResultsDir' does not exist."
    exit 1
}

# Recursively find all .json files under .\results
$jsonFiles = Get-ChildItem -Path $ResultsDir -Recurse -Filter "*.json" | Sort-Object FullName

if ($jsonFiles.Count -eq 0) {
    Write-Host "No .json files found in '$ResultsDir'." -ForegroundColor Yellow
    exit 0
}

Write-Host "Found $($jsonFiles.Count) JSON files. Consolidating into '$OutputFile'..." -ForegroundColor Cyan

# Build entire Markdown output in memory using StringBuilder
$sb = [System.Text.StringBuilder]::new()
[void]$sb.AppendLine("# Benchmark Raw Data Archive")
[void]$sb.AppendLine()
[void]$sb.AppendLine("_Generated: $((Get-Date).ToUniversalTime().ToString('yyyy-MM-dd HH:mm:ss')) UTC by consolidate_results.ps1._")
[void]$sb.AppendLine()
[void]$sb.AppendLine("_Note: result files were captured by different instrument versions over the life of this project (e.g. older files carry ``abort_retry_rate_pct`` while current code emits ``retries_per_request``/``client_failure_rate_pct``). Field sets are NOT uniform across files; check each JSON's fields before cross-file comparison._")

foreach ($file in $jsonFiles) {
    # Compute relative path (e.g., results\set1\OCC_skew_0.8_conc_50.json)
    $relativePath = (Resolve-Path -Relative $file.FullName) -replace '^\.\\', ''
    
    # Read raw JSON contents
    $jsonContent = Get-Content -Path $file.FullName -Raw
    
    # Append level-3 heading, code block backticks, and blank lines
    [void]$sb.AppendLine()
    [void]$sb.AppendLine("### $relativePath")
    [void]$sb.AppendLine()
    [void]$sb.AppendLine('```json')
    [void]$sb.AppendLine($jsonContent.Trim())
    [void]$sb.AppendLine('```')
    [void]$sb.AppendLine()
    [void]$sb.AppendLine()
}

# Write out accumulated content to file (overwriting existing content)
Set-Content -Path $OutputFile -Value $sb.ToString() -Encoding utf8

Write-Host "Success! Consolidated $($jsonFiles.Count) benchmark JSON files into '$OutputFile'." -ForegroundColor Green
