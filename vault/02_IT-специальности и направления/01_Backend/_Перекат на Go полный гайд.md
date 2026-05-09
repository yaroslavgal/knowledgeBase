---
parent: "[[20_IT-специальности]]"
type: concept
description: "Полный гайд перехода в Go-бэкенд для разработчиков с Python/JS/Java. Окно дефицита Go-разработчиков может закрыться к 2027, поэтому план — на 2-4 месяца."
---

# Перекат на Go — полный гайд (с Python/JS/Java)

## Главный тезис

Это уникальная ситуация — рекрутёры понимают что Go-разработчиков мало, а синтаксис учится за пару вечеров. Окно может закрыться к 2027 году по мере насыщения рынка.

> «Знать сам язык всё равно ОБЯЗАТЕЛЬНО.»

---

## Срок переката

| Темп | Срок |
|---|---|
| Платный курс с ментором, full-time | **2–3 недели** |
| Самостоятельно с pet-проектами | 2–6 месяцев |
| «Параллельно с работой по 1 часу/день» | 6–12 месяцев |

**Реальные кейсы:** Senior Python → Senior Go в Ozon с 2 неделями подготовки. Получилось пройти собес, но **не прошёл испытательный срок** из-за отсутствия онбординга.

---

## День 1–2: Tour of Go

[tour.golang.org](https://tour.golang.org) — официальный туториал. Обязательное минимум.

**Что покрыть:**
- Синтаксис (var/const, control flow)
- Функции и multiple return
- Структуры и методы (нет классов, есть structs + methods + interfaces)
- Slices, maps, channels — базово
- Errors as values (нет exceptions)

**Что пропустить пока:**
- **Дженерики** (с 1.18) — «ничем не примечательны», изучишь на работе
- **Reflection** — экзотика
- **CGO** — экзотика для системных задач

---

## Сразу после Tour: инструментарий

> «Не нужно знать все встроенные инструменты, нужно знать основные.»

**Must:**
- `go mod init` / `go mod tidy` — модули
- `go build` / `go run` — компиляция
- `go test ./...` — тесты
- `go test -race ./...` — race detector
- `go test -bench ./...` — бенчмарки
- `go vet` / `golangci-lint` — статанализ
- `pprof` — профилирование (CPU, memory, goroutines)

---

## Конкурентность — отдельная тема

Tour of Go даёт **очень базово**. Нужно отдельное видео/курс.

**Что освоить:**
1. Горутины, каналы, select
2. `sync.Mutex`, `sync.WaitGroup`, `sync.Once`, `sync.RWMutex`
3. `context.Context` — cancel, deadline, values
4. Race detector
5. Patterns: Worker Pool, Pipeline, Fan-out/Fan-in
6. Memory model: happens-before гарантии

См. отдельно: `_Go_горутины_каналы_концентрат.md`

---

## Что НЕ нужно учить до выхода на работу

> «Если будешь несколько месяцев бесплатно фигачить микросервисы, ты просто выгорешь.»

❌ **Не нужно:**
- Полностью изучать Postgres / Redis / ClickHouse — научишься на работе
- Становиться экспертом в трейсинге / мониторинге — статью прочитал и хватит
- Писать свой микросервис с нуля — на работе будут шаблоны
- Свою систему авторизации / свою БД
- Глубокий k8s — на работе там будут DevOps-инженеры
- Все паттерны GoF наизусть

---

## База обязательно

✅ **Один веб-фреймворк:**
- `Gin` (популярный, простой)
- `Fiber` (Express-like, для тех кто из Node.js)
- `Echo` (минималистичный)
- `chi` (стандартная библиотека + роутинг)
- Просто `net/http` (без фреймворка) — для понимания

✅ **Одна БД:**
- **Postgres** — стандарт. Драйвер `pgx`. ORM `gorm` или query builder `sqlc`

✅ **Базовые инструменты:**
- Docker (написать Dockerfile, compose)
- Git (commit, branch, merge)
- Базовое представление о k8s (`kubectl`, манифесты)
- Базовое представление о CI/CD (GitHub Actions / GitLab CI)

---

## Особенности Go vs другие языки

### vs Python

| Python | Go |
|---|---|
| Динамическая типизация | Статическая |
| Exceptions | Errors as values (`if err != nil`) |
| Классы + наследование | Structs + composition (embedding) |
| GIL | Нет GIL, настоящие goroutines |
| Списки/словари — типы первого класса | Slices/maps — встроенные |
| `pip` | `go modules` |

### vs JavaScript

| JS | Go |
|---|---|
| Event loop, async/await | Goroutines + channels |
| Динамическая | Статическая |
| Prototype inheritance | Composition |
| `npm install` | `go get` |
| Колбеки / промисы | Channels / context |

### vs Java

| Java | Go |
|---|---|
| Verbose | Compact |
| OOP с классами | Structs + interfaces (duck typing) |
| Checked exceptions | Errors as values |
| JVM, GC ~100ms | Native, GC <10ms |
| Spring Boot | Стандартная библиотека + лёгкие фреймворки |

---

## Что переносится из Python/JS/Java

✅ **Переносится почти 1:1:**
- Базовые алгоритмы и структуры данных
- Понимание HTTP / REST / gRPC
- Понимание баз данных и SQL
- Опыт debugging
- Soft skills, командная работа
- Знание Linux / shell

✅ **Адаптируется:**
- ООП → структуры + интерфейсы (composition over inheritance)
- Async patterns → goroutines + channels
- Error handling → явный возврат ошибок

❌ **Не переносится:**
- Specific framework knowledge (Spring → Gin)
- Specific ORM patterns
- Language-specific tooling

---

## Подготовка к собесу при перекате

### Roadmap 256 от Ozon (бесплатно)
Инженеры Ozon рассказывают что нужно знать. Идёт по топикам.

### Платные курсы
- Stepik / Coursera / Boot.dev / Udemy — есть много
- Выбирай по отзывам, не по цене

### Лайвкод-задачи на Go (не LeetCode)
- Не алгоритмический гринд — а 15–20 типовых задач конкурентности и бизнес-логики, которые гоняют на Go-собесах: Worker Pool, fan-in/out, кэш с TTL, rate limiter, parsing, простые transformations.
- LeetCode-гринд (300 задач, two pointers, DP, graphs) имеет смысл только под bigtech (Яндекс algo-секция, Stripe-аналоги). На РФ-СНГ middle Go-собес алгоритмы спрашивают редко и формально, гоняют как раз внутрянку Go и concurrency-паттерны.

### Pet-проект
**Что писать:**
- REST API на Gin / Fiber
- Postgres + миграции
- Docker compose
- Тесты (unit + integration)
- README на английском
- Pet-проект нужен чтобы рассказать о нём на собесе, а не для GitHub-профиля — туда никто не заходит

**Примеры pet-проектов:**
- TODO-list backend (CRUD + JWT auth)
- URL shortener (REST + Redis)
- Telegram-bot (long polling)
- Chat backend (WebSocket)
- E-commerce mini-backend (товары + корзина + заказы)

---

## Подвохи на собесе для свитчеров

### `nil` interface vs interface с nil-значением

```go
var p *Person = nil
var i interface{} = p
i == nil  // false! ⚠️
```

Interface = (type, value). Если type есть, а value = nil — interface не равен nil.

### Встроенные значения по умолчанию

```go
var s string  // ""
var i int     // 0
var b bool    // false
var p *int    // nil
var sl []int  // nil (НЕ пустой)
var m map[string]int  // nil — нельзя писать без make
```

### Multiple return + named return

```go
func divide(a, b int) (result int, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return  // naked return с named values
    }
    result = a / b
    return
}
```

### Defer + LIFO

```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
// Output: 3, 2, 1
```

### Receiver type

```go
type MyType struct{ x int }

func (m MyType) Foo() {}    // value receiver — копия
func (m *MyType) Bar() {}   // pointer receiver — мутация возможна
```

---

## Карьерные карты после перехода

### Карта grade в Go

| Grade | Обязанности |
|---|---|
| **Junior** | Простые задачи с подробным описанием, требует менторства |
| **Middle** | Оценивает время, понимает другие направления, помогает джунам |
| **Senior** | Архитектура, ментор, декомпозирует задачи на джуновские |
| **Tech Lead** | Принимает архитектурные решения, технический контакт с PM |
| **Staff/Principal** | Системы уровня компании, mentoring сеньоров |

---

## Альтернативный roadmap (Захар, 65+ собесов)

1. **Досконально язык** (G-M-P, escape analysis, runtime каналов, GC) — это и есть главное на Go-собесах в РФ-СНГ
2. **Базы данных** (запросы, индексы, транзакции, изоляция)
3. **Компьютерные сети** (TCP/IP, HTTP, gRPC)
4. **DevOps** (Docker, CI/CD, базовый k8s)
5. **Конкурентность глубоко** — паттерны (Worker Pool, fan-in/out, pipeline, rate limiter, singleflight) с тренировкой на доске под таймер. Это и есть «лайвкод» на Go-собесе, а не LeetCode.

> Захар изначально топил за «300 задач LeetCode». В контексте этого vault это неверная рекомендация — для middle Go в РФ это переоценённый труд. На большинстве собесов алгоритмы либо не спрашивают, либо они формальные. Время лучше вложить в внутрянку языка и concurrency-паттерны.

---

## Подготовка к собесу — параллельно

| Активность | Темп |
|---|---|
| Лайвкод-задачи под Go (Worker Pool, кэш, rate limiter) | 1–2 задачи / день под таймер 30 мин |
| Pet-проект | 5–10 часов / неделю |
| Системный дизайн | 1 кейс / неделю |
| Mock-собес | 1 раз в 2 недели |
| Английский | 30 мин / день видео + 1 урок / неделю |

**Итого:** ~15–20 часов / неделю на 2–3 месяца → готов к собесу.

---

## Платные курсы

> «Платные курсы есть, но выбирайте сами.»

**Что искать в курсе:**
- ✅ Включает реальные собесные задачи
- ✅ Code review со ментором
- ✅ Pet-проект завершённый
- ✅ Подготовка к лайвкодингу
- ✅ Помощь с резюме

**Что игнорировать:**
- ❌ «Полный курс по Go 100 часов» (не нужно)
- ❌ Курсы без code review
- ❌ Без подготовки к собесам
- ❌ Только теория без практики

**Бесплатные альтернативы:**

- **Roadmap 256 от Ozon** — инженеры Ozon рассказывают что нужно знать
- **Go by Example** (gobyexample.com) — справочник
- **Effective Go** (golang.org/doc/effective_go) — официальный гайд
- **Ultimate Go Tour** (Ardanlabs) — глубокий курс
- **«The Go Programming Language»** (Donovan & Kernighan) — книга

---

## Чек-лист перехода на Go

- [ ] Tour of Go пройден
- [ ] Один pet-проект (REST API + Postgres + Docker + тесты) — для рассказа на собесе, не для GitHub-витрины
- [ ] Понимание `sync` пакета и каналов
- [ ] Race detector использовал
- [ ] 15–20 типовых лайвкод-задач под Go-собес отрепетированы (Worker Pool, кэш с TTL, rate limiter, fan-in/out, parsing)
- [ ] Один веб-фреймворк освоен (Gin/Fiber/chi)
- [ ] Базовое понимание k8s (kubectl, манифесты)
- [ ] System Design 3+ кейсов
- [ ] Резюме адаптировано: Go-релевантный опыт впереди
- [ ] Mock-собес на Go пройден 3+ раз
- [ ] Готов к лайвкодингу: 1 Easy + 1 Medium / час

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Что такое middle Go]] — целевой уровень переката: к чему именно ведём.
- [[02_IT-специальности и направления/01_Backend/_Переход с Python JS на Go за 2 недели|Переход с Python/JS на Go за 2 недели]] — конкретный план для конкретного бэкграунда.
- [[02_IT-специальности и направления/01_Backend/_Перекат с фронтенда в Go|Перекат с фронтенда в Go]] — план для фронтендеров с другими стартовыми условиями.
- [[02_IT-специальности и направления/01_Backend/_5 ошибок не берут на middle Go|5 ошибок не берут на middle Go]] — слепые зоны переката, которые палят сразу.
- [[01_Образование и обучение/05_Путь обучения по специальностям/_Roadmap входа в Go|Roadmap входа в Go]] — карта обучения для тех, кто идёт без существующего бэкграунда.
- [[03_Зарплаты и грейды/01_По специализациям/_Go зарплаты рф снг 2026|Go зарплаты РФ/СНГ 2026]] — мотивация переката: +50–100к в среднем при прочих равных.
- [[04_Найм и трудоустройство/05_Проблемы рынка 2026/_Истории провалов собесов 2026|Истории провалов собесов 2026]] — кейс Петра в Ozon: что делать после переката, чтобы не слететь на испытательном.
<!-- auto-links:end -->

