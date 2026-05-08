---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "if err != nil {"
---

## Errors — паттерны

### ❌ Антипаттерн: bare error

```go
if err != nil {
    return err  // ← где произошла ошибка?
}
```

### ✅ Правильно: error wrapping

```go
if err != nil {
    return fmt.Errorf("getUser: %w", err)
}
```

### Sentinel errors

```go
var (
    ErrUserNotFound = errors.New("user not found")
    ErrInvalidInput = errors.New("invalid input")
)

// Caller:
if errors.Is(err, ErrUserNotFound) { ... }
```

### Custom error types

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Caller:
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Printf("Field: %s\n", validationErr.Field)
}
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

