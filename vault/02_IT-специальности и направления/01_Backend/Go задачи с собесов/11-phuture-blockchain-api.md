---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/phuture.finance - Компания: Phuture Finance - Уровень: Middle+ - Время на выполнение: 14-20 часов - Ключевые навыки: Blockchain, индексация, ABI decoding, Pos..."
type: concept
---

# Blockchain Indexer API

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/phuture.finance
- **Компания:** Phuture Finance
- **Уровень:** Middle+
- **Время на выполнение:** 14-20 часов
- **Ключевые навыки:** Blockchain, индексация, ABI decoding, PostgreSQL
- **Инфраструктура:** PostgreSQL, Docker, Ethereum RPC

## ТЗ

### Описание
Создать API-сервис для индексации событий блокчейна (EVM). Сервис отслеживает события смарт-контрактов, парсит их из логов и предоставляет REST API для получения проиндексированных данных.

### Функциональные требования

**Эндпоинты:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| GET | /events | Список событий с фильтрами |
| GET | /events/{id} | Детали конкретного события |
| GET | /contracts/{address}/events | События по контракту |
| GET | /sync/status | Статус синхронизации |

**GET /events:**
```
Query params:
- contract_address: фильтр по контракту
- event_name: Transfer, Approval, etc.
- from_block: начальный блок
- to_block: конечный блок
- page, page_size: пагинация

Response:
{
    "events": [
        {
            "id": "evt-uuid",
            "block_number": 18000000,
            "block_hash": "0x...",
            "tx_hash": "0x...",
            "log_index": 5,
            "contract_address": "0x...",
            "event_name": "Transfer",
            "args": {
                "from": "0x...",
                "to": "0x...",
                "value": "1000000000000000000"
            },
            "timestamp": "2024-01-15T10:30:00Z"
        }
    ],
    "total": 1500
}
```

**Бэкграунд индексатор:**
- Подписка на новые блоки через RPC
- Парсинг логов указанных контрактов
- Декодирование event data через ABI
- Сохранение в PostgreSQL

### Нефункциональные требования

1. **Архитектура:** Clean Architecture (handler → service → repository)
2. **Индексация:**
   - Background worker для синхронизации
   - Обработка reorgs (откатов цепочки)
   - Idempotent записи
3. **Конфигурация:**
   ```yaml
   rpc_url: https://eth.llamarpc.com
   contracts:
     - address: 0xA0b86a33E6441E3b4c5f8B5c5c5c5c5c5c5c5c
       abi: ./abi/erc20.json
       from_block: 18000000
   sync_interval: 15s
   ```
4. **API:** REST с пагинацией, фильтрами

### Ограничения
- Только EVM-совместимые цепочки
- PostgreSQL для хранения
- Нет real-time websocket (только polling)

### Бонусные задания
- [ ] Поддержка multiple chains
- [ ] WebSocket для real-time событий
- [ ] GraphQL API
- [ ] Кэширование популярных запросов (Redis)
- [ ] Reorg detection и rollback

## Критерии оценки

- [ ] Background индексация работает
- [ ] Парсинг логов через ABI
- [ ] REST API с фильтрами
- [ ] Пагинация
- [ ] PostgreSQL схема оптимизирована
- [ ] Обработка ошибок и retry
- [ ] Docker и docker-compose
- [ ] Swagger документация
- [ ] Тесты

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: indexer
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"

  indexer:
    build: .
    environment:
      - DATABASE_URL=postgres://user:password@postgres:5432/indexer
      - RPC_URL=${ETH_RPC_URL}
    depends_on:
      - postgres
```

### SQL-схема
```sql
CREATE TABLE indexed_blocks (
    block_number BIGINT PRIMARY KEY,
    block_hash VARCHAR(66) UNIQUE NOT NULL,
    parent_hash VARCHAR(66) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    indexed_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    block_number BIGINT REFERENCES indexed_blocks(block_number),
    block_hash VARCHAR(66) NOT NULL,
    tx_hash VARCHAR(66) NOT NULL,
    log_index INTEGER NOT NULL,
    contract_address VARCHAR(42) NOT NULL,
    event_name VARCHAR(100) NOT NULL,
    args JSONB NOT NULL,
    raw_data TEXT,
    timestamp TIMESTAMP NOT NULL,
    UNIQUE(block_hash, log_index)
);

CREATE INDEX idx_events_contract ON events(contract_address);
CREATE INDEX idx_events_name ON events(event_name);
CREATE INDEX idx_events_block ON events(block_number DESC);
CREATE INDEX idx_events_args ON events USING GIN(args);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет обработки reorgs** — база расходится с реальностью
2. **Не используют ABI** — пытаются парсить raw bytes руками
3. **Нет backoff** — DDoSят RPC при ошибках
4. **Нет idempotency** — дублируют события при рестарте
5. **Плохие индексы** — медленные запросы к events

### Ключевые паттерны
- **Repository pattern** — абстракция над БД
- **Worker pool** — для parallel processing блоков
- **Event decoder** — через go-ethereum/accounts/abi
- **Idempotent insert** — UPSERT или проверка

### Индексация блоков
```go
func (i *Indexer) processBlock(ctx context.Context, blockNum uint64) error {
    // Получаем логи
    logs, err := i.ethClient.FilterLogs(ctx, ethereum.FilterQuery{
        FromBlock: big.NewInt(int64(blockNum)),
        ToBlock:   big.NewInt(int64(blockNum)),
        Addresses: i.config.Contracts,
    })
    if err != nil { return err }
    
    // Парсим каждый лог
    for _, log := range logs {
        event, err := i.parseLog(log)
        if err != nil { continue }
        
        // Сохраняем
        if err := i.repo.SaveEvent(ctx, event); err != nil {
            return err
        }
    }
    return nil
}
```

### Follow-up вопросы
1. Как обнаружить и обработать chain reorg?
2. Как масштабировать до 100 контрактов?
3. Что если RPC пропустил блок?
4. Как реализовать real-time уведомления?

### Уровневые подсказки
- 🟢 **Направление:** Нужен background worker который polling'ом забирает блоки и парсит логи
- 🟡 **Конкретика:** Используй go-ethereum клиент и abi.Unpack для декодирования событий
- 🔴 **Фрагмент:** ```go
func (i *Indexer) parseLog(log types.Log) (*Event, error) {
    abi, _ := abi.JSON(strings.NewReader(erc20ABI))
    event := Event{
        BlockNumber: log.BlockNumber,
        TxHash:      log.TxHash.Hex(),
        Contract:    log.Address.Hex(),
    }
    
    // Декодируем топики и data
    if len(log.Topics) > 0 {
        sig := log.Topics[0].Hex()
        event.EventName = i.knownEvents[sig]
    }
    return &event, nil
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L5 — Context, background workers
- L6 — HTTP, JSON
- L7 — Clean architecture
- L8 — PostgreSQL, Docker

### Новые навыки
- Blockchain indexing
- ABI decoding
- Event-driven parsing
- Handling blockchain specifics (reorgs)

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

