---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "typeof null // \"object\""
---

### TRAP 7. typeof null

```js
typeof null   // "object"
typeof undefined // "undefined"
typeof []     // "object"
typeof NaN    // "number"
```

**Подвох:** `typeof null === "object"` — исторический баг JS, не исправляют по обратной совместимости. `NaN` — тоже number. Для проверки массива — `Array.isArray()`. Для проверки NaN — `Number.isNaN()` (НЕ глобальный isNaN — у него другая логика).

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

