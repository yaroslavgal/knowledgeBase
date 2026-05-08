---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Переход с Python JS на Go за 2 недели|Переход с Python / JS на Go за 2 недели]]"
type: concept
description: "Справочник синтаксиса: как перевести class, dict, try/except, asyncio, queue из Python в Go-конструкции"
---

### Python → Go

| Python | Go | Заметка |
|---|---|---|
| `def func(x):` | `func name(x int) error` | Возвращай ошибку явно |
| `class User:` | `type User struct { ... }` | Без конструктора, без наследования |
| `dict` | `map[string]int` | Жёсткие типы ключа и значения |
| `list` | `[]int` (slice) | Динамический массив |
| `try/except` | `if err != nil { return err }` | Ошибки — обычные значения |
| `asyncio.create_task` | `go func() { ... }()` | Горутина легче, чем coroutine |
| `await asyncio.gather(...)` | `errgroup.Wait()` | Группа горутин с ошибками |
| `queue.Queue` | `chan T` | Каналы — главный примитив общения |
| `with lock:` | `mu.Lock(); defer mu.Unlock()` | Мьютексы есть, но каналы лучше |
| `id(obj)` | `&obj` | Явные указатели |

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Переход с Python JS на Go за 2 недели|Переход с Python / JS на Go за 2 недели]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

