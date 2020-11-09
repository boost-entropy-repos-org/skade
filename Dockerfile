FROM golang:1.15-alpine3.12

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main cmd/skade/main.go
CMD ["/app/main"]
