---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]]"
type: concept
description: "3 состояния: `pending` → `fulfilled` или `rejected`. Переход необратим. `.then`, `.catch`, `.finally`."
---

## Promises и async/await

### 14. Promise состояния `J`
3 состояния: `pending` → `fulfilled` или `rejected`. Переход необратим. `.then`, `.catch`, `.finally`.

### 15. Promise.all/race/allSettled/any `M`

| Метод | Ждёт | Reject когда | Возврат |
|---|---|---|---|
| `all` | все | первом reject | массив |
| `race` | первый | первом reject | значение |
| `allSettled` | все | никогда | `{status, value/reason}` |
| `any` | первый success | все reject | значение |

### 16. ⚡ Что выведет? `M`
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
**Ответ:** `A 1 B 2 3`. Код до первого `await` — синхронный. После `await` — microtask, выполняется после синхронного.

### 17. Параллель vs последовательно `M`
```js
// ❌ 3 сек
const a = await fetchA(); const b = await fetchB(); const c = await fetchC();
// ✅ 1 сек
const [a,b,c] = await Promise.all([fetchA(), fetchB(), fetchC()]);
```
Если запросы НЕ зависят друг от друга → `Promise.all`. Зависят → последовательно.

### 18. Timeout для Promise `M`
```js
const withTimeout = (p, ms) => Promise.race([
  p,
  new Promise((_, r) => setTimeout(() => r(new Error("Timeout")), ms))
]);
```
Идея — кто первый: результат или таймер.

### 19. AbortController `M`
```js
const c = new AbortController();
fetch(url, { signal: c.signal });
c.abort(); // catch с err.name === "AbortError"
```
В React: создавай в useEffect, abort в cleanup → автоотмена при размонтировании.

### 20. Параллелизм с лимитом `S`
```js
async function parallelLimit(tasks, limit) {
  const results = []; const executing = new Set();
  for (const [i, t] of tasks.entries()) {
    const p = t().then(r => { executing.delete(p); results[i] = r; });
    executing.add(p);
    if (executing.size >= limit) await Promise.race(executing);
  }
  await Promise.all(executing);
  return results;
}
```

### 21. Retry с backoff `M`
```js
async function retry(fn, retries=3, delay=1000) {
  for (let i = 0; i < retries; i++) {
    try { return await fn(); }
    catch (e) {
      if (i === retries - 1) throw e;
      await new Promise(r => setTimeout(r, delay * (i + 1)));
    }
  }
}
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

