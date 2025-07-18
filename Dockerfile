# Используем официальный образ Golang версии 1.24 для сборки приложения
FROM golang:1.24 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы зависимостей go.mod и go.sum в контейнер
COPY go.mod go.sum ./

# Загружаем зависимости (кэшируем на этом этапе, чтобы не скачивать заново при отсутствии изменений)
RUN go mod download

# Копируем весь исходный код приложения в контейнер
COPY . .

# Устанавливаем инструмент swag для генерации Swagger документации
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Генерируем Swagger документацию на основе файла cmd/main.go
RUN swag init -g cmd/myapp/main.go

# Собираем исполняемый файл приложения с помощью команды go build
RUN go build -o proxy-api ./cmd/myapp/

# --- Второй этап: создаём минимальный образ для запуска приложения ---

# Используем облегчённый Debian образ для уменьшения итогового размера
FROM debian:bookworm-slim

# Устанавливаем рабочую директорию для финального контейнера
WORKDIR /app

# Обновляем индекс пакетов и устанавливаем необходимые системные пакеты:
# - ca-certificates для HTTPS-сертификатов
# - tzdata для временных зон
RUN apt-get update && apt-get install -y ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

# Копируем собранный бинарник из стадии сборки в текущий контейнер
COPY --from=builder /app/proxy-api /app/proxy-api

# Копируем дополнительные конфигурационные файлы
COPY .env ./
COPY config.yaml ./

# Открываем порт 8080 для внешнего доступа к приложению
EXPOSE 8080

# Команда запуска приложения при старте контейнера
CMD ["./proxy-api"]