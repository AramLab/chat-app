# Chat-app

Простое веб-приложение для обмена сообщениями между пользователями.

## Описание

Это веб-приложение разработано на языке Go с использованием фреймворка Gin и базы данных SQLite. Он предоставляет API для создания пользователей, каналов и отправки сообщений. Все данные сохраняются в базе данных SQLite `database.db`.

## Требования

Для запуска приложения вам потребуется установить следующие компоненты:

- Go (версия 1.16 и выше)
- SQLite (для работы с базой данных)

## Установка

1. Склонируйте репозиторий:

   ```bash
   git clone https://github.com/yourusername/chat-app.git
   cd chat-app
2. Установите зависимости:

   ```bash
   go mod download
3. Запустите приложение:
   
   ```bash
   go run .
Приложение будет доступно по адресу http://localhost:8080.

## Использование API

### Создание пользователя
  ```bash
  curl -X POST http://localhost:8080/users -d '{"username": "username", "password": "password"}' -H "Content-Type: application/json"

### Авторизация пользователя
  ```bash
  curl -X POST http://localhost:8080/login -d '{"username": "username", "password": "password"}' -H "Content-Type: application/json"

### Создание канала
  ```bash
  curl -X POST http://localhost:8080/channels -d '{"name": "general"}' -H "Content-Type: application/json"

### Отправка сообщения

  ```bash
  curl -X POST http://localhost:8080/messages -d '{"channel_id": 1, "user_id": 1, "text": "Hello, world!"}' -H "Content-Type: application/json"

### Получение списка каналов
  ```bash
  curl http://localhost:8080/channels

### Получение сообщений из канала
  ```bash
  curl http://localhost:8080/messages?channelID=1











