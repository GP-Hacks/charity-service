FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY charity/go.mod charity/go.sum ./
RUN go mod download

COPY . .

# Собираем сервис
WORKDIR /app/charity/cmd/charity
RUN go build -o charity_service

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/charity/cmd/charity/charity_service .
COPY --from=builder /app/charity/config/config.yaml ./config/config.yaml

EXPOSE 8080

CMD ["./charity_service"]
