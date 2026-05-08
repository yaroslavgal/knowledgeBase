---
parent: "[[40_Найм и трудоустройство]]"
description: "Не «прочитай книгу». 20 задач в правильном порядке = путь от Hello World до concurrency-задач BigTech."
type: concept
---

# От нуля до live coding senior — 20 задач (Go)

Не «прочитай книгу». **20 задач в правильном порядке = путь от Hello World до concurrency-задач BigTech.**

Каждая задача — реальный навык, который спрашивают на собесах в Ozon/Авито/Lamoda/VK. Прогресс измеряется задачами, не часами.

---

## Уровень 1 — Основы (1-5)

### 1. Hello, Go с CLI-аргументом

Принимает имя из `os.Args[1]`. Без аргумента — "Hello, World!".

Ловушка: `os.Args[1]` без проверки длины → panic.
```go
name := "World"
if len(os.Args) > 1 { name = os.Args[1] }
```

**Концепции:** package main, os.Args, fmt.Printf, проверка границ.

### 2. Калькулятор с обработкой ошибок

`divide(a, b float64) (float64, error)` — деление на ноль возвращает error.

```go
if b == 0 { return 0, errors.New("division by zero") }
return a / b, nil
```

**Концепции:** паттерн (value, error), множественный возврат, errors.New.

### 3. Фильтрация слайса

`filter(nums []int, pred func(int) bool) []int`.

Преаллокация обязательна:
```go
result := make([]int, 0, len(nums))
```

**Концепции:** slice cap, make, функции высшего порядка.

### 4. Подсчёт частоты слов

`wordCount(text string) map[string]int`.

```go
counts := make(map[string]int)
for _, w := range strings.Fields(text) {
    counts[strings.ToLower(w)]++
}
```

**Концепции:** map, идиома `m[k]++`, рандомный обход.

### 5. Интерфейс Shape

`Circle{Radius}`, `Rectangle{W,H}` реализуют `Shape{Area() float64}`.

**Концепции:** implicit реализация (duck typing), value receiver.

---

## Уровень 2 — Углублённый Go (6-10)

### 6. Кастомный тип ошибки

`ValidationError{Field, Message}`. Использовать `errors.Is` и `errors.As`.

```go
type ValidationError struct{ Field, Message string }
func (e *ValidationError) Error() string { return e.Field + ": " + e.Message }

var ve *ValidationError
if errors.As(err, &ve) { /* доступ к полям */ }
```

**Концепции:** error wrapping (`%w`), errors.Is vs As.

### 7. Генератор Фибоначчи через замыкание

`fibonacci() func() int` — каждый вызов даёт следующее число.

```go
func fibonacci() func() int {
    a, b := 0, 1
    return func() int { a, b = b, a+b; return a }
}
```

**Концепции:** замыкания, lifetime переменных.

### 8. Реверс Unicode-строки

```go
runes := []rune(s) // НЕ []byte — UTF-8 сломается
for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
}
return string(runes)
```

**Концепции:** byte vs rune, UTF-8.

### 9. Safe execute с panic/recover

`safeExecute(fn func()) error` — оборачивает и ловит panic, возвращает error.

```go
defer func() {
    if r := recover(); r != nil {
        err = fmt.Errorf("panic: %v", r)
    }
}()
```

**Концепции:** defer/panic/recover, порядок defer.

### 10. Counter с pointer receiver

`Increment()` через `*Counter`, `Value()` через `Counter`. Объясни разницу.

Если `Increment` value receiver — будет работать с копией, изменения не сохранятся. Pointer receiver обязателен для модификации.

**Концепции:** pointer vs value receiver, method set.

---

## Уровень 3 — Concurrency Junior (11-14)

### 11. Параллельный вывод 5 чисел

```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {           // ОБЯЗАТЕЛЬНО передать как аргумент
        defer wg.Done()
        fmt.Println(n)
    }(i)
}
wg.Wait()
```

Без передачи `i` как аргумента — на Go до 1.22 выведет 5 5 5 5 5.

**Концепции:** WaitGroup, замыкание на переменную цикла.

### 12. Генератор чисел через канал

Горутина пишет 1-10 в канал и закрывает. Main читает через `range`.

```go
ch := make(chan int)
go func() {
    defer close(ch)
    for i := 1; i <= 10; i++ { ch <- i }
}()
for v := range ch { fmt.Println(v) }
```

**Концепции:** channels, close, range по каналу.

### 13. Worker Pool (fan-out)

3 воркера читают из общего канала задач, пишут результаты в выходной.

```go
jobs := make(chan int, 10)
results := make(chan int, 10)
var wg sync.WaitGroup
for w := 0; w < 3; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs { results <- j*j }
    }()
}
for i := 1; i <= 10; i++ { jobs <- i }
close(jobs)
go func() { wg.Wait(); close(results) }()
for r := range results { fmt.Println(r) }
```

Это **самый частый паттерн на собесах**. Уметь написать на доске за 5 минут без подсказок.

**Концепции:** буферизированные каналы, координация закрытия.

### 14. Fetch с timeout через context

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
select {
case res := <-doRequest(ctx):
    return res
case <-ctx.Done():
    return ctx.Err()
}
```

**Концепции:** context.WithTimeout, обязательный defer cancel.

---

## Уровень 4 — Concurrency Middle (15-17)

### 15. Потокобезопасный кэш с RWMutex

```go
type Cache struct {
    mu   sync.RWMutex            // НЕ embedding — приватное поле
    data map[string]string
}

func (c *Cache) Get(k string) (string, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    v, ok := c.data[k]
    return v, ok
}

func (c *Cache) Set(k, v string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[k] = v
}
```

Ловушки:
- Embedding `sync.RWMutex` — Lock/Unlock уходят в публичный API, ломая инкапсуляцию
- Lock в Get вместо RLock — сериализует читатели, теряем смысл RWMutex
- Возврат внутреннего слайса/мапы наружу — пользователь модифицирует без блокировки

**Концепции:** RWMutex, инкапсуляция, mutex copying.

### 16. Merge каналов (fan-in)

`merge(channels ...<-chan int) <-chan int` — объединяет N в один.

```go
func merge(chans ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, c := range chans {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c { out <- v }
        }(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```

**Концепции:** select, координация закрытия выходного канала.

### 17. Pipeline (конвейер)

`generate → square → filter → print`. Каждая стадия — горутина с каналом.

```go
func generate(done <-chan struct{}) <-chan int { /* ... */ }
func square(done <-chan struct{}, in <-chan int) <-chan int { /* ... */ }
```

**Концепции:** chaining каналов, done-channel для отмены.

---

## Уровень 5 — Concurrency Senior (18-20)

### 18. Rate Limiter (Token Bucket)

`NewLimiter(rate, burst)` с методами `Allow() bool` и `Wait(ctx) error`.

Через `time.Ticker` и канал-буфер как «корзина токенов».

**Концепции:** time.Ticker, канал как семафор, graceful shutdown через context.

### 19. Singleflight — дедупликация запросов

Несколько горутин просят одно значение → реальный запрос делает только первая, остальные ждут её результат.

```go
type Group struct {
    mu    sync.Mutex
    calls map[string]*call
}
type call struct {
    wg  sync.WaitGroup
    val interface{}
    err error
}
```

**Концепции:** sync.WaitGroup для ожидания, дедуп.

### 20. Graceful Shutdown HTTP-сервера

SIGTERM → перестать принимать новые → дождаться текущих → закрыть БД.

```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
<-sigChan

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := server.Shutdown(ctx); err != nil {
    log.Fatal("Forced shutdown:", err)
}
db.Close()
```

Без graceful shutdown в k8s = текущие запросы обрываются, transactions откатываются, клиенты получают connection reset.

**Концепции:** os.Signal, signal.Notify, context.WithTimeout, http.Server.Shutdown.

---

## Чек-лист

| # | Задача | Уровень |
|---|--------|---------|
| 1-5 | Hello / Калькулятор / Фильтр / Words / Shapes | Новичок |
| 6-10 | Errors / Closures / Reverse / Recover / Receivers | Продвинутый |
| 11-14 | WG / Chan / Worker Pool / Context | Junior conc. |
| 15-17 | Кеш RWMutex / Merge / Pipeline | Middle conc. |
| 18-20 | Rate Limiter / Singleflight / Graceful Shutdown | Senior conc. |

---

## Дополнительные ловушки

- **`-race`** ловит data race, не race condition. Запускай в тестах: `go test -race ./...`
- **Nil-каналы в select** — присвоение nil блокирует case. Полезно для merge/fan-in
- **Копирование sync примитивов** — `go vet` ловит copylocks
- **Не отдавай защищённые данные наружу** — если `Data()` возвращает внутренний слайс, синхронизация бесполезна
- **False sharing** (senior) — горутины пишут в соседние участки памяти, cache line invalidation бьёт по производительности

---

## Как тренироваться

Каждую задачу решаешь **сам**, без подсказок. Минимум 30 минут на попытку. Только потом смотришь эталон и сравниваешь.

На concurrency-задачах **обязательно** запускай с `-race`. Если стрельнул — фикси, пока не зелёный.

Цель — не «знать решение», а **писать его за 5-7 минут на доске под давлением**, проговаривая вслух.

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Live coding Go реальные задачи|Live coding Go реальные задачи]] — финальный уровень, к которому ведут эти 20 задач.
- [[02_IT-специальности и направления/01_Backend/_Горутины и каналы за час - концентрат|Горутины и каналы за час]] — теоретический концентрат для concurrency-задач (последний этап).
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Подготовка к техсобесу|Подготовка к техсобесу]] — общая стратегия подготовки, в которой эти 20 задач — практический трек.
- [[01_Образование и обучение/05_Путь обучения по специальностям/_Roadmap входа в Go|Roadmap входа в Go]] — общий roadmap, в котором эти задачи — финальная фаза.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Структура мок-собеса Go (30 мин)|Структура мок-собеса Go]] — формат, в котором проверяется готовность по этим 20 задачам.
<!-- auto-links:end -->

