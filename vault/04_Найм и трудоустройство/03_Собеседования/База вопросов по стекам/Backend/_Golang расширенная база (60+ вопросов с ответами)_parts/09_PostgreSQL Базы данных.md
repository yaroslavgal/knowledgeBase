---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "Read Uncommitted → Read Committed → Repeatable Read → Serializable."
---

## PostgreSQL / Базы данных

### 55. ACID `M`
- **A** — атомарность (всё или ничего)
- **C** — консистентность (CHECK, FK)
- **I** — изоляция (транзакции не мешают)
- **D** — durability (после commit — не пропадёт)

### 56. Уровни изоляции `M`
Read Uncommitted → Read Committed → Repeatable Read → Serializable.
- PostgreSQL default = Read Committed
- MySQL InnoDB default = Repeatable Read

⚡ В PG Repeatable Read защищает от phantom read через MVCC (snapshot isolation), хотя в стандарте SQL — нет.

### 57. ⚡ HTTP внутри транзакции `M`
```go
tx, _ := db.Begin()
tx.Exec("UPDATE accounts SET balance = balance - 100 WHERE id = 1")
http.Post("https://payment.api/charge", ...)  // 3 сек!
tx.Exec("UPDATE accounts SET paid = true WHERE id = 1")
tx.Commit()
```
Long-lived tx, блокировки строк на 3 сек, очереди, idle-in-transaction, bloat. Фикс: **Outbox pattern**.

### 58. Index работает или нет `M`
```sql
CREATE INDEX idx ON users(city, age);
WHERE city='Moscow' AND age>25  -- да
WHERE age>25                    -- нет (нет leftmost prefix)
WHERE city='Moscow'             -- да
WHERE age>25 AND city='Moscow'  -- да (порядок AND не важен)
```

Правило составного индекса: поля с `=` первыми, поля с `<`/`>`/`BETWEEN` последними, высокая кардинальность первой.

### 59. ⚡ Seq scan быстрее Index scan когда `M`
Когда возвращается > ~5-10% таблицы. Index scan делает random I/O (прыжки по heap), Seq scan — последовательное чтение, cache-friendly.

### 60. Index Only Scan `M`
Не ходит в heap, всё из индекса. Включается через `INCLUDE (other_cols)` в индексе (covering) + актуальная visibility map (нужен VACUUM).

### 61. EXPLAIN — что смотреть `M`
- `Seq Scan` vs `Index Scan` vs `Bitmap Index Scan`
- `cost` — оценка планировщика
- `rows` — оценка vs реальное (если разница большая — статистика устарела, нужен `ANALYZE`)
- `buffers` — операции с диском/кэшем

⚡ Запрос внезапно стал медленным: первое — `EXPLAIN (ANALYZE, BUFFERS)`, не лезть в индексы. Возможно: статистика устарела, конкурентные блокировки, table bloat без autovacuum.

### 62. MVCC `M`
Каждая строка в нескольких версиях (xmin, xmax). Читатели видят snapshot на момент tx. Не блокируют писателей и наоборот. UPDATE/DELETE не удаляет — помечает как dead tuple. VACUUM чистит.

### 63. SELECT FOR UPDATE SKIP LOCKED `M`
FOR UPDATE — эксклюзивная блокировка строк до конца tx. SKIP LOCKED — пропускает уже заблокированные. Применение: очереди задач в БД, несколько воркеров параллельно.

### 64. ⚡ idle in transaction `M`
Открытая tx, ничего не делает (например код ждёт HTTP-ответ). Держит снапшот, мешает VACUUM, накапливает bloat. Ставить `idle_in_transaction_session_timeout` в конфиге.

### 65. SQL injection защита `M`
Только параметризованные запросы (`$1`, `?`), драйвер сам экранирует. `fmt.Sprintf("...WHERE name='%s'", input)` — мгновенный красный флаг на собесе.

### 66. database/sql пул соединений `M`
- `SetMaxOpenConns` — обычно 20-50 на инстанс (не 1000!)
- `SetMaxIdleConns` — половина от max
- `SetConnMaxLifetime` — 5-30 мин

Для 1000 RPS — критично pool tuning + pgbouncer.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

