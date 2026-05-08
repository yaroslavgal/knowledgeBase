---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]]"
type: concept
description: "/* 1. Flex */ .p { display: flex; justify-content: center; align-items: center; }"
---

## CSS

### 43. Flexbox vs Grid `J`
- Flex — одномерный (строка ИЛИ столбец)
- Grid — двумерный (строки И столбцы)

### 44. Центрирование — 5+ способов `J`
```css
/* 1. Flex */ .p { display: flex; justify-content: center; align-items: center; }
/* 2. Grid */ .p { display: grid; place-items: center; }
/* 3. Absolute+transform */ .c { position:absolute; top:50%; left:50%; transform:translate(-50%,-50%); }
/* 4. margin auto */ .c { margin: 0 auto; }
/* 5. text-align + line-height */ .p { text-align:center; line-height: 200px; }
/* 6. table-cell + vertical-align */
```

### 45. Специфичность CSS `M`
Тег (0,0,0,1) → класс/атрибут (0,0,1,0) → ID (0,1,0,0) → inline (1,0,0,0) → !important. При равной — последнее правило.

### 46. ⚡ z-index не работает `M`
`z-index` работает только на positioned. Сравнивается внутри **stacking context**. Контекст создают: `position+z-index`, `opacity<1`, `transform`, `filter`, `will-change`. Типичный баг — `z-index: 9999` бесполезен, потому что родитель создал свой stacking context с меньшим z-index.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Frontend/_FE расширенная база (50+ вопросов с ответами)|FE расширенная база — 50+ вопросов с короткими ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

