# Используем официальный образ Go для сборки и выполнения приложения
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Устанавливаем Air для live-reload (только для процесса разработки)
RUN go install github.com/air-verse/air@latest

# Копируем исходный код
COPY . ./

# Собираем приложение
RUN go build -o server ./cmd/server

# Запускаем приложение с использованием Air для live-reload (если требуется для разработки)
CMD ["air", "-c", "/app/.air.toml"]
