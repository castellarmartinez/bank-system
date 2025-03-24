FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY transaction-service transaction-service

WORKDIR /app/transaction-service

RUN go mod tidy

RUN go build -o transaction-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/transaction-service/transaction-service .

CMD ["/root/transaction-service"]