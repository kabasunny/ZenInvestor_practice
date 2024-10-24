Function CreateAnalysisDirectories {
    Write-Host "Creating directory structure for data-analysis-python..."

    
    $basePythonDir  = "../data-analysis-python/src"

    $pythonSubDirs = @(
        "routes",
        "services"
    )

    foreach ($subDir in $pythonSubDirs) {
        $dir = Join-Path $basePythonDir $subDir
        if (-not (Test-Path $dir)) {
            New-Item -Path $dir -ItemType Directory -Force
            Write-Host "Created directory: $dir"
        } else {
            Write-Host "Directory already exists: $dir"
        }
    }

    Write-Host "Python directory structure creation complete."
}

# 関数を呼び出す
CreateAnalysisDirectories
