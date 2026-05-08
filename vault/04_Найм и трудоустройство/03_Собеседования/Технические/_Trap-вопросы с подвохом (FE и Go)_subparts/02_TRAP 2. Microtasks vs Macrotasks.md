---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "console.log(\"1\");"
---

### TRAP 2. Microtasks vs Macrotasks

```js
console.log("1");
setTimeout(() => console.log("2"), 0);
Promise.resolve().then(() => console.log("3"));
console.log("4");
```

**Ответ:** `1 4 3 2`.

**Подвох:** многие говорят `1 4 2 3` потому что setTimeout(..., 0) кажется быстрее. На самом деле:
1. Сначала весь синхронный код (1, 4)
2. Потом ВСЯ microtask queue (Promise — это microtask) → 3
3. Только потом ОДНА macrotask (setTimeout) → 2

Microtasks ВСЕГДА выполняются до следующей macrotask. Даже setTimeout с 0 ждёт.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

