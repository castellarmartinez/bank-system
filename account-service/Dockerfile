FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY account-service account-service

WORKDIR /app/account-service

RUN go mod tidy

RUN go build -o account-service ./cmd/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/account-service/account-service .

CMD ["/root/account-service"]
