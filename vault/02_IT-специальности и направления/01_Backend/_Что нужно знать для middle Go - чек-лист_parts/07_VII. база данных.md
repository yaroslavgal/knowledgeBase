---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]]"
type: concept
description: "db.Query(\"SELECT * FROM users WHERE name = '\" + name + \"'\") // SQL-инъекция"
---

## VII. БАЗА ДАННЫХ

### Junior
```go
db.Query("SELECT * FROM users WHERE name = '" + name + "'")  // SQL-инъекция
```

### Middle
```go
db.QueryContext(ctx, "SELECT id, name, email FROM users WHERE name = $1", name)
// явные колонки + параметризованный запрос
```

**Ключи:**
- Всегда параметризованные запросы (`$1`, `$2` или `?`)
- Явный список колонок, не `SELECT *`
- `defer rows.Close()` после `Query`
- Транзакции через `BeginTx(ctx, ...)`
- HTTP-вызовы НИКОГДА внутри транзакции (долгоживущая транзакция = локи)
- `sqlc` или `goqu` для типобезопасности
- Понимаешь N+1 проблему

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

