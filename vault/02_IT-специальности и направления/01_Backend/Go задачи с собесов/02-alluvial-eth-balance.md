---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/alluvial - Компания: Alluvial - Уровень: Middle - Время на выполнение: 8-12 часов - Ключевые навыки: HTTP API, Ethereum RPC, балансировка нагрузки, Docker - И..."
type: concept
---

# Ethereum Balance Checker

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/alluvial
- **Компания:** Alluvial
- **Уровень:** Middle
- **Время на выполнение:** 8-12 часов
- **Ключевые навыки:** HTTP API, Ethereum RPC, балансировка нагрузки, Docker
- **Инфраструктура:** Docker, без внешних зависимостей (zero external deps)

## ТЗ

### Описание
Создать микросервис для получения баланса Ethereum-адресов. Сервис должен поддерживать работу с несколькими RPC-нодами, балансировать нагрузку между ними и обходить rate limiting бесплатных провайдеров.

### Функциональные требования

**Эндпоинты:**

| Метод | URL | Описание |
|-------|-----|----------|
| GET | `/eth/balance/{address}` | Получить ETH баланс адреса |
| GET | `/eth/{token}/balance/{address}` | Получить баланс ERC20 токена |
| GET | `/health` | Health check (всегда `ok`) |
| GET | `/ready` | Readiness check (всегда `ok`) |
| GET | `/metrics` | Метрики приложения |

**Примеры:**
- `http://localhost:8000/eth/balance/0xDf8ac28156209F5cbc89cDec419dd3dB3D3E326a`
- `http://localhost:8000/eth/usdc/balance/0xDf8ac28156209F5cbc89cDec419dd3dB3D3E326a`

### Нефункциональные требования

1. **Балансировка RPC-нод:**
   - Ротация между несколькими бесплатными RPC-провайдерами
   - Обработка ограничений rate limit
   - Fallback между нодами при ошибках

2. **Метрики:**
   - Метрики балансировщика с префиксом `balancer_proxy`
   - Endpoint `/metrics` для Prometheus-стиля метрик

3. **Graceful degradation:**
   - Zero external dependencies для health/ready endpoints
   - Всегда возвращают `ok` даже если RPC-ноды недоступны

### Ограничения
- Использовать несколько бесплатных RPC-нод (Infura, Alchemy, QuickNode и т.д.)
- Без кэширования (можно добавить как бонус)
- Docker для запуска

### Бонусные задания
- [ ] Кэширование известных адресов
- [ ] Отслеживание изменений через новые транзакции
- [ ] Поддержка множества ERC20 токенов из конфига

## Критерии оценки

- [ ] Корректность получения баланса ETH и ERC20
- [ ] Балансировка между RPC-нодами
- [ ] Обработка ошибок и fallback
- [ ] Метрики и мониторинг
- [ ] Структура проекта (clean architecture)
- [ ] Docker и Makefile
- [ ] Тесты (unit + integration)
- [ ] Документация (README)

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "8000:8000"
    environment:
      - RPC_NODES=https://eth.llamarpc.com,https://rpc.ankr.com/eth
```

### Makefile цели
```makefile
test      # запуск тестов
lint      # проверка линтером
dev_up    # запуск в Docker
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Не используют retry** — при 429 ошибке сразу паникуют
2. **Синхронные запросы к нодам** — не используют таймауты
3. **Хардкод RPC URL** — нет конфигурации через env
4. **Нет graceful shutdown** — при остановке теряются соединения

### Ключевые паттерны
- **Round-robin / Weighted round-robin** — для балансировки
- **Circuit breaker** — для отключения "больных" нод
- **Retry with backoff** — для обработки rate limits
- **Interface for RPC client** — для тестирования

### Follow-up вопросы
1. Как бы ты добавил кэширование? Что кэшировать и на какое время?
2. Как обработать случай когда все RPC-ноды недоступны?
3. Как бы ты мониторил "здоровье" каждой ноды?
4. Что делать если один пользователь DDoSит endpoint одним адресом?

### Уровневые подсказки
- 🟢 **Направление:** Нужно создать обёртку над RPC-клиентом, которая будет чередовать URL
- 🟡 **Конкретика:** Используй `sync.Mutex` или `atomic` для счётчика текущей ноды
- 🔴 **Фрагмент:** ```go
type Balancer struct {
    nodes []string
    current uint32
}
func (b *Balancer) Next() string {
    idx := atomic.AddUint32(&b.current, 1) % uint32(len(b.nodes))
    return b.nodes[idx]
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L6 — HTTP сервер, middleware
- L7 — Структура проекта, clean architecture
- L8 — Docker, PostgreSQL (опционально)

### Новые навыки
- Интеграция с Ethereum RPC
- Балансировка нагрузки
- Graceful degradation паттерны

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

