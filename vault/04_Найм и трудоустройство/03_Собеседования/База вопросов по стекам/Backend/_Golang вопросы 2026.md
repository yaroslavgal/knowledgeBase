---
parent: "[[40_Найм и трудоустройство]]"
description: "Статистика по 70+ мок-собесам (Ozon, Lamoda, Avito, Тинькофф, Яндекс) и записям реальных интервью."
type: concept
---

# Golang вопросы 2026 — что реально спрашивают

Статистика по 70+ мок-собесам (Ozon, Lamoda, Avito, Тинькофф, Яндекс) и записям реальных интервью.

В РФ копают ВНУТРЯНКУ рантайма (планировщик, эвакуация мап, GMP). За рубежом — больше прикладные навыки и стек.

---

## Топ-15 по вероятности

| # | Тема | % | Грейд |
|---|------|---|-------|
| 1 | Слайсы — устройство, append, cap | 95% | Все |
| 2 | Мапы — бакеты, эвакуация, load factor 6.5 | 90% | Все |
| 3 | Горутины + GMP-планировщик | 90% | Все |
| 4 | Каналы — буф/небуф, select, deadlock | 85% | Все |
| 5 | PostgreSQL — индексы, EXPLAIN, изоляция | 85% | Все |
| 6 | Интерфейсы — устройство, nil-ловушка | 80% | Все |
| 7 | Лайвкодинг (worker pool, кеш, defer) | 80% | Все |
| 8 | sync — Mutex, RWMutex, WaitGroup | 75% | Все |
| 9 | Строки — иммутабельность, byte vs rune | 70% | Все |
| 10 | Kafka — гарантии, consumer groups | 60-70% | Middle+ |
| 11 | Паттерны concurrency — fan-in/out, errgroup | 60% | Middle+ |
| 12 | Микросервисы — Saga, circuit breaker | 50-70% | Middle+ |
| 13 | gRPC vs REST | 50% | Middle+ |
| 14 | Алгоритмы (LeetCode) | 40-60% | Зависит от компании |
| 15 | System Design | 30% мидл / 70% сеньор | Senior |

---

## Слайсы (95% — спросят 100%)

«Как устроен слайс? Чем отличается от массива?»

Слайс = `{pointer, len, cap}` — 24 байта на 64-bit. View на underlying array.

```go
// Классическая ловушка
func modify(s []int) {
    s = append(s, 4) // caller НЕ увидит — append вернул новый header
}
// Правильно: возвращаем
func modify(s []int) []int {
    return append(s, 4)
}
```

Cap растёт ×2 до ~1024, потом ~×1.25.

Trap: `s[1:3]` → `sub`, `append(sub, 99)` затрагивает исходный underlying array. Защита — full slice expression `s[1:3:3]`.

---

## Мапы (90%)

«Как устроена map?»

- Бакеты по 8 пар key/value, overflow-бакеты по цепочке
- Load factor 6.5 — триггер эвакуации
- Эвакуация **инкрементальная** — размазана по операциям, не сразу
- Порядок обхода **намеренно рандомизирован** — чтобы код не опирался
- Нельзя получить указатель на значение в мапе
- Запись в `nil` map → panic, чтение → zero value

Trap: «Вставка в map — потокобезопасна?» Нет. Конкурентная запись = fatal error (не panic — fatal). `sync.Map` или `sync.RWMutex + map`.

---

## Горутины + GMP (90%)

GMP-модель:
- **G** — горутина (~2-8 KB стек, contiguous, растёт удвоением)
- **M** — поток ОС (kernel thread)
- **P** — логический процессор с локальной очередью (= GOMAXPROCS, обычно по числу ядер)

Work stealing — пустой P ворует половину очереди соседа.

Sysmon — отдельный поток, мониторит каждые ~10 мс: вытесняет долгие горутины (с Go 1.14 — async preemption через сигналы), отбирает P у заблокированных на syscall M.

Trap: «Сколько потоков создаст программа?» Минимум GOMAXPROCS + дополнительные при блокирующих syscall. Лимит M по умолчанию 10 000.

---

## Каналы (85%)

| Действие | Nil chan | Закрытый |
|----------|---------|----------|
| Чтение | блок навсегда | zero value, ok=false |
| Запись | блок навсегда | panic |
| Закрыть | panic | panic |

Nil-канал в `select` — способ «выключить» ветку.

Trap-задача (утечка):
```go
func slow() string {
    ch := make(chan string)              // НЕ буферизированный
    go func() { ch <- fetchSlow() }()
    select {
    case s := <-ch: return s
    case <-time.After(1*time.Second): return "timeout"
    }
    // Если timeout сработал — горутина повисла навсегда на ch <- ...
}
// Фикс: make(chan string, 1)
```

---

## Интерфейсы (80%)

Под капотом: `(type, value)`. Implicit реализация (structural typing) — никаких `implements`.

**Главная ловушка** — nil-интерфейс:
```go
type MyError struct{}
func (*MyError) Error() string { return "bad" }

func doIt() error {
    var err *MyError = nil
    return err
}

if err := doIt(); err != nil {
    fmt.Println("WTF?") // выведется! type=*MyError, value=nil → интерфейс != nil
}
```

Фикс: `return nil` (литерал nil интерфейса), а не nil-указатель.

Принцип: «Accept interfaces, return structs» — потребитель описывает минимальный контракт.

Compile-time check: `var _ MyInterface = (*MyType)(nil)`.

---

## sync (75%)

| Примитив | Когда |
|----------|-------|
| Mutex | Любая защита состояния |
| RWMutex | Read-heavy (90%+ чтения), долгие чтения |
| WaitGroup | Дождаться N горутин |
| Once | Ленивая инициализация ровно один раз |
| atomic | Одна переменная, простая операция |
| Pool | Переиспользование аллокаций (буферы) |

Trap: встраивание `sync.Mutex` анонимным полем — Lock/Unlock уходят в публичный API. Делай явное поле `mu sync.Mutex`.

Trap: копирование структуры с Mutex по значению — `go vet` ловит copylocks.

---

## PostgreSQL (85%)

**ACID** — каждую букву с примером. PG default isolation = Read Committed. MySQL InnoDB = Repeatable Read.

**Уровни изоляции и аномалии:**

| Уровень | Dirty | Non-repeatable | Phantom |
|---------|-------|----------------|---------|
| Read Uncommitted | + | + | + |
| Read Committed | — | + | + |
| Repeatable Read | — | — | +/— |
| Serializable | — | — | — |

В PostgreSQL Repeatable Read защищает от phantom (благодаря MVCC snapshot). В стандарте SQL — нет.

**Индексы:**
- B-tree (default), Hash, GIN (полнотекст, JSONB), GiST (гео)
- Составной индекс — leftmost prefix rule
- Не работает: `WHERE LOWER(col)=...` без functional index, `LIKE '%x%'`, низкая кардинальность, приведение типов
- Index Only Scan — данные берутся из индекса, нужен covering + актуальный VACUUM

**EXPLAIN ANALYZE** — обязательно уметь читать. Seq Scan быстрее Index Scan если возвращается >5% строк (random I/O vs sequential).

---

## Kafka (60-70%)

- Partition = упорядоченный лог. Внутри партиции — порядок гарантирован, между — нет.
- Consumer Group — каждая партиция читается одним consumer'ом группы. Больше consumer'ов чем партиций = idle.
- Hash(key) % partitions → выбор партиции. Один key всегда в одной партиции.

Гарантии:
- at-most-once — commit до обработки, потеря возможна
- at-least-once — commit после, дубли возможны (стандарт прода)
- exactly-once — Kafka transactions + idempotent producer, дорого

Прод: at-least-once + idempotent consumer (dedup по event_id).

---

## Микросервисы (50-70%)

- **Saga** — распределённый процесс с компенсациями
- **Outbox pattern** — событие в одну tx с основными данными, воркер публикует. `FOR UPDATE SKIP LOCKED` для параллельных воркеров
- **Circuit breaker, retry с exponential backoff, dead letter queue**
- **Graceful shutdown** — обязательно. SIGTERM → перестать принимать → дождаться текущих → закрыть БД

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
<-sigChan

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
server.Shutdown(ctx)
```

---

## Что копают на сеньора отдельно

- Defer + named returns — defer может изменить возвращаемое значение
- errors.Is vs errors.As, оборачивание `%w`
- Context — почему `defer cancel()` обязателен, что класть в WithValue
- sync.Map — когда оправдан (append-only, read-heavy с редкими записями)
- pprof — goroutine/heap/cpu/mutex/block профили
- MVCC — dead tuples, VACUUM vs VACUUM FULL, idle in transaction

См. `_Worker_pool_и_лайвкод_задачи.md` для практики.

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
<!-- auto-links:end -->

