---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/field-materials - Компания: Field Materials - Уровень: Junior/Middle - Время на выполнение: 6-10 часов - Ключевые навыки: HTTP API, обработка изображений, asy..."
type: concept
---

# Image Resizer API

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/field-materials
- **Компания:** Field Materials
- **Уровень:** Junior/Middle
- **Время на выполнение:** 6-10 часов
- **Ключевые навыки:** HTTP API, обработка изображений, async/sync обработка
- **Инфраструктура:** Нет (in-memory или файловая система)

## ТЗ

### Описание
Создать API-сервис для ресайза изображений. Поддержка синхронной и асинхронной обработки, возврат URL для скачивания resized изображений.

### Функциональные требования

**Эндпоинты:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| POST | /v1/resize | Синхронный ресайз |
| POST | /v1/resize/async | Асинхронный ресайз |
| GET | /v1/resize/{jobId}/status | Проверка статуса async задачи |
| GET | /v1/download/{filename} | Скачать обработанное изображение |

**POST /v1/resize (sync):**
```json
{
    "image_url": "https://example.com/image.jpg",
    "width": 800,
    "height": 600,
    "format": "jpeg",
    "quality": 85
}

Response:
{
    "url": "http://localhost:8080/v1/download/resized_abc123.jpg",
    "width": 800,
    "height": 600,
    "processing_time_ms": 150
}
```

**POST /v1/resize/async:**
```json
{
    "image_url": "https://example.com/large-image.jpg",
    "width": 1920,
    "height": 1080
}

Response:
{
    "job_id": "job-uuid",
    "status": "pending",
    "check_url": "http://localhost:8080/v1/resize/job-uuid/status"
}
```

**GET /v1/resize/{jobId}/status:**
```json
{
    "job_id": "job-uuid",
    "status": "completed",
    "result_url": "http://localhost:8080/v1/download/resized_def456.jpg"
}
```

### Нефункциональные требования

1. **Форматы:** JPEG, PNG, WebP (опционально)
2. **Обработка:**
   - Sync — для быстрых операций (< 5 сек)
   - Async — для больших файлов (background processing)
3. **Хранение:** Локальная файловая система или in-memory
4. **Валидация:**
   - Проверка URL (должен быть валидным)
   - Лимиты на размер (max 10MB)
   - Поддерживаемые форматы
5. **Cleanup:** Удаление старых файлов (TTL 1 час)

### Ограничения
- In-memory очередь для async задач (без Redis/RabbitMQ)
- Локальное хранилище (без S3)
- Нет аутентификации

### Бонусные задания
- [ ] Поддержка WebP
- [ ] Crop mode (cover, contain, fill)
- [ ] Watermark overlay
- [ ] Bulk resize (несколько размеров за раз)
- [ ] Caching частых запросов

## Критерии оценки

- [ ] Синхронный ресайз работает
- [ ] Асинхронный ресайз работает
- [ ] Проверка статуса async задачи
- [ ] Скачивание результата
- [ ] Валидация входных данных
- [ ] Обработка ошибок (невалидный URL, timeout)
- [ ] Тесты
- [ ] Makefile с командами
- [ ] README с примерами

## Инфраструктура

### Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache gcc musl-dev libpng-dev libjpeg-turbo-dev
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates libpng libjpeg-turbo
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

### Структура проекта
```
field-materials/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   │   └── resize.go
│   ├── service/
│   │   └── resizer.go
│   ├── processor/
│   │   ├── sync.go
│   │   └── async.go
│   └── storage/
│       └── local.go
├── uploads/              # для хранения файлов
├── Makefile
└── README.md
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Утечка памяти** — не закрывают HTTP response body
2. **Нет таймаутов** — висят на больших файлах
3. **Race condition** — concurrent map access для job status
4. **Disk overflow** — не чистят старые файлы
5. **Нет валидации** — принимают любые URL (SSRF риск)

### Ключевые паттерны
- **Worker pool** — для async обработки
- **Job queue** — канал для задач
- **Streaming download** — не грузить весь файл в память
- **Semaphore** — ограничить concurrent обработку

### Async обработка
```go
type Job struct {
    ID       string
    URL      string
    Width    int
    Height   int
    Status   string // pending, processing, completed, failed
    ResultURL string
}

type AsyncProcessor struct {
    queue chan Job
    jobs  map[string]*Job
    mu    sync.RWMutex
}

func (p *AsyncProcessor) Start(workers int) {
    for i := 0; i < workers; i++ {
        go p.worker()
    }
}
```

### Follow-up вопросы
1. Как бы ты масштабировал до 1000 запросов/сек?
2. Что если оригинальный URL недоступен?
3. Как защититься от SSRF (запросы к internal IPs)?
4. Как реализовать прогресс-бар для долгих задач?

### Уровневые подсказки
- 🟢 **Направление:** Для async используй канал как очередь и горутины-воркеры
- 🟡 **Конкретика:** Создай структуру Job с ID, статусом, и храни в мапе
- 🔴 **Фрагмент:** ```go
func (h *Handler) AsyncResize(w http.ResponseWriter, r *http.Request) {
    job := Job{ID: uuid.New().String(), Status: "pending"}
    h.processor.Submit(job)
    
    json.NewEncoder(w).Encode(map[string]string{
        "job_id": job.ID,
        "check_url": "/v1/resize/" + job.ID + "/status",
    })
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L6 — HTTP сервер, JSON
- L4 — Горутины, каналы (для async)

### Новые навыки
- Обработка изображений
- Async job processing
- File I/O в вебе

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

