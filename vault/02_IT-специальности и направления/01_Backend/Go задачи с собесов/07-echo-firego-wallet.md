---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/echo - Компания: Echo - Уровень: Middle+ - Время на выполнение: 14-20 часов - Ключевые навыки: Fireblocks API, крипто-кошельки, PostgreSQL, интеграция - Инфра..."
type: concept
---

# FireGo Wallet Service

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/echo
- **Компания:** Echo
- **Уровень:** Middle+
- **Время на выполнение:** 14-20 часов
- **Ключевые навыки:** Fireblocks API, крипто-кошельки, PostgreSQL, интеграция
- **Инфраструктура:** PostgreSQL, Docker, Fireblocks API (testnet)

## ТЗ

### Описание
Создать backend-сервис для работы с криптовалютными кошельками через Fireblocks API. Сервис управляет Vault Accounts, получает балансы, адреса для депозита и иницииирует переводы.

### Функциональные требования

**Эндпоинты:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| POST | /wallets | Создать новый кошелёк (Vault Account) |
| GET | /wallets/{walletId}/assets/{assetId}/balance | Получить баланс |
| GET | /wallets/{walletId}/assets/{assetId}/address | Получить deposit адрес |
| POST | /wallets/{walletId}/transactions | Инициировать перевод |

**Создание кошелька:**
```json
POST /wallets
{
    "name": "My Personal Wallet"
}

Response:
{
    "id": "wallet-uuid",
    "name": "My Personal Wallet",
    "vault_account_id": "123",
    "created_at": "2024-01-15T10:30:00Z"
}
```

**Получение баланса:**
```
GET /wallets/{walletId}/assets/TEST_BTC/balance

Response:
{
    "asset_id": "TEST_BTC",
    "balance": "0.5",
    "available": "0.5",
    "pending": "0"
}
```

**Инициирование перевода:**
```json
POST /wallets/{walletId}/transactions
{
    "assetId": "TEST_BTC",
    "amount": "0.001",
    "destinationAddress": "tb1q...",
    "note": "Test transfer"
}

Response:
{
    "transaction_id": "tx-uuid",
    "status": "pending",
    "created_at": "2024-01-15T10:30:00Z"
}
```

**Поддерживаемые assets:**
- `TEST_BTC` — Bitcoin testnet
- `TEST_ETH_TEST3` — Goerli ETH

### Нефункциональные требования

1. **Интеграция:** Fireblocks API (sandbox/testnet)
2. **База данных:** PostgreSQL для хранения wallet mappings
3. **Архитектура:** Clean Architecture
4. **Валидация:** Проверка входных данных
5. **Обработка ошибок:** Корректные HTTP статусы и сообщения

### Ограничения
- Только testnet (нет real крипты)
- API ключи через environment variables
- Нет webhook обработки (только polling опционально)

### Бонусные задания
- [ ] Polling статуса транзакции
- [ ] Webhook обработка callback'ов от Fireblocks
- [ ] История транзакций
- [ ] Поддержка дополнительных assets

## Критерии оценки

- [ ] Корректная интеграция с Fireblocks API
- [ ] Создание Vault Accounts
- [ ] Получение балансов и адресов
- [ ] Инициирование переводов
- [ ] Сохранение wallet mappings в PostgreSQL
- [ ] Обработка ошибок
- [ ] Валидация входных данных
- [ ] Docker и docker-compose
- [ ] README с инструкциями по запуску

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: firego
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"

  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      - DATABASE_URL=postgres://user:password@postgres:5432/firego
      - FIREBLOCKS_API_KEY=${FIREBLOCKS_API_KEY}
      - FIREBLOCKS_SECRET_KEY_PATH=/secrets/fireblocks_secret.key
      - FIREBLOCKS_BASE_URL=https://sandbox-api.fireblocks.io
```

### SQL-схема
```sql
CREATE TABLE wallets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    vault_account_id VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    wallet_id UUID REFERENCES wallets(id),
    fireblocks_tx_id VARCHAR(255) NOT NULL,
    asset_id VARCHAR(50) NOT NULL,
    amount VARCHAR(50) NOT NULL,
    destination_address VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет маппинга** — не сохраняют связь local ID ↔ vault account ID
2. **Секреты в коде** — API ключи захардкожены
3. **Нет валидации** — принимают любые asset_id
4. **Путают testnet/mainnet** — пытаются работать с реальными деньгами
5. **Нет idempotency** — дублирование создания кошельков

### Ключевые паттерны
- **API Client wrapper** — абстракция над Fireblocks SDK
- **Repository pattern** — работа с БД
- **Idempotency key** — предотвращение дублей
- **Circuit breaker** — для внешнего API

### Fireblocks интеграция
```go
type FireblocksClient struct {
    client *fireblocks.Client
}

func (c *FireblocksClient) CreateVaultAccount(ctx context.Context, name string) (string, error) {
    resp, err := c.client.CreateVaultAccount(ctx, &fireblocks.CreateVaultAccountRequest{
        Name: name,
    })
    return resp.ID, err
}
```

### Follow-up вопросы
1. Как бы ты добавил поддержку hardware wallet?
2. Что делать если Fireblocks API недоступен?
3. Как реализовать multi-signature?
4. Как защитить от double-spending?

### Уровневые подсказки
- 🟢 **Направление:** Нужен клиент для Fireblocks API + сервис для управления кошельками
- 🟡 **Конкретика:** Используй Fireblocks Go SDK, храни API credentials в env
- 🔴 **Фрагмент:** ```go
func (s *WalletService) CreateWallet(ctx context.Context, name string) (*Wallet, error) {
    // Создаём в Fireblocks
    vaultID, err := s.fireblocks.CreateVaultAccount(ctx, name)
    if err != nil { return nil, err }
    
    // Сохраняем маппинг
    wallet := &Wallet{Name: name, VaultAccountID: vaultID}
    return s.repo.Create(ctx, wallet)
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L6 — HTTP, middleware, JSON
- L7 — Clean architecture
- L8 — PostgreSQL, Docker, external APIs

### Новые навыки
- Интеграция с Fireblocks
- Работа с крипто-API
- Vault account management
- Financial API security

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

