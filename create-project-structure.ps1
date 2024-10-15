# create-project-structure.ps1

# Set console encoding to UTF-8
[Console]::InputEncoding = [System.Text.Encoding]::UTF8
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# Path to the .env file
$envFilePath = ".env"

# Check if the .env file exists
if (-Not (Test-Path $envFilePath)) {
    Write-Error ".env file not found. Please check the path: $envFilePath"
    exit 1
}

# Load .env file and set environment variables
$envVars = @{}

Get-Content $envFilePath | ForEach-Object {
    # Ignore empty lines and comments
    if ($_ -match "^\s*#") { return }
    if ($_ -match "^\s*$") { return }
    # Set environment variables
    $name, $value = $_ -split "=", 2
    $name = $name.Trim()
    $value = $value.Trim()
    $envVars[$name] = $value
}

# Retrieve necessary environment variables
$rootPath = $envVars["ROOT_PATH"]
$mysqlRootPassword = $envVars["MYSQL_ROOT_PASSWORD"]
$mysqlDatabase = $envVars["MYSQL_DATABASE"]
$mysqlUser = $envVars["MYSQL_USER"]
$mysqlPassword = $envVars["MYSQL_PASSWORD"]

# Check if root path is set
if (-Not $rootPath) {
    Write-Error "ROOT_PATH is not set in .env file."
    exit 1
}

# Create root directory
Write-Output "Creating root directory: $rootPath"
New-Item -ItemType Directory -Path $rootPath -Force

# Create subdirectories
$directories = @("frontend", "api", "data-analysis")
foreach ($dir in $directories) {
    $fullPath = Join-Path $rootPath $dir
    New-Item -ItemType Directory -Path $fullPath -Force
    Write-Output "Creating directory: $fullPath"
}

# Initialize frontend project (React with TypeScript)
Write-Output "Initializing React TypeScript project"
cd (Join-Path $rootPath "frontend")
npx create-react-app . --template typescript

# Create Dockerfile for frontend
$frontendDockerfileContent = @'
# Dockerfile for frontend
FROM node:20.11.0-alpine

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 3000
CMD ["npm", "start"]
'@

$frontendDockerfilePath = Join-Path $rootPath "frontend\Dockerfile"
Set-Content -Path $frontendDockerfilePath -Value $frontendDockerfileContent -Force -Encoding UTF8
Write-Output "Created Dockerfile for frontend: $frontendDockerfilePath"

# Initialize Go API project
Write-Output "Initializing Go API project"
$apiPath = Join-Path $rootPath "api"
cd $apiPath
go mod init api

# Install Gin and GORM
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# Create sample main.go
$goMainContent = @'
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run()
}
'@

$goMainPath = Join-Path $apiPath "main.go"
Set-Content -Path $goMainPath -Value $goMainContent -Force -Encoding UTF8
Write-Output "Created api/main.go: $goMainPath"

# Create Dockerfile for API (using air for hot reload)
$apiDockerfileContent = @'
# Dockerfile for API
FROM golang:1.21-alpine

WORKDIR /go/src/app

# Install Air
RUN go install github.com/cosmtrek/air@latest

# Copy Go modules and download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Copy Air config file if necessary
# COPY .air.toml .

EXPOSE 8080

CMD ["air"]
'@

$apiDockerfilePath = Join-Path $apiPath "Dockerfile"
Set-Content -Path $apiDockerfilePath -Value $apiDockerfileContent -Force -Encoding UTF8
Write-Output "Created Dockerfile for API: $apiDockerfilePath"

# Initialize data-analysis (Python Flask) project
Write-Output "Initializing Python Flask project"
$dataAnalysisPath = Join-Path $rootPath "data-analysis"
cd $dataAnalysisPath

# Create sample app.py
$flaskAppContent = @'
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return 'Hello from Flask!'

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0')
'@

$flaskAppPath = Join-Path $dataAnalysisPath "app.py"
Set-Content -Path $flaskAppPath -Value $flaskAppContent -Force -Encoding UTF8
Write-Output "Created data-analysis/app.py: $flaskAppPath"

# Create requirements.txt
$requirementsContent = "flask"
$requirementsPath = Join-Path $dataAnalysisPath "requirements.txt"
Set-Content -Path $requirementsPath -Value $requirementsContent -Force -Encoding UTF8
Write-Output "Created requirements.txt: $requirementsPath"

# Create Dockerfile for data-analysis (Flask with hot reload)
$dataAnalysisDockerfileContent = @'
# Dockerfile for data-analysis
FROM python:3.11-alpine

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

EXPOSE 5000

CMD ["flask", "run", "--host=0.0.0.0", "--port=5000", "--reload"]
'@

$dataAnalysisDockerfilePath = Join-Path $dataAnalysisPath "Dockerfile"
Set-Content -Path $dataAnalysisDockerfilePath -Value $dataAnalysisDockerfileContent -Force -Encoding UTF8
Write-Output "Created Dockerfile for data-analysis: $dataAnalysisDockerfilePath"

# Create docker-compose.yml
$dockerComposeContent = @'
version: "3.8"

services:
  # Frontend service
  frontend:
    build: ./frontend
    ports:
      - "3006:3000"
    volumes:
      - ./frontend:/usr/src/app
    environment:
      - CHOKIDAR_USEPOLLING=true
    stdin_open: true
    tty: true

  # API service
  api:
    build: ./api
    ports:
      - "8086:8080"
    volumes:
      - ./api:/go/src/app
    environment:
      - MYSQL_HOST=mysql
      - MYSQL_PORT=3306
      - MYSQL_USER=${MYSQL_USER}
      - MYSQL_PASSWORD=${MYSQL_PASSWORD}
      - MYSQL_DATABASE=${MYSQL_DATABASE}
    depends_on:
      - mysql

  # Data Analysis service
  data-analysis:
    build: ./data-analysis
    ports:
      - "5006:5000"
    volumes:
      - ./data-analysis:/usr/src/app
    environment:
      - FLASK_APP=app.py
      - FLASK_ENV=development

  # MySQL database
  mysql:
    image: mysql:latest
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: ${MYSQL_DATABASE}
      MYSQL_USER: ${MYSQL_USER}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}
    ports:
      - "3306:3306"
    volumes:
      - db_data:/var/lib/mysql

  # phpMyAdmin service
  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    restart: always
    ports:
      - "86:80"
    environment:
      PMA_HOST: mysql
    depends_on:
      - mysql

volumes:
  db_data:
'@

$dockerComposePath = Join-Path $rootPath "docker-compose.yml"
Set-Content -Path $dockerComposePath -Value $dockerComposeContent -Force -Encoding UTF8
Write-Output "Created docker-compose.yml: $dockerComposePath"

# Create .gitignore
$gitignoreContent = @'
# Node modules
/frontend/node_modules

# Python cache
/data-analysis/__pycache__/
/data-analysis/*.pyc

# Go binary
/api/*.exe
/api/air

# Docker
**/Dockerfile
docker-compose.yml

'@
$gitignorePath = Join-Path $rootPath ".gitignore"
Set-Content -Path $gitignorePath -Value $gitignoreContent -Force -Encoding UTF8
Write-Output "Created .gitignore: $gitignorePath"

# Project initialization complete
Write-Output "Project structure has been created."

# Start Docker containers
Write-Output "Starting Docker containers..."
cd $rootPath
docker-compose up --build