---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/bitburst - Компания: Bitburst - Уровень: Middle - Время на выполнение: 10-14 часов - Ключевые навыки: PostgreSQL, конкурентность, worker pool, graceful shutdo..."
type: concept
---

# Callback Service

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/bitburst
- **Компания:** Bitburst
- **Уровень:** Middle
- **Время на выполнение:** 10-14 часов
- **Ключевые навыки:** PostgreSQL, конкурентность, worker pool, graceful shutdown
- **Инфраструктура:** PostgreSQL, Docker

## ТЗ

### Описание
Создать REST-сервис для обработки callback-запросов. Сервис принимает списки объектов, асинхронно запрашивает их статус из внешнего API и сохраняет в PostgreSQL. Объекты, не появляющиеся более 30 секунд, должны удаляться.

### Функциональные требования

**Эндпоинты:**

| Метод | URL | Тело запроса | Описание |
|-------|-----|--------------|----------|
| POST | `/callback` | `{"object_ids": [1,2,3,4,5,6]}` | Принять callback со списком ID |

**Внешний сервис (предоставляется):**
- URL: `localhost:9010/objects/:id`
- Response time: 300ms - 4s (непредсказуемо!)
- Response: `{"id": <id>, "online": true|false}`

**Бизнес-логика:**
1. Для каждого `object_id` запросить статус из внешнего API
2. Отфильтровать объекты где `online: true`
3. Сохранить в PostgreSQL с timestamp последнего появления
4. Удалять объекты, не полученные более 30 секунд

**Важно:**
- Нельзя пропускать callback'и (business constraint)
- До 200 ID в одном запросе
- Callback приходят каждые 5 секунд

### Нефункциональные требования

1. **Конкурентность:**
   - Параллельная обработка ID (worker pool)
   - Таймауты на внешние запросы
   - Graceful shutdown с завершением текущих задач

2. **База данных:**
   - PostgreSQL для хранения объектов
   - Индекс по `last_seen` для cleanup

3. **Надёжность:**
   - Не терять callback'и при рестарте
   - Обработка паник в воркерах

### API Внешнего сервиса

```
GET localhost:9010/objects/:id

Response:
{
    "id": 123,
    "online": true  // или false
}

Latency: 300ms - 4000ms (random)
```

### Ограничения
- Нельзя использовать существующие cron parser библиотеки [дополнено: это из другого задания, здесь нет]
- Нельзя терять callback'и
- Обязательно использовать PostgreSQL

### Бонусные задания
- [ ] Graceful shutdown с drain
- [ ] Retry logic для внешних запросов
- [ ] Метрики (prometheus)
- [ ] Cleanup worker как отдельный goroutine

## Критерии оценки

- [ ] Корректность обработки callback'ов
- [ ] Потокобезопасность (нет race conditions)
- [ ] Worker pool для параллельной обработки
- [ ] Graceful shutdown
- [ ] Тесты (unit + integration)
- [ ] Docker и docker-compose
- [ ] Структура проекта
- [ ] Обработка ошибок (нет panic в проде)
- [ ] **Code style:** Linter, gofmt, читаемость кода (как в ревью — «читаемый как книга»)

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: callback
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
  
  app:
    build: .
    depends_on:
      - postgres
    environment:
      DATABASE_URL: postgres://user:password@postgres:5432/callback
      EXTERNAL_API_URL: http://localhost:9010
```

### SQL-схема
```sql
CREATE TABLE objects (
    id SERIAL PRIMARY KEY,
    object_id INTEGER UNIQUE NOT NULL,
    online BOOLEAN NOT NULL,
    last_seen TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_last_seen ON objects(last_seen);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Синхронная обработка** — 200 ID × 4s = 800s = disaster
2. **Нет таймаутов** — висящие goroutines
3. **Race condition** — одновременный доступ к мапе/счётчику
4. **Нет cleanup** — база растёт бесконечно
5. **Утечка goroutines** — не закрывают каналы

### Ключевые паттерны
- **Worker pool** — ограничить concurrency (10-20 workers)
- **Context with timeout** — для внешних запросов
- **errgroup** — для ожидания всех воркеров
- **Background cleanup job** — ticker каждые 10s

### Follow-up вопросы
1. Как бы ты масштабировал это до 10k ID за раз?
2. Что если внешний API падает на 30 секунд?
3. Как бы ты сделал "exactly once" обработку при рестарте?
4. Как обнаружить deadlock в worker pool?

### Уровневые подсказки
- 🟢 **Направление:** Нужен worker pool — горутины, которые читают ID из канала и обрабатывают
- 🟡 **Конкретика:** Используй `context.WithTimeout` для каждого HTTP-запроса
- 🔴 **Фрагмент:** ```go
func (s *Service) ProcessBatch(ids []int) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    g, ctx := errgroup.WithContext(ctx)
    sem := make(chan struct{}, 10) // max 10 concurrent
    
    for _, id := range ids {
        id := id // capture
        g.Go(func() error {
            sem <- struct{}{}
            defer func() { <-sem }()
            return s.processOne(ctx, id)
        })
    }
    return g.Wait()
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L4 — Горутины, каналы, mutex
- L5 — Context, errgroup, graceful shutdown
- L6 — HTTP клиент/сервер
- L8 — PostgreSQL, Docker

### Новые навыки
- Worker pool паттерн
- Graceful shutdown с drain
- Интеграция с внешними API
- Background jobs

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

