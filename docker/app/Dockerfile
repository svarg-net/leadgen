# Stage 1: Сборка приложения
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

# Stage 2: Создание финального образа
FROM alpine:latest

WORKDIR /app

# Устанавливаем необходимые утилиты
RUN apk add --no-cache bash

# Копируем бинарник из первого этапа
COPY --from=builder /app/main .

# Запускаем приложение
CMD ["./main"]