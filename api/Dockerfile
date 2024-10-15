FROM golang:1.23

WORKDIR /app/api

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# air縺ｮ繧､繝ｳ繧ｹ繝医・繝ｫ縺ｨPATH縺ｫ霑ｽ蜉
RUN go install github.com/air-verse/air@latest

# PATH繧定ｨｭ螳壹☆繧九◆繧√↓CMD繧堤峩謗･螳溯｡・
ENV PATH="/root/go/bin:"

EXPOSE 8080
CMD ["air"]
