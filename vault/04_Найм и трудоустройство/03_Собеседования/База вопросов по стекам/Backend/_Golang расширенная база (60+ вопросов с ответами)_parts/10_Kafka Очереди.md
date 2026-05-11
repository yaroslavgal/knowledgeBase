---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "⚡ 3 партиции, 5 консьюмеров → 2 простаивают. Параллелизм = `min(consumers, partitions)`."
---

## Kafka / Очереди

### 67. Kafka vs RabbitMQ `M`
- **Kafka** — распределённый лог, pull, replay, высокий throughput, партиции. Сообщения не удаляются после чтения. Для event streaming
- **RabbitMQ** — classic queue, push, удаление после ACK, гибкая маршрутизация (exchanges). Для task queues

### 68. Partition / Consumer Group / Offset `M`
- Partition — упорядоченный лог сообщений
- Consumer Group — несколько консьюмеров делят партиции
- Offset — позиция в партиции

⚡ 3 партиции, 5 консьюмеров → 2 простаивают. Параллелизм = `min(consumers, partitions)`.

### 69. ⚡ Ключ сообщения `M`
`hash(key) % numPartitions` → выбор партиции. Гарантирует: одинаковые ключи → одна партиция → упорядоченность для ключа. Без ключа — round-robin.

### 70. Гарантии доставки `M`
- **At-most-once** — commit ДО обработки, могут потеряться
- **At-least-once** — commit ПОСЛЕ обработки, могут задублироваться (норма в проде)
- **Exactly-once** — Kafka transactions + idempotent producer, дорого

В проде: at-least-once + idempotent consumer (дедуп по ключу).

### 71. Идемпотентный обработчик `M`
```go
// Дедуп по event_id
INSERT INTO processed_events (id) VALUES ($1) ON CONFLICT DO NOTHING;
// или upsert вместо insert
```

### 72. Consumer rebalancing `M`
При изменении числа consumers/partitions — Kafka переназначает. На время — пауза в обработке (stop-the-world группы). Минимизация: sticky assignor, правильный session.timeout, не деплоить всё одновременно, static membership.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
- [[02_IT-специальности и направления/01_Backend/_Что нужно знать для middle Go - чек-лист_parts/13_XIII. экосистема бэкенда|XIII. экосистема бэкенда]] — Kafka/RabbitMQ — обязательная часть экосистемы middle Go-бэкендера.
<!-- auto-links:end -->

