# Live coding Go — реальные задачи с собесов 2026

В 60% собесов на Go есть лайвкодинг (Антон, 25 собесов).

## Контекст

В **60% собесов на Go** есть лайвкодинг (Антон, 25 собесов).

> «Если не умеешь писать код в блокноте, общаться, проектировать, тестировать — не пройдёшь 6 из 10 собесов.»

Готовность к Senior уровню: **1 Easy + 1 Medium за 1 час** на твоём языке.

---

## 4-этапный гайд решения лайвкодинга

### 1. Получение задачи (3–5 мин)
- Прочитать задачу
- Уточнить функциональные/нефункциональные требования
- Обсудить ограничения (размер данных, особые условия)
- **Написать тест-кейсы** (граничные + обычные) до кода

### 2. Проектирование (5–10 мин)
- Для алгоритмов: проверить Big-O (time + space)
- Для платформенных: обсудить архитектуру/паттерны (DRY, KISS, SOLID)
- Naive решение вслух → оптимизация

### 3. Кодирование (15–25 мин)
- **Проговаривай вслух где находишься**
- Разбивай на подзадачи
- Проси подсказку если застрял — «**любая решённая задача лучше нерешённой**»

### 4. Тестирование (5–10 мин)
- Прогон граничных кейсов
- Обработка ошибок (упал бэкенд, плохой ответ сервиса)

---

## Реальные задачи с Go-собесов

### Задача 1: In-memory cache с TTL (Купер, Middle+/Senior)

**Условие:** структура `Profile` с заказами внутри. Сделать конкурентнобезопасный кэш с TTL.

**Скелет решения:**
```go
type Cache struct {
    mu   sync.RWMutex
    data map[string]entry
}

type entry struct {
    profile Profile
    expires time.Time
}

func (c *Cache) Set(key string, p Profile, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = entry{profile: deepCopy(p), expires: time.Now().Add(ttl)}
}

func (c *Cache) Get(key string) (Profile, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    e, ok := c.data[key]
    if !ok || time.Now().After(e.expires) {
        return Profile{}, false
    }
    return deepCopy(e.profile), true
}
```

**Подвохи интервьюера:**
1. **Возврат указателя** позволит мутацию профиля в кэше → нужно **глубокое копирование**
2. **Заказы в профиле** — slice, который тоже нужно копировать (тот же reference issue)
3. **Get не должен возвращать просрочки** — проверять время прямо в Get
4. Corner case: «что если кто-то заберёт профиль, поменяет имя — изменится ли в кэше?»

---

### Задача 2: Worker Pool (Ozon, частая)

**Условие:** N горутин-воркеров слушают канал тасок, после `Close()` ждать выполнения всех задач.

**Скелет решения:**
```go
type Pool struct {
    mu       sync.Mutex
    isClosed bool
    tasks    chan func()
    done     chan struct{}
}

func New(workers int) *Pool {
    p := &Pool{
        tasks: make(chan func(), 100),
        done:  make(chan struct{}),
    }
    var wg sync.WaitGroup
    wg.Add(workers)
    for i := 0; i < workers; i++ {
        go func() {
            defer wg.Done()
            for task := range p.tasks {
                task()
            }
        }()
    }
    go func() {
        wg.Wait()
        close(p.done)
    }()
    return p
}

func (p *Pool) AddTask(task func()) error {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.isClosed {
        return errors.New("pool closed")
    }
    p.tasks <- task
    return nil
}

func (p *Pool) Close() {
    p.mu.Lock()
    p.isClosed = true
    close(p.tasks)
    p.mu.Unlock()
    <-p.done
}
```

**Подвохи:**
- `wg.Add(workers)` **до** запуска горутин (иначе race с Wait)
- `close(tasks)` после `isClosed=true`, чтобы не было send в закрытый канал

---

### Задача 3: Конкурентный стек

**Условие:** реализация стека на slice + sync.Mutex с методами Top/Pop.

**Подвохи:**
1. **`sync.Mutex` нельзя копировать** («золотое правило») → нужны pointer receivers
2. Доступ к `len(slice)` тоже надо синхронизировать (data race на чтении длины во время append)
3. **API проблема:** между Top() и Pop() другая горутина может изменить стек → нужен **один атомарный метод** вместо двух (например, `PopTop()` который возвращает и удаляет)

```go
type Stack[T any] struct {
    mu   sync.Mutex
    data []T
}

func (s *Stack[T]) Push(v T) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.data = append(s.data, v)
}

// Атомарный метод — вместо отдельных Top + Pop
func (s *Stack[T]) PopTop() (T, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    var zero T
    if len(s.data) == 0 {
        return zero, false
    }
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return v, true
}
```

---

### Задача 4: Буфер с возвратом slice

**Условие:** `Add(item)` под мьютексом, `Data()` возвращает slice.

**Подвох:** возврат slice наружу — пользователь может **читать без блокировки + append'ить** → race condition.

**Решение:** возвращать **копию** через `append([]T{}, original...)`:

```go
func (b *Buffer) Data() []int {
    b.mu.RLock()
    defer b.mu.RUnlock()
    return append([]int{}, b.data...)
}
```

---

## Концепции которые проверяют

### sync пакет — золотые правила

- Объекты `sync` (Mutex, WaitGroup, Once, RWMutex) **нельзя копировать** — pointer receivers обязательны
- `sync.Once.Do(f)` — f вызовется ровно один раз даже если N горутин зайдут одновременно
- `sync.WaitGroup.Add(n)` **всегда до** запуска горутин (иначе race с Wait)
- `sync.Map` хорош для read-heavy + редкие записи; для write-heavy → `map+RWMutex`
- `sync.Pool` — переиспользование объектов, ослабляет GC давление (но Pool не гарантирует сохранение)

### Race detector

`go test -race`, `go run -race` — **обязательно перед PR**.

### Закрытие каналов

- `close` nil channel = **panic**
- `close` уже закрытого канала = **panic**
- Send в закрытый канал = **panic**
- Receive из закрытого = **zero value + ok=false** (нормально)

### Memory model: happens-before гарантии

- Channel send happens-before receive
- Mutex Unlock happens-before следующего Lock
- `once.Do.f` happens-before любое другое Do

---

## Типичные ошибки на лайвкодинге

❌ **Молчание** — интервьюер не может помочь, не понимает где ты застрял
❌ **Спешка** — пропуск проектирования → переписывать к концу
❌ **Громоздкий код** — нечитаемый, не модульный
❌ **Неоптимальное решение без обсуждения** — даже если работает, минус
❌ **Игнорирование тестирования** — не покрыть corner cases
❌ **Бросание задачи** — лучше частично решить чем не решить совсем

---

## Топ-3 концептуальные западни Go-конкурентности

1. **Copy of sync.Mutex** — value receiver на структуре с Mutex = race
2. **Неатомарный API** из синхронизированных методов — между Top() и Pop() гонка
3. **Возврат slice/pointer наружу** из под mutex — пользователь меняет внутреннее состояние

---

## Чек-лист подготовки

- [ ] 300+ LeetCode задач решено (Easy + Medium на Go)
- [ ] Все паттерны конкурентности отрепетированы (Worker Pool, Pipeline, Fan-out/Fan-in)
- [ ] Знаешь `sync` пакет (Mutex, RWMutex, WaitGroup, Once, Map, Pool)
- [ ] Понимаешь модель G-M-P scheduler
- [ ] Понимаешь как работают каналы под капотом (hchan, sendq/recvq)
- [ ] Знаешь race detector (`go test -race`)
- [ ] 5+ mock-собесов с лайвкодингом пройдено
- [ ] Привычка проговаривать решение вслух
- [ ] Тесты пишешь до кода (TDD-подход на собесе)

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Структура мок-собеса Go (30 мин)|Структура мок-собеса Go]] — формат для тренировки этих задач.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Теория и лайвкодинг|Теория и лайвкодинг]] — общий разбор: чем лайвкодинг отличается от теории.
- [[02_IT-специальности и направления/01_Backend/_Горутины и каналы за час - концентрат|Горутины и каналы за час]] — concurrency-концентрат, обязательный для лайвкодинга Go.
- [[02_IT-специальности и направления/01_Backend/_5 ошибок не берут на middle Go|5 ошибок не берут на middle Go]] — типовые ошибки, которые палятся именно на лайвкодинге.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_От нуля до live coding senior — 20 задач|От нуля до live coding senior]] — путь подготовки до уровня этих реальных задач.
- [[05_Рабочие процессы и эффективность/01_AI инструменты/AI-ассистенты на собесе/_Mock-собес через Claude Code|Mock-собес через Claude Code]] — тренировка лайвкодинга с AI как спарринг-партнёром.
<!-- auto-links:end -->

