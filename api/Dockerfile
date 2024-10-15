# Dockerfile
FROM golang:1.23

WORKDIR /app/api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# srcチE��レクトリを正しい場所にコピ�E
COPY src/ ./

# golang と air のインスト�Eル
RUN apt-get update && apt-get install -y golang-go
RUN go install github.com/air-verse/air@v1.61.0

ENV PATH="/go/bin:/usr/local/go/bin:"

EXPOSE 8080
CMD ["air"]


