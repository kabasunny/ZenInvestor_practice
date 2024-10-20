# CreateDirectory-react-src.ps1
Function CreateDirectoryStructure {
    Write-Host "Creating directory structure..."

    $directories = @(
        "src",
        "src/components",
        "src/components/LoginForm",
        "src/pages",
        "src/hooks",
        "src/services",
        "src/utils",
        "src/types"
    )
    
    foreach ($dir in $directories) {
        if (-not (Test-Path $dir)) {
            New-Item -Path $dir -ItemType Directory -Force
            Write-Host "Created directory: $dir"
        } else {
            Write-Host "Directory already exists: $dir"
        }
    }

    Write-Host "Directory structure creation complete."
}

# 関数を呼び出す
CreateDirectoryStructure
