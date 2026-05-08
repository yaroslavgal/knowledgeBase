---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/kima-finance - Компания: Kima Finance - Уровень: Middle+ - Время на выполнение: 12-16 часов - Ключевые навыки: gRPC, protobuf, кэширование, конкурентность - И..."
type: concept
---

# Currency Conversion gRPC Service

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/kima-finance
- **Компания:** Kima Finance
- **Уровень:** Middle+
- **Время на выполнение:** 12-16 часов
- **Ключевые навыки:** gRPC, protobuf, кэширование, конкурентность
- **Инфраструктура:** Docker (опционально внешний API)

## ТЗ

### Описание
Создать gRPC микросервис для конвертации валют с поддержкой batch операций, пагинацией и синхронизацией курсов из внешних API.

### Задание состоит из 2 частей

#### Часть 1: gRPC сервис

**Proto файл:**
```protobuf
service CurrencyService {
    // Конвертация одной валюты
    rpc Convert(ConvertRequest) returns (ConvertResponse);
    
    // Batch конвертация
    rpc ConvertBatch(ConvertBatchRequest) returns (ConvertBatchResponse);
    
    // Список курсов с пагинацией
    rpc ListRates(ListRatesRequest) returns (ListRatesResponse);
}

message ConvertRequest {
    string from_currency = 1;
    string to_currency = 2;
    double amount = 3;
}

message ConvertResponse {
    double converted_amount = 1;
    double rate = 2;
    google.protobuf.Timestamp timestamp = 3;
}

message ConvertBatchRequest {
    repeated ConvertRequest requests = 1;
}

message ConvertBatchResponse {
    repeated ConvertResponse responses = 1;
}

message ListRatesRequest {
    int32 page = 1;
    int32 page_size = 2;
}

message ListRatesResponse {
    repeated Rate rates = 1;
    int32 total = 2;
}

message Rate {
    string from_currency = 1;
    string to_currency = 2;
    double rate = 3;
    google.protobuf.Timestamp updated_at = 4;
}
```

**Требования:**
- Описать proto (без реализации gRPC методов)
- Реализация `GetRates` — получение курсов из memory
- Поддержка множественных источников (CoinGecko, CurrencyLayer)

#### Часть 2: Синхронизация курсов

**Архитектура:**
```
Внешний API (CoinGecko/CurrencyLayer)
         ↓
   Fetcher (интерфейс)
         ↓
   In-Memory Cache
         ↓
   gRPC Service
```

**Интерфейс Fetcher:**
```go
type PriceFetcher interface {
    FetchRates(ctx context.Context) (map[string]float64, error)
    Source() string // "coingecko" или "currencylayer"
}
```

**Требования:**
1. Конкурентная синхронизация курсов в память
2. Поддержка fallback между источниками
3. Потокобезопасный доступ к кэшу
4. Периодическое обновление (background sync)

### Нефункциональные требования

1. **Конкурентность:**
   - Несколько fetcher'ов работают параллельно
   - Безопасный доступ к shared cache

2. **Конфигурация:**
   ```yaml
   app:
     sync_interval: 60s
     fallback_enabled: true
   
   coingecko:
     api_key: "..."
     base_url: "https://api.coingecko.com"
   
   currencylayer:
     api_key: "..."
     base_url: "https://api.currencylayer.com"
   ```

3. **Makefile:**
   - `make run` — локальный запуск
   - `make lint && make test` — проверки

### Бонусные задания
- [ ] gRPC Gateway для REST API
- [ ] Prometheus метрики
- [ ] Circuit breaker для внешних API
- [ ] Graceful shutdown

## Критерии оценки

- [ ] Корректный proto файл
- [ ] Интерфейс PriceFetcher
- [ ] Конкурентная синхронизация
- [ ] Потокобезопасный кэш
- [ ] Fallback между источниками
- [ ] Тесты (unit + integration)
- [ ] Docker

## Инфраструктура

### Docker Compose (опционально)
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "50051:50051"  # gRPC
      - "8080:8080"    # HTTP gateway (optional)
```

### Структура проекта
```
kima-finance/
├── api/
│   └── service.proto       # protobuf определения
├── internal/
│   ├── service/
│   │   └── pricer/
│   │       └── service.go  # GetRates
│   └── service/fetchers/
│       ├── abstract.go     # PriceFetcher interface
│       ├── coingecko.go    # реализация
│       └── currencylayer.go
├── configs/
│   └── app_conf.yml
├── Makefile
└── README.md
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет интерфейса** — хардкод конкретного fetcher
2. **Data race** — обновляют кэш без синхронизации
3. **Блокировка** — долгий fetch блокирует все запросы
4. **Нет fallback** — один упал = всё упало
5. **Memory leak** — неограниченный кэш

### Ключевые паттерны

**Fetch с fallback:**
```go
func (s *Service) syncRates(ctx context.Context) {
    for _, fetcher := range s.fetchers {
        rates, err := fetcher.FetchRates(ctx)
        if err != nil {
            log.Printf("%s failed: %v", fetcher.Source(), err)
            continue // пробуем следующий
        }
        s.updateCache(rates)
        return // успех
    }
}
```

**Потокобезопасный кэш:**
```go
type RateCache struct {
    mu     sync.RWMutex
    rates  map[string]float64
    updated time.Time
}

func (c *RateCache) Get(from, to string) (float64, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    rate, ok := c.rates[from+"_"+to]
    return rate, ok
}
```

### Follow-up вопросы
1. Как бы ты добавил Redis для кэша?
2. Что если курсы нужны с задержкой < 100ms?
3. Как обработать rate limiting внешних API?
4. Как тестировать без реальных API ключей?

### Уровневые подсказки
- 🟢 **Направление:** Нужен интерфейс для fetcher'ов и фоновая горутина для синхронизации
- 🟡 **Конкретика:** Используй `sync.RWMutex` для кэша и `time.Ticker` для периодического обновления
- 🔴 **Фрагмент:** ```go
func (s *Service) StartSync(ctx context.Context) {
    ticker := time.NewTicker(s.config.SyncInterval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            s.syncOnce(ctx)
        case <-ctx.Done():
            return
        }
    }
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L4 — Конкурентность, мьютексы
- L5 — Context, graceful shutdown
- L8 — gRPC, protobuf, Docker

### Новые навыки
- gRPC сервисы
- Интеграция с внешними API
- Background jobs
- Failover паттерны

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

