FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем сервис
WORKDIR /app/cmd/charity
RUN go build -o charity_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/cmd/charity/charity_service .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["./charity_service"]

