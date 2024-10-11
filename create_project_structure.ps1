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
New-Item -ItemType Directory -Path frontend
New-Item -ItemType Directory -Path api
New-Item -ItemType Directory -Path data-analysis

# frontendディレクトリのファイルを作成
$frontendDockerfile = @"
FROM node:18
WORKDIR /app/frontend
COPY package.json package-lock.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "start"]
"@
New-Item -ItemType File -Path frontend\Dockerfile -Value $frontendDockerfile

$frontendPackageJson = @"
{
  "name": "frontend",
  "version": "1.0.0",
  "main": "index.js",
  "dependencies": {
    "react": "^17.0.2",
    "react-dom": "^17.0.2",
    "react-scripts": "4.0.3",
    "typescript": "^4.1.3"
  },
  "scripts": {
    "start": "react-scripts start",
    "build": "react-scripts build",
    "test": "react-scripts test",
    "eject": "react-scripts eject"
  },
  "eslintConfig": {
    "extends": [
      "react-app",
      "react-app/jest"
    ]
  },
  "browserslist": {
    "production": [
      ">0.2%",
      "not dead",
      "not op_mini all"
    ],
    "development": [
      "last 1 chrome version",
      "last 1 firefox version",
      "last 1 safari version"
    ]
  }
}
"@
New-Item -ItemType File -Path frontend\package.json -Value $frontendPackageJson

# apiディレクトリのファイルを作成
$apiDockerfile = @"
FROM golang:1.20
WORKDIR /app/api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080
CMD ["air"]
"@
New-Item -ItemType File -Path api\Dockerfile -Value $apiDockerfile

$apiMainGo = @"
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
    CreatedAt time.Time
}

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
}
"@
New-Item -ItemType File -Path api\main.go -Value $apiMainGo

# data-analysisディレクトリのファイルを作成
$dataAnalysisDockerfile = @"
FROM python:3.11
WORKDIR /app/data-analysis
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 5000
CMD ["flask", "run", "--host=0.0.0.0", "--reload"]
"@
New-Item -ItemType File -Path data-analysis\Dockerfile -Value $dataAnalysisDockerfile

$requirementsTxt = @"
yfinance
pandas
numpy
flask
"@
New-Item -ItemType File -Path data-analysis\requirements.txt -Value $requirementsTxt

# Flask アプリケーションの main ファイルを作成
$dataAnalysisMainPy = @"
from flask import Flask, jsonify
import yfinance as yf
import pandas as pd

app = Flask(__name__)

@app.route('/stocks/<ticker>', methods=['GET'])
def get_stock_data(ticker):
    stock = yf.Ticker(ticker)
    hist = stock.history(period="1mo")
    data = hist[['Open', 'High', 'Low', 'Close']].to_dict(orient='index')
    return jsonify(data)

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
"@
New-Item -ItemType File -Path data-analysis\main.py -Value $dataAnalysisMainPy

# プロジェクトルートに docker-compose.yaml を作成
$dockerComposeYaml = @"
version: '3.8'
services:
  frontend:
    build: ./frontend
    volumes:
      - ./frontend:/app/frontend
      - /app/frontend/node_modules
    ports:
      - "3004:3000"

  api:
    build: ./api
    volumes:
      - ./api:/app/api
    ports:
      - "8084:8080"
    depends_on:
      - mysql

  data-analysis:
    build: ./data-analysis
    volumes:
      - ./data-analysis:/app/data-analysis
    ports:
      - "5004:5000"
    depends_on:
      - mysql
    environment:
      - FLASK_PORT=${FLASK_PORT}

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

# ファイルのパス
$dockerComposePath = "$rootPath\docker-compose.yml"

# docker-compose.yaml を作成
New-Item -ItemType File -Path $dockerComposePath -Value $dockerComposeYaml

Write-Output "docker-compose.yml has been created at $dockerComposePath"

