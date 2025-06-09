# UrlShorter

Простой сервис для сокращения URL.  
Позволяет создавать короткие ссылки, которые перенаправляют на исходные длинные URL.

---

## Описание

UrlShorter — это минималистичный URL shortener, реализованный на Go с использованием PostgreSQL в качестве базы данных.  
Поддерживает создание коротких ссылок и редирект по ним.

---

## Функциональность

- Генерация коротких URL для длинных ссылок
- Сохранение ссылок в базе данных PostgreSQL
- Перенаправление с короткой ссылки на оригинальную

---

## Технологии

- Go (pgx для работы с PostgreSQL)
- PostgreSQL
- Миграции базы данных через `migrate`
- Gorilla/mux router

---

### Требования

- Go 1.20+  
- PostgreSQL  
- Утилита migrate (https://github.com/golang-migrate/migrate)
- Роутер gorilla/mux (https://github.com/gorilla/mux)

---

### Запуск сервера
1. Клонируйте репозиторий:
```bash
git clone https://github.com/folada2007/UrlShorter.git
cd UrlShorter
```
2. Установите зависимости:
```bash
go mod tidy
```
3. Запустите приложение:
```bash
go run main.go
```

---

### Настройка базы данных

1. Создайте базу данных PostgreSQL:

```sql
CREATE DATABASE shorterdb;
```

2. Запустите миграции
```bash
migrate -source file://internal/db/migrations -database "postgres://postgres:yourpassword@localhost:5432/shorterdb?search_path=public" up
```

---

# API

### Создание короткой ссылки
```bash
POST "/"
Content-Type: application/json

{
  "long_url": "https://example.com/very/long/url"
}
```
Ответ:
```bash
{
  "short_url": "http://localhost:8080/abc123"
}
```
Переход по короткой ссылке
```bash
GET /{short_url}
```

