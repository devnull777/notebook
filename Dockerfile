# Используем официальный образ Go для сборки
FROM golang:1.16 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы из текущего местоположения в рабочую директорию внутри контейнера
COPY . .

# Собираем наше приложение
RUN go build -o main .

# Начинаем новую стадию сборки на основе образа alpine
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исполняемый файл приложения и директорию static из стадии builder
COPY --from=builder /app/main /app/main
COPY --from=builder /app/static /app/static

# Запускаем приложение
CMD ["/app/main"]
