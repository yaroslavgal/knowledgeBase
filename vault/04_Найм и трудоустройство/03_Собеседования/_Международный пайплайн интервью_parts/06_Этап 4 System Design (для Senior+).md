---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/_Международный пайплайн интервью|Международный пайплайн интервью (от LinkedIn до оффера)]]"
type: concept
description: "Классические задачи:"
---

## Этап 4: System Design (для Senior+)

### Что спрашивают

Классические задачи:
- Design Twitter / Instagram feed
- Design URL shortener (TinyURL)
- Design WhatsApp / messaging system
- Design YouTube / Netflix streaming
- Design Uber / Lyft ride matching
- Design rate limiter
- Design distributed cache
- Design search engine indexing
- Design payment processing system
- Design ad recommendation system

### Структура ответа (45–60 мин)

1. **Requirements** (5 мин)
   - Functional: что система делает
   - Non-functional: scale, latency, availability
   - Constraints: budget, team, timeline
2. **Estimation** (5 мин)
   - QPS / RPS
   - Storage
   - Bandwidth
3. **High-level design** (15 мин)
   - Components: services, DBs, queues, caches
   - Data flow
4. **Deep dive** (15 мин)
   - 1–2 компонента подробнее по запросу интервьюера
5. **Bottlenecks / scaling** (5 мин)
   - Where to optimize
   - How to handle 10× load
6. **Q&A** (5 мин)

### Подготовка

- **«Designing Data-Intensive Applications»** (Martin Kleppmann) — must read
- **ByteByteGo** (alex xu) — сборники классических задач
- **Hired In Tech** — mock интервью с разбором
- **System Design Primer** (github.com/donnemartin/system-design-primer) — бесплатный гайд

### Типичные ошибки

- Прыгаешь сразу в детали без requirements
- Не делаешь back-of-envelope estimation
- Игнорируешь non-functional requirements
- Перегружаешь сразу всеми технологиями (Kafka, Redis, ElasticSearch, Cassandra…)
- Не обосновываешь trade-offs (CAP theorem, ACID vs BASE)

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/_Международный пайплайн интервью|Международный пайплайн интервью (от LinkedIn до оффера)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

