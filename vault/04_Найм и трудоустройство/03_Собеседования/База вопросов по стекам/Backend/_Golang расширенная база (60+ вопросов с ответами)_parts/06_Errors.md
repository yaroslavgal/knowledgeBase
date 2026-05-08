---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "return fmt.Errorf(\"doing X: %w\", err)"
---

## Errors

### 43. Оборачивание ошибок `M`
```go
return fmt.Errorf("doing X: %w", err)
```
`%w` оборачивает (errors.Is/As работают). `%v` — только подставляет строку, тип теряется.

### 44. errors.Is vs errors.As `M`
- `Is` — сравнение со sentinel error: `errors.Is(err, sql.ErrNoRows)`
- `As` — извлечение по типу: `var pgErr *pgconn.PgError; errors.As(err, &pgErr)`

⚡ `err == sql.ErrNoRows` — НЕ гарантировано работает, если кто-то в цепочке обернул через `%w`. Всегда `errors.Is`.

### 45. errors.Join (Go 1.20+) `M`
```go
return errors.Join(err1, err2, err3)
```
Объединяет несколько ошибок. До 1.20 — multierror библиотеки.

### 46. Defer и tx.Rollback `M`
```go
defer func() {
    if p := recover(); p != nil { tx.Rollback(); panic(p) }
    else if err != nil { tx.Rollback() }
    else { err = tx.Commit() }
}()
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

