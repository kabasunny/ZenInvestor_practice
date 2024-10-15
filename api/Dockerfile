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
