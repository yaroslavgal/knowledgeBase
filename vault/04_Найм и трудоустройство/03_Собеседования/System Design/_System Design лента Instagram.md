# System Design: проектирование ленты Instagram (Mock от TL Яндекса)

Спроектировать ленту Instagram на 50 миллионов пользователей в СНГ.

## Условие задачи

---

## Структура ответа (45–60 минут)

### 1. Функциональные требования (5 мин)

Что система делает:
- Загрузка поста (фото + описание)
- Получение ленты подписок
- Лайки / комментарии со звёздочкой
- Подписки / отписки
- Удаление поста (опционально)

### 2. Нефункциональные требования (5 мин)

| Параметр | Значение |
|---|---|
| MAU | 50M |
| Гео | СНГ (один регион) |
| Загрузка поста | ~2 сек |
| Получение ленты | ~1 сек |
| Хранение | Навсегда |

**Обсудить с интервьюером:** почему лента дольше? — Нужно агрегировать посты от подписок.

### 3. Поведение пользователя (3 мин)

- 1 пост / 5 дней
- 5 просмотров ленты / день
- 10 постов на запрос ленты

### 4. Расчёт нагрузки (5 мин)

```
RPS write = 50M / 5 / 86400 ≈ 115 RPS
RPS read  = 50M × 5 / 86400 ≈ 2900 RPS
```

### 5. Расчёт трафика (3 мин)

```
Размер поста = 100 символов × 2 байта + 0.5 MB картинка ≈ 500 KB

Write throughput = 115 × 500 KB = 57 MB/s
Read throughput  = 2900 × 500 KB × 10 = 14 GB/s
```

### 6. Capacity на год (2 мин)

```
Только write: 80 MB/s × 86400 × 365 ≈ 2.5 PB
```

### 7. API псевдокод (5 мин)

```
POST   /post           {description, image}
GET    /feed           ?limit=10&offset=0
POST   /follow         {target_user_id}
DELETE /follow/{id}
POST   /post/{id}/like
POST   /post/{id}/comment {text}
```

### 8. Архитектура (15 мин)

```
[Client]
   ↓
[CDN/Image Storage] ← S3 (картинки)
   ↓
[Load Balancer] / [API Gateway]
   ↓
┌─────────┬─────────────┬──────────────┐
│ User    │ Feed        │ Post         │  Микросервисы
│ Service │ Service     │ Service      │
└────┬────┴──────┬──────┴───────┬──────┘
     ↓           ↓              ↓
[Postgres]  [Redis]      [Postgres]
  users      hot feed     posts/comments
[Neo4j или Postgres] follows  ←  граф подписок
```

### 9. Дата-модель (5 мин)

```sql
users (id, name, created_at)
posts (id, user_id, description, image_url, created_at)
follows (follower_id, following_id, created_at)  -- composite PK
likes (user_id, post_id, created_at)
comments (id, user_id, post_id, text, created_at)
```

### 10. Железо (часто не успевают)

- API: 5 серверов × 16 cores × 64GB RAM
- Postgres: master + 2 read replicas
- Redis: 3-node cluster, 100GB total
- S3: managed, оплата по объёму
- Кол-во шардов / партиций

---

## Топ-проблемы и решения

### Проблема 1: Fan-out (миллион подписок)

**Сценарий:** пользователь подписан на миллион аккаунтов. Как собрать ленту?

**Решения:**
- **Pull (read-time aggregation):** при запросе ленты — `SELECT * FROM posts WHERE user_id IN (subscriptions) ORDER BY created_at LIMIT 10` — медленно
- **Push (write-time fan-out):** при публикации — кладём пост в feed-cache всех подписчиков сразу
- **Hybrid:** для celebrity (>10K подписчиков) — Pull (на read), для обычных — Push

### Проблема 2: Атомарность загрузки

**Сценарий:** картинка загрузилась в S3, а описание не сохранилось в Postgres → orphan-картинка.

**Решения:**
- **Two-phase commit** (сложно, плохая производительность)
- **Saga pattern:** компенсирующие действия при сбое
- **Outbox pattern:** транзакционно записать «событие» + воркер публикует в S3

### Проблема 3: Hot key (популярный пост)

**Сценарий:** пост звезды получает 10K RPS на чтение.

**Решения:**
- Кэшируем агрессивно (Redis + CDN)
- Repltcaten на чтение
- Sharding по post_id

### Проблема 4: Кэш-инвалидация

**Сценарий:** пользователь удалил пост — он остаётся в Redis ленты у подписчиков.

**Решения:**
- TTL на feed-cache (5–10 мин)
- Pub/Sub для уведомления о изменениях
- Версионность ленты

---

## Типичные ошибки кандидата (фидбэк интервьюера)

❌ **Не уточнили лимиты** (макс. число подписок, размер картинки)
❌ **Не формализовали что такое «пост»** — сразу к дизайну
❌ **Баги в расчётах трафика** — «из-за этих расчётов определяется архитектура»
❌ **Выбор БД пальцем в небо** — Redis для основного хранилища без оценки объёма; «Postgres потому что я работал» без аргументов нагрузка/железо
❌ **Не обсуждены corner case'ы** (fan-out, атомарность, hot keys)
❌ **Не дошли до железа** = недозащищён дизайн

---

## Шаблоны для других задач

### Design URL Shortener (TinyURL)
- Hashing (base62, MD5 truncate)
- Hot keys → CDN cache
- Analytics через async event queue

### Design Twitter
- Pull vs Push разделение
- Timeline materialization
- Hashtag indexing (Inverted index)

### Design WhatsApp
- WebSocket connections (sticky session)
- Message persistence (Cassandra)
- Read receipts через ACK

### Design Uber
- Real-time location (Redis Geo)
- Driver-rider matching (geo-hash)
- Dynamic pricing (event-driven)

### Design Rate Limiter
- Token bucket
- Sliding window
- Distributed Redis-based

### Design Distributed Cache
- Consistent hashing
- LRU eviction
- Write-through vs write-behind

---

## Паттерны на System Design в Ozon

«В Озоне точно спросят»:
- **Saga** (распределённые транзакции, компенсирующие действия)
- **Transactional Outbox** (надёжная отправка событий из сервиса в кафку)
- **CQRS** (Command Query Responsibility Segregation)
- **Event Sourcing**

---

## Подготовка

| Источник | Что |
|---|---|
| **«Designing Data-Intensive Applications»** (Kleppmann) | Must read |
| **ByteByteGo** (Alex Xu) | Сборники задач |
| **Hired In Tech** | Mock-интервью с разбором |
| **System Design Primer** (github.com/donnemartin/system-design-primer) | Бесплатный гайд |
| **Grokking System Design Interview** | Educative.io |
| **High Scalability blog** | Реальные кейсы крупных компаний |

---

## Чек-лист системного дизайна

- [ ] 5+ классических задач разобрано
- [ ] Знаешь как делать back-of-envelope estimation (RPS, throughput, capacity)
- [ ] Понимаешь CAP theorem и trade-offs
- [ ] Знаешь основные БД и когда какую использовать (SQL/NoSQL/Graph)
- [ ] Знаешь паттерны (Pub/Sub, Saga, CQRS, Outbox)
- [ ] Можешь обсуждать кэширование (Redis, Memcached, CDN)
- [ ] Знаешь load balancing (L4 vs L7, sticky sessions)
- [ ] Понимаешь sharding / partitioning стратегии
- [ ] Можешь обсудить мониторинг и observability
- [ ] Mock-интервью пройдено 3+ перед реальным

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/System Design/_Проектирование систем|Проектирование систем]] — общий теоретический материал, на котором стоит этот кейс.
- [[05_Рабочие процессы и эффективность/01_AI инструменты/AI-ассистенты на собесе/_System Design подготовка через Claude Code|System Design подготовка через Claude Code]] — как тренироваться на этом кейсе через AI.
- [[04_Найм и трудоустройство/05_Проблемы рынка 2026/_Истории провалов собесов 2026|Истории провалов собесов 2026]] — кейс из паттерна 5 (Senior без System Design).
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Что такое middle Go]] — System Design 3+ кейсов — пункт middle-чек-листа.
<!-- auto-links:end -->

