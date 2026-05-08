---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "NaN === NaN // false"
---

### TRAP 8. NaN === NaN

```js
NaN === NaN          // false
NaN == NaN           // false
Number.isNaN(NaN)    // true
isNaN("hello")       // true!
Number.isNaN("hello")// false
```

**Подвох:** глобальный `isNaN()` сначала приводит к Number, и "hello" → NaN → true. Используй `Number.isNaN()` — без приведения.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

