# .env ファイルのパス
$envFilePath = ".env"

# .env ファイルの読み込み
if (Test-Path $envFilePath) {
    Get-Content $envFilePath | ForEach-Object {
        # 空行やコメント行を無視
        if ($_ -match "^\s*#") { return }
        if ($_ -match "^\s*$") { return }
        # 環境変数を設定
        $name, $value = $_ -split "=", 2
        $name = $name.Trim()
        $value = $value.Trim()
        [System.Environment]::SetEnvironmentVariable($name, $value)
    }
}

# 環境変数を使用
$rootPath = $env:ROOT_PATH
Write-Output "Root path is: $rootPath"

# プロジェクトディレクトリへ移動
cd $rootPath

# プロジェクトのルートディレクトリを作成
New-Item -ItemType Directory -Path api -Force
New-Item -ItemType Directory -Path "$rootPath/api/migration" -Force  # migrationフォルダを作成

# API Dockerfile の作成
$apiDockerfileContent = @"
FROM golang:1.23

WORKDIR /app/api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# airのインストールとPATHに追加
RUN go install github.com/air-verse/air@latest

# PATHを設定するためにCMDを直接実行
ENV PATH="/root/go/bin:${PATH}"

EXPOSE 8080
CMD ["air"]

"@
New-Item -ItemType File -Path "$rootPath/api/Dockerfile" -Value $apiDockerfileContent -Force
Write-Output "Dockerfile for API has been created."

# API migration.go の作成
$apiMainGoContent = @"
package main

import (
    ""fmt""
    ""gorm.io/driver/mysql""
    ""gorm.io/gorm""
    ""os""
    ""time""
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    CreatedAt time.Time
}

func main() {
    dsn := fmt.Sprintf(""%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"", 
                      os.Getenv(""MYSQL_USER""), 
                      os.Getenv(""MYSQL_PASSWORD""), 
                      os.Getenv(""MYSQL_HOST""), 
                      os.Getenv(""MYSQL_DB""))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(""failed to connect database"")
    }

    db.AutoMigrate(&User{})
}
"@
New-Item -ItemType File -Path "$rootPath/api/migration/migration.go" -Value $apiMainGoContent -Force
Write-Output "migration.go for API has been created."

# go.modとgo.sumの存在を確認し、存在しない場合は作成
if (-Not (Test-Path "$rootPath/api/go.mod")) {
    cd "$rootPath/api"
    go mod init api
    go get gorm.io/gorm
    go get gorm.io/driver/mysql
    go mod tidy
    cd .. 
}
if (-Not (Test-Path "$rootPath/api/go.sum")) {
    cd "$rootPath/api"
    go mod tidy
    cd .. 
}

# docker-compose.yml の作成
$dockerComposeContent = @"
version: '3.8'
services:
  api:
    build: ./api
    ports:
      - "8084:8080"
    environment:
      MYSQL_HOST: ${MYSQL_HOST}
      MYSQL_USER: ${MYSQL_USER}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}
      MYSQL_DB: ${MYSQL_DB}
    depends_on:
      - mysql
    tty: true

  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: ${MYSQL_DATABASE}
      MYSQL_USER: ${MYSQL_USER}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    environment:
      PMA_HOST: mysql
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
    ports:
      - "84:80"
    depends_on:
      - mysql

volumes:
  mysql_data:

"@
New-Item -ItemType File -Path "$rootPath/docker-compose.yml" -Value $dockerComposeContent -Force
Write-Output "docker-compose.yml has been created."


# Dockerコンテナの起動
docker-compose up --build

Write-Output "Directory structure has been created."

Write-Output "Project structure has been created."

# .\create-project-structure.ps1