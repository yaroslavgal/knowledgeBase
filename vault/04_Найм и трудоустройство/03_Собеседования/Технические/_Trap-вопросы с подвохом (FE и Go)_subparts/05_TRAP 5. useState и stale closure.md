---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]]"
type: concept
description: "function Counter() {"
---

### TRAP 5. useState и stale closure

```js
function Counter() {
  const [count, setCount] = useState(0);
  const handleClick = () => {
    setCount(count + 1);
    setCount(count + 1);
    setCount(count + 1);
  };
  return <button onClick={handleClick}>{count}</button>;
}
```

**После клика:** count = 1 (не 3).

**Подвох:** `count` — замыкание, зафиксированное при рендере. Все три вызова используют одно и то же старое значение. Нужно функциональное обновление: `setCount(prev => prev + 1)`.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Trap-вопросы с подвохом (FE и Go)|Trap-вопросы с подвохом — FE и Go]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

