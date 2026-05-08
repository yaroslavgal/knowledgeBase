---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Структура мок-собеса Go (30 мин)|Структура мок-собеса Go (30 мин)]]"
type: concept
description: "1. «Что такое ACID? Пример нарушения каждой буквы.»"
---

## БЛОК 3: SQL / БД / Kafka (6 мин) — ОБЯЗАТЕЛЬНО

### Цепочка J — Транзакции и ACID

1. «Что такое ACID? Пример нарушения каждой буквы.»
2. «Уровни изоляции? Default в PG? В MySQL InnoDB?» (RC / RR)
3. «Dirty / non-repeatable / phantom — каждое одним сценарием.»
4. ⚡ «PG Repeatable Read защищает от phantom? А по стандарту SQL?»
5. ⚡ Long-lived tx с HTTP внутри:
```go
tx.Exec("UPDATE accounts SET balance=balance-100 WHERE id=1")
http.Post("https://payment.api/charge", ...)  // 3 секунды
tx.Exec("UPDATE accounts SET paid=true WHERE id=1")
tx.Commit()
```
   «Что не так? Что под нагрузкой?» (locks, queue, idle-in-tx, bloat)
6. «Outbox pattern. Зачем `FOR UPDATE SKIP LOCKED`?»
7. ⚡ «Как откатить tx при panic в Go?» (defer с recover)
8. «Как защититься от SQL injection? Почему `fmt.Sprintf` плохо?»

### Цепочка K — Индексы PostgreSQL

1. «B-tree под капотом?»
2. «Когда индекс НЕ помогает / замедляет?» (низкая кардинальность, LIKE '%x%', функция от столбца, type cast, >5% строк)
3. ⚡ «Seq scan быстрее Index scan когда >5% строк. Почему?» (random vs sequential I/O)
4. ```sql
   CREATE INDEX idx ON users(city, age);
   SELECT * FROM users WHERE city='Moscow' AND age>25;
   SELECT * FROM users WHERE age>25;
   SELECT * FROM users WHERE city='Moscow';
   ```
   «На каких сработает?» (1, 3 — да; 2 — нет, leftmost prefix)
5. «Правило выбора порядка полей в составном индексе?» (равенство первое, диапазон последнее)
6. «Index Only Scan — чем отличается? Как заставить?» (covering, VACUUM)
7. ⚡ «Индекс есть, EXPLAIN показывает Seq Scan. Шаги?» (ANALYZE, селективность, type cast, функции в WHERE)
8. «Запрос внезапно стал медленным, код не менялся. Дебажить — не сразу в индексы.» (EXPLAIN BUFFERS, статистика, pg_locks, bloat)

### Цепочка M — Kafka

1. «Kafka vs RabbitMQ. Для чего что?»
2. «Partition / Consumer Group / Offset. На пальцах.»
3. ⚡ «3 партиции, 5 consumers в группе. Что с 2 лишними?» (idle)
4. «Ключ сообщения — для чего? Без ключа?»
5. «at-most/at-least/exactly-once. Что в проде?» (at-least + idempotent consumer)
6. ⚡ «order.created прилетел дважды. Сделай idempotent.» (dedup по event_id, INSERT ON CONFLICT DO NOTHING)
7. «Consumer rebalancing — почему боль, как минимизировать?» (sticky assignor, static membership)
8. «Consumer lag растёт. Как разобраться?»

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Структура мок-собеса Go (30 мин)|Структура мок-собеса Go (30 мин)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

