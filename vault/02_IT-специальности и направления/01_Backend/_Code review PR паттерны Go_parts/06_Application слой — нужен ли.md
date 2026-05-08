---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "// usecase/user_service.go"
---

## Application слой — нужен ли

### ❌ Антипаттерн: application слой который ничего не делает

```go
// usecase/user_service.go
func (s *UserService) GetUser(id int) (*User, error) {
    return s.repo.GetUser(id)  // ← просто прокси
}
```

**Проблема:** этот слой не добавляет ценности → удалить.

### ✅ Когда application слой нужен

- Композиция нескольких репозиториев
- Бизнес-логика
- Транзакции
- Применение domain events

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

