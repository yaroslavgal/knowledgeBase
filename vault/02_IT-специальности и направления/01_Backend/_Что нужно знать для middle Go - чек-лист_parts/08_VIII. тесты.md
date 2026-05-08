---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]]"
type: concept
description: "В пет-проекте — ноль тестов. На собесе: \"Я знаю про unit-тесты, просто не успел добавить\"."
---

## VIII. ТЕСТЫ

### Junior
В пет-проекте — ноль тестов. На собесе: "Я знаю про unit-тесты, просто не успел добавить".

### Middle
- Юнит-тесты для бизнес-логики (`*_test.go`)
- Table-driven tests:
  ```go
  tests := []struct{
      name string
      input string
      want int
      wantErr bool
  }{
      {"valid", "abc", 3, false},
      {"empty", "", 0, true},
  }
  for _, tt := range tests {
      t.Run(tt.name, func(t *testing.T) { ... })
  }
  ```
- Моки через интерфейсы (или `gomock`/`mockery`)
- Интеграционные тесты с `testcontainers-go`
- `t.Parallel()` где можно
- `httptest.NewRecorder` для тестирования хэндлеров

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

