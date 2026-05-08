---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "Простой CRUD сервис на 5 endpoints с CQRS, Event Sourcing, Saga. **Слишком сложно** для требований."
---

## CQRS — когда оверкилл

> «CQRS overkill для простой задачи микросервиса — рады что знаете, но не нужно.»

### ❌ Антипаттерн на собеседовании

Простой CRUD сервис на 5 endpoints с CQRS, Event Sourcing, Saga. **Слишком сложно** для требований.

### ✅ Когда уместно

- Read load в **10×+** больше write
- Сложная domain-логика на write side
- Event-driven архитектура
- Audit log требуется

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

