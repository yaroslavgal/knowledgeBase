---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/codenotary - Компания: CodeNotary - Уровень: Middle - Время на выполнение: 10-14 часов - Ключевые навыки: gRPC, REST API, ImmuDB, Docker, clean architecture -..."
type: concept
---

# ImmuDB Log API

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/codenotary
- **Компания:** CodeNotary
- **Уровень:** Middle
- **Время на выполнение:** 10-14 часов
- **Ключевые навыки:** gRPC, REST API, ImmuDB, Docker, clean architecture
- **Инфраструктура:** ImmuDB, Docker

## ТЗ

### Описание
Создать сервис для хранения и извлечения логов с использованием ImmuDB (иммутабельная база данных). Сервис предоставляет REST/gRPC API для записи логов и их просмотра с пагинацией.

### Функциональные требования

**API методы:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| POST | /logs | Сохранить одну строку лога |
| POST | /logs/batch | Сохранить batch логов |
| GET | /logs | Получить историю логов (с пагинацией) |
| GET | /logs/last?n=10 | Получить последние N логов |
| GET | /logs/count | Получить общее количество логов |

**gRPC сервис:**
```protobuf
service LogService {
    rpc StoreLog(LogEntry) returns (StoreResponse);
    rpc StoreBatch(BatchRequest) returns (StoreResponse);
    rpc GetHistory(HistoryRequest) returns (HistoryResponse);
    rpc GetCount(Empty) returns (CountResponse);
}
```

**Бизнес-логика:**
1. Каждый лог содержит: timestamp, level (INFO/WARN/ERROR), message, source (опционально)
2. ImmuDB гарантирует иммутабельность — логи нельзя изменить или удалить
3. Поддержка "buckets" — разделение логов по приложениям/источникам
4. Batch запросы для эффективной записи множества логов

### Нефункциональные требования

1. **Архитектура:** Clean Architecture (handler → service → repository)
2. **API:** gRPC + gRPC-gateway для REST
3. **База данных:** ImmuDB в SQL-режиме
4. **Docker:** Полный docker-compose для запуска
5. **Тестирование:** Postman коллекция или curl примеры

### API Спецификация

**POST /logs**
```json
{
    "timestamp": "2024-01-15T10:30:00Z",
    "level": "INFO",
    "message": "User logged in",
    "source": "auth-service",
    "bucket": "production"
}
```

**GET /logs?bucket=production&page=1&page_size=50**
```json
{
    "logs": [
        {
            "timestamp": "2024-01-15T10:30:00Z",
            "level": "INFO",
            "message": "User logged in",
            "source": "auth-service"
        }
    ],
    "total": 150,
    "page": 1
}
```

### Ограничения
- ImmuDB в SQL-режиме (проще для замены на другую БД)
- gRPC-gateway для REST API
- Docker-compose обязателен

### Бонусные задания
- [ ] Аутентификация (JWT или API keys)
- [ ] Разделение прав (read-only vs write)
- [ ] Логирование самого сервиса
- [ ] Health checks

## Критерии оценки

- [ ] Корректная работа с ImmuDB
- [ ] gRPC + REST API (gateway)
- [ ] Clean architecture
- [ ] Docker и docker-compose
- [ ] Batch операции
- [ ] Пагинация
- [ ] Тесты
- [ ] Postman коллекция
- [ ] README с инструкциями

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  immudb:
    image: codenotary/immudb:latest
    ports:
      - "3322:3322"
    environment:
      - IMMUDB_ADDRESS=0.0.0.0
    volumes:
      - immudb_data:/var/lib/immudb

  logapi:
    build: .
    ports:
      - "8080:8080"
      - "50051:50051"
    depends_on:
      - immudb
    environment:
      - IMMUDB_HOST=immudb:3322

volumes:
  immudb_data:
```

### SQL-схема (ImmuDB)
```sql
CREATE TABLE IF NOT EXISTS loglines (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    level VARCHAR(10) NOT NULL,
    message TEXT NOT NULL,
    source VARCHAR(255),
    bucket VARCHAR(255) DEFAULT 'default'
);

CREATE INDEX idx_timestamp ON loglines(timestamp);
CREATE INDEX idx_bucket ON loglines(bucket);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет batch** — вставляют логи по одному (медленно)
2. **Нет пагинации** — выгружают все логи сразу
3. **Путают ImmuDB API** — используют неправильный клиент
4. **Нет graceful shutdown** — теряют логи при остановке
5. **Нет индексов** — медленный поиск по времени

### Ключевые паттерны
- **Repository pattern** — абстракция над ImmuDB
- **gRPC-gateway** — единый код для gRPC и REST
- **Batch insert** — эффективная запись
- **Cursor-based pagination** — для больших данных

### Follow-up вопросы
1. Почему ImmuDB а не обычная PostgreSQL?
2. Как бы ты масштабировал до 1M логов в секунду?
3. Как реализовать ротацию старых логов в иммутабельной БД?
4. Что если ImmuDB недоступна — как не потерять логи?

### Уровневые подсказки
- 🟢 **Направление:** Нужен gRPC сервис с gateway + репозиторий для ImmuDB
- 🟡 **Конкретика:** Используй github.com/codenotary/immudb/pkg/client
- 🔴 **Фрагмент:** ```go
type LogRepository struct {
    client immudb.ImmuClient
}

func (r *LogRepository) Store(ctx context.Context, log LogEntry) error {
    _, err := r.client.SQLExec(ctx, 
        "INSERT INTO loglines (timestamp, level, message) VALUES (?, ?, ?)",
        log.Timestamp, log.Level, log.Message,
    )
    return err
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L5 — gRPC, protobuf
- L6 — HTTP, middleware
- L7 — Clean architecture
- L8 — Docker, specialized DB

### Новые навыки
- Работа с ImmuDB
- gRPC-gateway
- Иммутабельные данные

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

