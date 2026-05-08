---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "Закрывать должен только **отправитель**, и только когда никто больше не пишет."
---

## Channels

### 12. Запись/чтение в закрытый канал `M`

| Операция | Открытый | Закрытый | nil |
|----------|---------|---------|-----|
| Отправка | блок/запись | **паника** | вечная блокировка |
| Чтение | блок/чтение | zero value, ok=false | вечная блокировка |
| Закрытие | закрыл | **паника** | **паника** |

Закрывать должен только **отправитель**, и только когда никто больше не пишет.

### 13. Buffered vs unbuffered `M`
- Unbuffered (`make(chan T)`) — отправитель блокируется до читателя. Точка синхронизации, рукопожатие.
- Buffered (`make(chan T, N)`) — отправитель блокируется только когда буфер полон. Асинхронная очередь.

Буфер — кольцевой массив с двумя индексами sendx и recvx.

Когда какой:
- Unbuffered — сигналы, `done chan struct{}`, важно «получатель взял»
- Buffered — очереди задач, развязывание producer/consumer, паттерн «первый ответ»

### 14. ⚡ nil-канал — зачем `M`
```go
var ch chan int  // nil
```
Чтение/запись в nil — блок навечно. Закрытие — паника. **Зачем нужен**: «выключить» ветку в `select`. Присвоил `ch = nil` — case больше не сработает.
```go
for ch1 != nil || ch2 != nil {
    select {
    case v, ok := <-ch1:
        if !ok { ch1 = nil; continue }
        out <- v
    case v, ok := <-ch2:
        if !ok { ch2 = nil; continue }
        out <- v
    }
}
```

### 15. Утечка горутин — типичные причины `M`
1. Горутина пишет в небуферизированный канал, читателя нет
2. Горутина читает range по каналу, никто не закроет
3. Забыли отменить context (`defer cancel()`)
4. HTTP без timeout

Обнаружение: `runtime.NumGoroutine()` метрикой, pprof goroutine `/debug/pprof/goroutine?debug=2`, библиотека `goleak` в тестах.

### 16. ⚡ Утечка в этом коде? `M`
```go
func slow() string {
    ch := make(chan string)
    go func() { ch <- fetchSlow() }()
    select {
    case s := <-ch: return s
    case <-time.After(1*time.Second): return "timeout"
    }
}
```
**ДА, утечка**. Если сработал `time.After`, горутина зависает на `ch <- ...` навсегда. Фикс: `make(chan string, 1)` — буферизированный канал примет значение даже без читателя.

### 17. select — что если несколько готовы `M`
**Случайный выбор**, не первый. Это by design — иначе голодание второго канала. Под капотом — `selectgo` перемешивает case, проверяет lock-free, паркует горутину сразу на всех каналах.

`default` — неблокирующая операция. Антипаттерн: busy-wait `for { select { case: ...; default: } }` = 100% CPU.

### 18. Fan-in (merge) `M`
```go
func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)
    for _, c := range channels {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch { out <- n }
        }(c)
    }
    go func() { wg.Wait(); close(out) }()
    return out
}
```

### 19. ⚡ First response wins (hedged requests) `S`
```go
func firstResponse(replicas []string) string {
    ch := make(chan string, len(replicas))  // ВАЖНО: буфер!
    for _, r := range replicas {
        go func(addr string) { ch <- fetch(addr) }(r)
    }
    return <-ch
}
```
Буфер критичен. Без буфера: первый ответ читается, остальные горутины блокируются на `ch <- ...` → утечка. С буфером — пишут и завершаются.

### 20. Pipeline pattern `M`
```go
func generate(nums ...int) <-chan int { /* шлёт в канал, close */ }
func square(in <-chan int) <-chan int { /* читает, шлёт в канал, close */ }
// generate(1,2,3) → square → ...
```
Каждая стадия — горутина. Параллелизм автоматический: пока square обрабатывает 1, generate уже шлёт 2. Backpressure встроенный: медленный consumer → канал заполнен → producer блокируется.

### 21. Направленные каналы `J`
```go
chan T     // двусторонний
<-chan T   // receive-only
chan<- T   // send-only
```
Контракт через тип. Двусторонний неявно конвертится в направленный. Используется в сигнатурах функций.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

