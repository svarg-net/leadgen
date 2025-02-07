# Stage 1: Сборка приложения
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

COPY config /app/config
COPY migrations /migrations

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz \
    && mv migrate /usr/local/bin/migrate

# Stage 2: Создание финального образа
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache bash

COPY --from=builder /app/main .
COPY --from=builder /app/config /app/config
COPY --from=builder /migrations /migrations
RUN test -d /migrations || (echo "Migrations directory not found" && exit 1)
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

# Запускаем приложение
CMD ["./main"]