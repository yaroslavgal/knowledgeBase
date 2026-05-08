---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "name: CI"
---

## CI: GitHub Actions для Go-проекта

```yaml
name: CI
on: [push, pull_request]

jobs:
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with: { go-version: '1.21' }

      - name: Lint
        uses: golangci/golangci-lint-action@v3

      - name: Test
        run: go test -race -coverprofile=coverage.out ./...

      - name: Coverage
        run: go tool cover -func=coverage.out
```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
- [[04_Найм и трудоустройство/_Полный пайплайн поиска работы_parts/08_Тестовое задание — правила|Тестовое задание — правила]] — рабочий CI поверх Go-тестового — конкретный сигнал «у кандидата есть продакшен-привычки».
<!-- auto-links:end -->

