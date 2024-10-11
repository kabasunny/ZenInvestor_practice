FROM golang:1.20
WORKDIR /app/api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080
CMD ["air"]