---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "async function foo() {"
---

### TRAP 3. Async ordering with await

```js
async function foo() {
  console.log("1");
  const result = await Promise.resolve("2");
  console.log(result);
  console.log("3");
}
console.log("A");
foo();
console.log("B");
```

**Ответ:** `A 1 B 2 3`.

**Подвох:** код до первого `await` — СИНХРОННЫЙ. Поэтому A → 1 → B → 2 → 3, а не A → 1 → 2 → 3 → B. Всё после `await` становится microtask.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

