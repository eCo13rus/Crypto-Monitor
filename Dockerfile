# Используем официальный образ Golang как исходный для сборки
FROM golang:1.20-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./

# Качаем зависимости
RUN go mod download

# Копируем остальные файлы приложения в контейнер
COPY . .

# Сборка приложения
RUN go build -o /crypto-monitor ./cmd

# Используем минимальный образ для финального контейнера
FROM alpine:latest

WORKDIR /root/

# Копируем скомпилированное приложение из билд-контейнера
COPY --from=builder /crypto-monitor .
# Копируем конфигурационный файл
COPY config.json .

# Устанавливаем зависимости для работы с Redis
RUN apk add --no-cache redis

# Устанавливаем переменную окружения для пути конфигурационного файла
ENV CONFIG_PATH=/root

# Запускаем приложение при старте контейнера
CMD ["./crypto-monitor"]

# Открываем нужный порт (например, 8080)
EXPOSE 8080