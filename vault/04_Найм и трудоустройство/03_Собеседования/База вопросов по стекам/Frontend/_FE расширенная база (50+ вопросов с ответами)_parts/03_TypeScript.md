---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]]"
type: concept
description: "Правило: объект/API → interface. Union/утилитарный/комбинация → type."
---

## TypeScript

### 22. interface vs type `M`
- `interface` — формы объектов, поддерживает `extends` и declaration merging
- `type` — для всего остального (union, intersection, primitives, mapped, conditional)

Правило: объект/API → interface. Union/утилитарный/комбинация → type.

### 23. Generics + constraints `M`
```ts
function first<T>(arr: T[]): T | undefined { return arr[0]; }
function getProp<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}
```
`T extends ...` — ограничение. `keyof T` — union ключей.

### 24. Utility Types `M`
- `Partial<T>` — все опциональны
- `Required<T>` — все обязательны
- `Pick<T, K>` / `Omit<T, K>`
- `Record<K, V>` — словарь
- `ReturnType<F>` — тип возврата
- `Readonly<T>` — readonly поля

### 25. Реализуй свой Partial `M`
```ts
type MyPartial<T> = { [K in keyof T]?: T[K] };
type MyPick<T, K extends keyof T> = { [P in K]: T[P] };
type MyOmit<T, K extends keyof T> = MyPick<T, Exclude<keyof T, K>>;
```

### 26. Discriminated Unions `M`
```ts
type Shape = { kind: "circle"; r: number } | { kind: "square"; s: number };
function area(s: Shape) {
  switch (s.kind) {
    case "circle": return Math.PI * s.r ** 2; // TS знает: circle
    case "square": return s.s ** 2;
  }
}
```
Exhaustive check через `never` — компилятор ругнётся если забудешь case.

### 27. Type Guards и narrowing `M`
- `typeof x === "string"`
- `"prop" in x`
- `x instanceof Class`
- Custom: `function isUser(x: unknown): x is User { ... }`

### 28. unknown vs any `J`
- `any` — отключает проверку, ошибки в рантайме
- `unknown` — безопасно, нельзя ничего делать пока не проверишь

Правило: никогда `any`. Используй `unknown` + type narrowing.

### 29. Conditional Types + infer `S`
```ts
type UnwrapPromise<T> = T extends Promise<infer U> ? U : T;
```
`infer` — извлекает тип «изнутри» (Promise, массив, функция).

### 30. Mapped + Template Literal Types `S`
```ts
type Getters<T> = { [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K] };
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]] — документ-оглавление, частью которого является эта заметка.
- [[02_IT-специальности и направления/02_Frontend/_Реальный фронтенд 2026 - не курсы|Реальный фронтенд 2026]] — TypeScript — обязательный 60–70% всех Frontend-вакансий 2026, без него собес не пройти.
<!-- auto-links:end -->

