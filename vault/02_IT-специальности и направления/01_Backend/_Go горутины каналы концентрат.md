# Go: горутины и каналы — концентрат для собесов

G (Goroutine) → M (Machine = OS thread) → P (Processor = logical CPU) GOMAXPROCS = число P (по умолчанию = число CPU cores)

## Главные концепции

### Scheduler модель G-M-P

```
G (Goroutine) → M (Machine = OS thread) → P (Processor = logical CPU)
GOMAXPROCS = число P (по умолчанию = число CPU cores)
```

- Горутина ≈ **2KB** стека (растёт динамически)
- Не привязана к OS-потоку (М может переключаться между G)
- P содержит локальную очередь горутин + глобальная очередь
- Когда M блокируется (syscall) — другой M подхватывает P

### Каналы под капотом

```go
type hchan struct {
    qcount   uint           // элементов в буфере
    dataqsiz uint           // размер кольцевого буфера
    buf      unsafe.Pointer // кольцевой буфер
    sendx    uint           // позиция записи
    recvx    uint           // позиция чтения
    sendq    waitq          // ожидающие на send
    recvq    waitq          // ожидающие на recv
    lock     mutex          // под капотом — futex
}
```

**Главное:** каналы реализованы через **mutex + кольцевой буфер + списки ожидающих**.

---

## Главные ошибки на собесах (реальные кейсы)

### Кейс 1: deadlock при close после WaitGroup

```go
ch := make(chan string, 3)
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        ch <- "msg"
    }()
}

// ❌ ОШИБКА: close(ch) тут = panic на send в closed channel
close(ch)
wg.Wait()
```

**✅ Исправление:**
```go
go func() { wg.Wait(); close(ch) }()
for msg := range ch { fmt.Println(msg) }
```

### Кейс 2: бесконечный select без выхода

```go
// ❌ После close ch будет вечно печатать пустые строки
for {
    select {
    case v := <-ch:
        fmt.Println(v)
    }
}
```

**✅ Исправление через `for range`:**
```go
for v := range ch { fmt.Println(v) }
// Выходит автоматически на close
```

### Кейс 3: длина буфера < числа продьюсеров

```go
// Буфер 3, продьюсеров 5 → deadlock пока никто не читает
ch := make(chan int, 3)
for i := 0; i < 5; i++ {
    go func() { ch <- i }()
}
// Зависнет на 4-м продьюсере
```

**✅ Решение:** запустить close в отдельной горутине + `for range ch`.

---

## Закрытие каналов — правила

| Действие | Результат |
|---|---|
| `close(ch)` где ch == nil | **panic** |
| `close(ch)` где ch уже закрыт | **panic** |
| `ch <- val` где ch закрыт | **panic** |
| `<-ch` где ch закрыт | **zero value + ok=false** (нормально) |
| `select{}` без кейсов | блокировка навсегда |

**Кто закрывает канал:** только **producer** (тот кто пишет). Никогда consumer.

---

## sync пакет

### Mutex / RWMutex

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

// ❗ ВАЖНО: pointer receiver обязателен — иначе copy of mutex
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```

**RWMutex когда:** read-heavy (90%+ чтений). Иначе обычный Mutex быстрее.

### WaitGroup

```go
var wg sync.WaitGroup
wg.Add(n)  // ❗ ВСЕГДА до запуска горутин
for i := 0; i < n; i++ {
    go func() {
        defer wg.Done()
        // работа
    }()
}
wg.Wait()
```

### Once

```go
var once sync.Once
func init() {
    once.Do(func() {
        // выполнится ровно 1 раз даже если 10 горутин зайдут
    })
}
```

### sync.Map

- Хорош для **read-heavy + редкие записи**
- Для write-heavy → обычный `map + RWMutex` быстрее

### sync.Pool

- Переиспользование объектов (буферы, парсеры)
- **Не гарантирует** сохранение объекта между вызовами
- Снимает давление с GC

---

## Memory Model: happens-before гарантии

Гарантии порядка операций:

1. **Channel send** happens-before **receive** этой же переменной
2. **Mutex Unlock** happens-before следующего **Lock**
3. **`once.Do(f)`** happens-before любое другое **Do**
4. **Init** happens-before **main**
5. **Goroutine creation** happens-before её **выполнения**

Из этих правил выводится: **видимость записей** в памяти.

---

## Типичные западни

### `time.After` в горячем цикле

```go
// ❌ Создаёт горутину и канал на каждой итерации — memory leak
for {
    select {
    case <-ch:
        // ...
    case <-time.After(time.Second):
        // ...
    }
}
```

**✅ Используй `time.NewTimer` с reset:**
```go
timer := time.NewTimer(time.Second)
defer timer.Stop()
for {
    select {
    case <-ch:
        // ...
    case <-timer.C:
        timer.Reset(time.Second)
    }
}
```

### `nil` interface vs interface с nil-значением

```go
var p *Person = nil
var i interface{} = p
fmt.Println(i == nil) // false! ⚠️
```

Interface = (type, value). Если type есть, а value = nil — interface не равен nil.

### Race condition на чтение слайса

```go
// ❌ Race на чтение len() если другая горутина делает append
type Stack struct {
    mu   sync.Mutex
    data []int
}

func (s *Stack) Len() int { return len(s.data) }  // RACE
```

**✅ Защити чтение тем же мьютексом:**
```go
func (s *Stack) Len() int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return len(s.data)
}
```

---

## Race Detector

```bash
go test -race ./...
go run -race main.go
go build -race
```

**Включает overhead в ~10x** — только для тестов и dev. Не для прода.

---

## Context.Context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Передача через все функции
result, err := doWork(ctx, ...)

// Внутри функции
func doWork(ctx context.Context, ...) {
    select {
    case <-ctx.Done():
        return ctx.Err()  // context.Canceled или context.DeadlineExceeded
    case result := <-doActualWork():
        return result
    }
}
```

**Правила:**
- ВСЕГДА первый аргумент функции
- ВСЕГДА `defer cancel()` после `WithTimeout` / `WithCancel`
- Не клади в struct, передавай через параметры

---

## Топ-вопросы по Go в 2026

1. **GOMAXPROCS vs runtime.NumCPU** — в чём разница, когда менять
2. **Generics (с 1.18)** — синтаксис, ограничения, type parameters
3. **Escape analysis** — когда переменная попадает на heap vs stack
4. **GC: tri-color marking** — как работает, когда STW
5. **Defer / panic / recover** — порядок выполнения, edge cases
6. **slice внутреннее устройство** — pointer + len + cap
7. **map внутреннее устройство** — buckets, growing, evacuation
8. **strings vs runes** — байты vs Unicode, range особенности
9. **interface internals** — itab, методы
10. **net/http internals** — middleware, server, client

---

## Чек-лист готовности к Go-собесу

- [ ] Понимаешь G-M-P scheduler
- [ ] Знаешь когда `panic` от каналов (close, send)
- [ ] `sync.Mutex` нельзя копировать → pointer receiver
- [ ] `WaitGroup.Add` ВСЕГДА до запуска горутин
- [ ] `Race detector` запускал и понимаешь
- [ ] Понимаешь happens-before гарантии
- [ ] Знаешь `time.After` western leak
- [ ] Понимаешь `nil interface` подвох
- [ ] Знаешь `context.Context` правила
- [ ] Можешь объяснить escape analysis
- [ ] Понимаешь как работает map (growing, hash collisions)
- [ ] Mock-собес с лайвкодингом пройден 5+ раз

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

