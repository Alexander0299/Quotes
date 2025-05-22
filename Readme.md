1. Запуск файла:
В терминале ввести команду:

go run main.go

В терминале должны увидеть : "Сервер запущен на порту: 8080"

2. Запустить дополнительный терминал. В дополнительном терминале ввести команды:
Для Linux/macOS:
# Добавить цитату
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

# Получить все цитаты
curl http://localhost:8080/quotes

# Получить случайную цитату
curl http://localhost:8080/quotes/random

# Фильтрация по автору
curl http://localhost:8080/quotes?author=Confucius

# Удалить цитату по ID
curl -X DELETE http://localhost:8080/quotes/1

///////////////////////
 Для Windows:
# Добавить цитату

Invoke-RestMethod -Method POST http://localhost:8080/quotes `
  -ContentType "application/json" `
  -Body '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

  # Получить все цитаты
  Invoke-RestMethod -Method GET http://localhost:8080/quotes

  # Получить случайную цитату
  Invoke-RestMethod -Method GET http://localhost:8080/quotes/random

  # Фильтрация по автору
  Invoke-RestMethod -Method GET "http://localhost:8080/quotes?author=Confucius"

  # Удалить цитату по ID
  Invoke-RestMethod -Method DELETE http://localhost:8080/quotes/1