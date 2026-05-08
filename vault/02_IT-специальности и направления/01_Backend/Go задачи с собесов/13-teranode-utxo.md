---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/teranode - Компания: Teranode (Spacemesh) - Уровень: Middle - Время на выполнение: 10-14 часов - Ключевые навыки: UTXO модель, транзакции, in-memory структуры..."
type: concept
---

# UTXO Store

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/teranode
- **Компания:** Teranode (Spacemesh)
- **Уровень:** Middle
- **Время на выполнение:** 10-14 часов
- **Ключевые навыки:** UTXO модель, транзакции, in-memory структуры
- **Инфраструктура:** Нет (in-memory)

## ТЗ

### Описание
Реализовать in-memory UTXO (Unspent Transaction Output) store — упрощённый Bitcoin-стиль кошелёк. Поддержка депозитов, spending, nested транзакций с rollback/commit.

### Функциональные требования

**Интерфейс:**
```go
type UTXOStore interface {
    // Deposit добавляет unspent output
    Deposit(amount uint64) (OutputID, error)
    
    // Spend списывает средства
    Spend(amount uint64) error
    
    // Balance возвращает текущий баланс
    Balance() uint64
    
    // BalanceAt возвращает баланс после конкретного output
    BalanceAt(id OutputID) (uint64, error)
    
    // Begin начинает транзакцию
    Begin() Transaction
}

type Transaction interface {
    Deposit(amount uint64) (OutputID, error)
    Spend(amount uint64) error
    Commit() error
    Rollback() error
}
```

**Бизнес-логика:**

1. **UTXO модель:**
   - Каждый deposit создаёт output с уникальным ID
   - Balance = сумма всех unspent outputs
   - Spend выбирает outputs (FIFO или любая стратегия) и создаёт change

2. **Транзакции (nested):**
   ```
   Begin()
     Deposit(100)     // +100
     Begin()          // nested
       Deposit(50)    // +50
       Spend(30)      // -30
       Commit()       // nested committed
     Spend(120)       // -120
   Commit()           // all committed
   ```
   - Работает как nested transactions
   - Rollback откатывает только изменения в текущем уровне
   - Commit применяет изменения к родительскому уровню

3. **BalanceAt:**
   - Возвращает баланс на момент конкретного output
   - Работает для committed и uncommitted outputs

### Нефункциональные требования

1. **In-memory:** Без persistence, без БД
2. **Потокобезопасность:** Необязательно (single-threaded использование)
3. **Тесты:** Полный покрытие всех сценариев
4. **Ошибки:** Insufficient funds, invalid output ID

### Примеры использования

```go
store := NewInMemoryUTXOStore()

// Простой депозит
id1, _ := store.Deposit(1000)
fmt.Println(store.Balance()) // 1000

// Spend с change
store.Spend(300)             // spent: 1000, change: 700
fmt.Println(store.Balance()) // 700

// Nested transaction
tx := store.Begin()
tx.Deposit(500)
tx.Spend(200)
tx.Commit()

fmt.Println(store.Balance()) // 1000 (700 + 500 - 200)
```

### Ограничения
- In-memory only
- Сатоши (uint64), без плавающей точки
- Нет криптографии, просто accounting
- Nested transactions unlimited depth

### Бонусные задания
- [ ] Persistence в файл
- [ ] MVCC (Multi-Version Concurrency Control)
- [ ] Оптимизация spend (minimize outputs)
- [ ] Proof of reserves

## Критерии оценки

- [ ] Deposit создаёт outputs
- [ ] Spend работает правильно
- [ ] Change calculation
- [ ] Nested transactions
- [ ] Commit/Rollback
- [ ] BalanceAt работает
- [ ] Обработка ошибок
- [ ] Тесты (100% покрытие)

## Инфраструктура

Не требуется — pure Go code.

```
teranode/
├── utxostore.go       # интерфейс + реализация
├── utxostore_test.go  # тесты (failing)
├── go.mod
└── README.md
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Не UTXO** — пытаются считать баланс как account-based
2. **Нет nested** — плоские транзакции без уровней
3. **Утечка памяти** — не удаляют spent outputs
4. **Нет change** — spend съедает весь output
5. **Нет валидации** — отрицательные amounts

### Ключевые паттерны

**UTXO структуры:**
```go
type Output struct {
    ID     OutputID
    Amount uint64
    Spent  bool
}

type InMemoryStore struct {
    outputs  map[OutputID]*Output
    balance  uint64
    pending  []*Change // для nested transactions
}
```

**Nested транзакции:**
```go
type Tx struct {
    store    *InMemoryStore
    parent   *Tx
    changes  []Change
    committed bool
}

func (t *Tx) Commit() error {
    if t.parent != nil {
        // Переносим changes к родителю
        t.parent.changes = append(t.parent.changes, t.changes...)
    } else {
        // Применяем к store
        t.store.apply(t.changes)
    }
    t.committed = true
    return nil
}
```

**Spend стратегия (FIFO):**
```go
func (s *InMemoryStore) spendFIFO(amount uint64) error {
    var toSpend []*Output
    var collected uint64
    
    // Идём по порядку creation, собираем outputs
    for _, out := range s.unspentOutputs() {
        toSpend = append(toSpend, out)
        collected += out.Amount
        if collected >= amount {
            break
        }
    }
    
    if collected < amount {
        return ErrInsufficientFunds
    }
    
    // Помечаем spent
    for _, out := range toSpend {
        out.Spent = true
    }
    
    // Change если нужно
    if change := collected - amount; change > 0 {
        s.Deposit(change)
    }
    return nil
}
```

### Follow-up вопросы
1. Как бы ты добавил persistence?
2. Какие стратегии spend бывают (FIFO, LIFO, smallest-first)?
3. Как реализовать concurrent доступ?
4. Что такое coin selection problem?

### Уровневые подсказки
- 🟢 **Направление:** Нужна структура для хранения outputs + стек для nested transactions
- 🟡 **Конкретика:** Для nested tx храни changes в списке, применяй только при commit
- 🔴 **Фрагмент:** ```go
func (s *InMemoryStore) Begin() Transaction {
    return &nestedTx{
        store:   s,
        outputs: make(map[OutputID]*Output),
    }
}

func (t *nestedTx) Commit() error {
    for id, out := range t.outputs {
        t.store.outputs[id] = out
    }
    return nil
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L3 — Структуры, методы, интерфейсы
- L4 — (опционально) конкурентность

### Новые навыки
- UTXO модель (Bitcoin-style)
- Nested transactions
- In-memory data structures
- Transaction patterns

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

