---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/trust-wallet - Компания: Trust Wallet - Уровень: Middle+ - Время на выполнение: 14-20 часов - Ключевые навыки: Ethereum RPC, background workers, subscription,..."
type: concept
---

# Ethereum Indexer

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/trust-wallet
- **Компания:** Trust Wallet
- **Уровень:** Middle+
- **Время на выполнение:** 14-20 часов
- **Ключевые навыки:** Ethereum RPC, background workers, subscription, indexing
- **Инфраструктура:** PostgreSQL (опционально), Docker

## ТЗ

### Описание
Создать сервис для индексации Ethereum блоков и транзакций. Поддержка подписки на адреса, фоновая синхронизация блоков, получение истории транзакций.

### Функциональные требования

**Эндпоинты:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| GET | /get_current_block | Текущий проиндексированный блок |
| GET | /subscribe | Подписаться на адрес |
| GET | /get_transactions | Получить транзакции по адресу |

**GET /get_current_block:**
```json
{
    "block_number": 18000000,
    "block_hash": "0x...",
    "timestamp": "2024-01-15T10:30:00Z"
}
```

**GET /subscribe?address=0x...:**
```json
{
    "address": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
    "status": "subscribed"
}
```

**GET /get_transactions?address=0x...:**
```json
{
    "address": "0x...",
    "transactions": [
        {
            "hash": "0x...",
            "block_number": 18000000,
            "from": "0x...",
            "to": "0x...",
            "value": "1000000000000000000",
            "gas_price": "20000000000",
            "timestamp": "2024-01-15T10:30:00Z"
        }
    ],
    "total": 150
}
```

**Бэкграунд индексатор:**

1. **Background Block Sync:**
   - Постоянная синхронизация новых блоков
   - Сохранение заголовков блоков
   - Обработка reorgs

2. **Background Transaction Sync:**
   - Для каждого подписанного адреса
   - Парсинг транзакций из блоков
   - Сохранение релевантных транзакций

### Нефункциональные требования

1. **RPC Calls:**
   - Pure RPC calls (без go-ethereum клиента для обучения)
   - Поддержка fallback между несколькими нодами
   - Обработка rate limiting

2. **Синхронизация:**
   - Background goroutines
   - Graceful shutdown
   - Retry с backoff

3. **Хранение:**
   - In-memory или PostgreSQL
   - Индексация по address

### Ограничения
- Нет websocket (HTTP polling)
- Нет ERC20 decoding (только ETH transfers)
- Нет real-time push

### Бонусные задания
- [ ] ERC20 token transfers decoding
- [ ] WebSocket subscriptions
- [ ] REST streaming (SSE)
- [ ] Multi-chain support
- [ ] GraphQL API

## Критерии оценки

- [ ] Background block sync работает
- [ ] Подписка на адреса
- [ ] Transaction parsing для subscribed адресов
- [ ] Get transactions endpoint
- [ ] Fallback между RPC нодами
- [ ] Retry и error handling
- [ ] Graceful shutdown
- [ ] Docker
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
    volumes:
      - postgres_data:/var/lib/postgresql/data

  indexer:
    build: .
    ports:
      - "8000:8000"
    environment:
      - DATABASE_URL=postgres://user:password@postgres:5432/indexer
      - ETH_RPC_URLS=https://rpc1.com,https://rpc2.com
    depends_on:
      - postgres

volumes:
  postgres_data:
```

### SQL-схема (опционально)
```sql
CREATE TABLE blocks (
    number BIGINT PRIMARY KEY,
    hash VARCHAR(66) UNIQUE NOT NULL,
    parent_hash VARCHAR(66) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    indexed_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE subscriptions (
    address VARCHAR(42) PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
    hash VARCHAR(66) PRIMARY KEY,
    block_number BIGINT REFERENCES blocks(number),
    "from" VARCHAR(42) NOT NULL,
    "to" VARCHAR(42),
    value NUMERIC NOT NULL,
    gas_price NUMERIC,
    nonce INTEGER,
    input TEXT,
    timestamp TIMESTAMP NOT NULL
);

CREATE INDEX idx_tx_from ON transactions("from");
CREATE INDEX idx_tx_to ON transactions("to");
CREATE INDEX idx_tx_block ON transactions(block_number DESC);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет background sync** — синхронный polling по запросу
2. **Нет fallback** — одна RPC нода, нет retry
3. **Нет graceful shutdown** — теряют данные при остановке
4. **Blocking channels** — небуферизованные каналы висят
5. **Нет reorg handling** — база расходится с реальностью

### Ключевые паттерны

**Background block sync:**
```go
func (i *Indexator) Start(ctx context.Context) {
    ticker := time.NewTicker(15 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            i.syncBlocks(ctx)
        case <-ctx.Done():
            return
        }
    }
}

func (i *Indexator) syncBlocks(ctx context.Context) {
    current := i.getCurrentBlock()
    latest, _ := i.ethClient.BlockNumber(ctx)
    
    for n := current + 1; n <= latest; n++ {
        block, _ := i.ethClient.BlockByNumber(ctx, big.NewInt(int64(n)))
        i.processBlock(block)
    }
}
```

**RPC с fallback:**
```go
func (c *EthClient) Call(method string, params []interface{}) (json.RawMessage, error) {
    for _, node := range c.nodes {
        result, err := c.callNode(node, method, params)
        if err == nil {
            return result, nil
        }
        log.Printf("Node %s failed: %v", node, err)
    }
    return nil, ErrAllNodesFailed
}
```

### Follow-up вопросы
1. Как масштабировать до 1000 подписок?
2. Как обработать реорганизацию цепочки?
3. Что если RPC нода отстаёт на 100 блоков?
4. Как реализовать real-time уведомления?

### Уровневые подсказки
- 🟢 **Направление:** Нужны две background горутины: одна для блоков, другая для транзакций подписанных адресов
- 🟡 **Конкретика:** Используй time.Ticker для периодического polling и канал для graceful shutdown
- 🔴 **Фрагмент:** ```go
func (i *Indexer) Start() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    go i.syncBlocks(ctx)
    go i.syncTransactions(ctx)
    
    <-i.shutdownChan
    cancel()
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L4 — Горутины, channels
- L5 — Context, graceful shutdown
- L6 — HTTP API
- L8 — PostgreSQL, Docker

### Новые навыки
- Ethereum RPC integration
- Background job patterns
- Blockchain indexing
- Subscription management

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

