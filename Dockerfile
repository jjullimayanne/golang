
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-reciclagem ./cmd/main.go

FROM debian:buster

WORKDIR /app

COPY --from=builder /app/api-reciclagem .

EXPOSE 8080

CMD ["./api-reciclagem"]
