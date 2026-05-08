---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "type Counter struct {"
---

## Receiver type — value vs pointer

### Правило (для собесов с компаниями типа Ozon, Купер)

| Используй pointer receiver если | Используй value receiver если |
|---|---|
| Метод **меняет** состояние | Объект **неизменяемый** (immutable) |
| Структура содержит **mutex** или другие sync-объекты | Объект **маленький** (≤2-3 поля) |
| Структура **большая** (копирование дорого) | Безопасность (защита от мутации) |

### ❌ Антипаттерн

```go
type Counter struct {
    mu sync.Mutex
    n  int
}

func (c Counter) Inc() {  // ← value receiver
    c.mu.Lock()           // ← мьютекс копировался → не работает!
    defer c.mu.Unlock()
    c.n++
}
```

### ✅ Правильно

```go
func (c *Counter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.n++
}
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

