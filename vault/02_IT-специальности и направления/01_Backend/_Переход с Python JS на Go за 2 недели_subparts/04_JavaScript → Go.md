---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Переход с Python JS на Go за 2 недели|Переход с Python / JS на Go за 2 недели]]"
type: concept
description: "---"
---

### JavaScript → Go

| JS | Go | Заметка |
|---|---|---|
| `function name(x) {}` | `func name(x int) error` | Типы обязательны |
| `class User { }` | `type User struct { ... }` | Композиция вместо наследования |
| `object {}` | `struct{...}` или `map[string]any` | Лучше struct |
| `array []` | `[]int` (slice) | Типизированный |
| `try { } catch (e)` | `if err != nil` | Без исключений |
| `Promise.all([...])` | `errgroup.Wait()` | Реальные потоки |
| `await fetch()` | `http.Get()` (синхронно) | Параллелизм через горутины |
| `EventEmitter` | `chan T` | Каналы вместо событий |
| Spread `{...obj}` | `*newObj = *obj` или копия | Указатель / значение явно |
| `null/undefined` | `nil` | Только указатели и ссылочные могут быть nil |

---

## 6 концепций, на которых спотыкаются все

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Переход с Python JS на Go за 2 недели|Переход с Python / JS на Go за 2 недели]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

