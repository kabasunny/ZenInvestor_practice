Function CreateProtoFiles {
    Write-Host "Creating directory structure for shared-protos..."

    $baseProtoDir = "../shared-protos"

    $protoFiles = @(
        "data.proto",
        "other_service.proto"
    )

    foreach ($file in $protoFiles) {
        $filePath = Join-Path $baseProtoDir $file
        if (-not (Test-Path $filePath)) {
            New-Item -Path $filePath -ItemType File -Force
            Write-Host "Created proto file: $filePath"
        } else {
            Write-Host "Proto file already exists: $filePath"
        }
    }

    Write-Host "Proto directory structure creation complete."
}

# 関数を呼び出す
CreateProtoFiles
