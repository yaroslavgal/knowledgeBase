---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "Promise.resolve()"
---

### TRAP 4. Promise unwrapping

```js
Promise.resolve()
  .then(() => { console.log(1); return Promise.resolve(2); })
  .then(console.log);

Promise.resolve()
  .then(() => console.log(3))
  .then(() => console.log(4))
  .then(() => console.log(5));
```

**Ответ:** `1 3 4 2 5`.

**Подвох:** `return Promise.resolve(2)` — это thenable. По спецификации, unwrapping thenable создаёт **2 дополнительных microtask'а**, поэтому `2` появляется позже ожидаемого. Это уровень Senior, спрашивают чтобы отделить «знаю Promise» от «понимаю спецификацию».

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

