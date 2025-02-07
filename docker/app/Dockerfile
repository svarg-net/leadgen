# Stage 1: Сборка приложения
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go get github.com/lib/pq
RUN go get gopkg.in/yaml.v3
# Копируем исходный

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go
COPY config /app/config
# Stage 2: Создание финального образа
FROM alpine:latest

WORKDIR /app

# Устанавливаем необходимые утилиты
RUN apk add --no-cache bash

# Копируем бинарник из первого этапа
COPY --from=builder /app/main .
COPY --from=builder /app/config /app/config
# Запускаем приложение
CMD ["./main"]