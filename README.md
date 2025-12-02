## Описание
Полнофункциональный микросервис на Go для управления пользователями с CRUD-операциями, конкурентной обработкой, rate limiting, метриками Prometheus и контейнеризацией.

## Функциональность
- RESTful API для пользователей: GET /api/users, GET /api/users/{id}, POST /api/users, PUT /api/users/{id}, DELETE /api/users/{id}.
- Конкурентная обработка: асинхронное логирование и уведомления через goroutines.
- Rate limiting: 1000 запросов в секунду.
- Метрики: RPS и latency для Prometheus на /metrics.
- Контейнеризация: Docker и docker-compose.

## Запуск
1. Клонируйте репозиторий: `git clone https://github.com/yourusername/gohighload.git`.
2. Перейдите в папку: `cd gohighload`.
3. Запустите через Docker: `docker-compose up --build`.
4. Сервис доступен на http://localhost:8080.

## Тестирование
- Нагрузочное тестирование: `wrk -t12 -c500 -d60s http://localhost:8080/api/users`.
- Метрики: `curl http://localhost:8080/metrics`.

## Зависимости
- Go 1.24
- gorilla/mux, minio-go, golang.org/x/time/rate, prometheus/client_golang
