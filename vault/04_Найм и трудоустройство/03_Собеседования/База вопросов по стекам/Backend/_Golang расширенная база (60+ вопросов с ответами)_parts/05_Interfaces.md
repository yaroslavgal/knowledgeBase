---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "Интерфейс = пара `(type, value)` — динамический тип + указатель на значение. Реализация неявная (structural/duck typing) — если тип имеет нужные методы, он реализует."
---

## Interfaces

### 37. Структура интерфейса `M`
Интерфейс = пара `(type, value)` — динамический тип + указатель на значение. Реализация неявная (structural/duck typing) — если тип имеет нужные методы, он реализует.

### 38. ⚡ nil-интерфейс ловушка `M`
```go
type MyError struct{}
func (*MyError) Error() string { return "bad" }
func doIt() error {
    var err *MyError = nil
    return err
}
if err := doIt(); err != nil {
    fmt.Println("WTF?") // выведется!
}
```
Интерфейс != nil, потому что type != nil (`*MyError`), хотя value = nil. Фикс: `return nil` явно вместо `return err`.

### 39. ⚡ Method set для T vs *T `M`
Если метод через pointer receiver `(*T)`:
- `t := T{}` — НЕ реализует интерфейс
- `&t` — реализует

Method set у T — только value receivers, у *T — и value и pointer.

### 40. Pointer receiver vs value receiver `M`
- Pointer receiver — нужно изменять состояние, тип большой (избежать копирования), консистентность с другими методами
- Value receiver — не меняет, маленький тип, immutable

Правило: если есть хоть один метод с pointer receiver — все методы тоже с pointer receiver (для консистентности).

### 41. Accept interfaces, return structs `M`
Потребитель зависит от минимального контракта (легко тестировать/мокать). Возврат struct даёт caller полный API. Интерфейс определять рядом с потребителем, не с реализацией.

### 42. Compile-time interface check `M`
```go
var _ MyInterface = (*MyType)(nil)
```
Идиома. Если *MyType не реализует MyInterface — compile error.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

