# builder
FROM golang:1.23 AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/transport ./main.go

# runtime
FROM debian:bookworm-slim
WORKDIR /app

# Установим pg_isready и базовые утилиты
RUN apt-get update && apt-get install -y ca-certificates curl postgresql-client && rm -rf /var/lib/apt/lists/*

# Установим golang-migrate CLI
ARG MIGRATE_VERSION=v4.15.2
RUN curl -L "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz" \
    | tar xz -C /usr/local/bin && chmod +x /usr/local/bin/migrate

# Копируем бинарник приложения
COPY --from=builder /app/transport /app/transport

# Копируем entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Место для монтирования миграций
VOLUME ["/migrations"]

EXPOSE 8080

ENV PORT=8080

ENTRYPOINT ["/app/entrypoint.sh"]
