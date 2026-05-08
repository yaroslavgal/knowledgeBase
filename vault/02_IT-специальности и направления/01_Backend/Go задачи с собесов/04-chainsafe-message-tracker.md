---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/chainsafe-gossamer - Компания: ChainSafe (Gossamer — Go-реализация Polkadot) - Уровень: Middle+ - Время на выполнение: 12-16 часов - Ключевые навыки: Конкурен..."
type: concept
---

# P2P Message Tracker

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/chainsafe-gossamer
- **Компания:** ChainSafe (Gossamer — Go-реализация Polkadot)
- **Уровень:** Middle+
- **Время на выполнение:** 12-16 часов
- **Ключевые навыки:** Конкурентность, sync.Map, мьютексы, бенчмарки
- **Инфраструктура:** Нет (in-memory)

## ТЗ

### Описание
Реализовать thread-safe трекер сообщений для P2P-сети. Сообщения уникальны по ID, могут приходить от разных peers. Хранение FIFO с лимитом, дедупликация, потокобезопасность.

### Функциональные требования

**Тип сообщения:**
```go
type Message struct {
    ID     string  // уникальный ID сообщения
    PeerID string  // ID отправителя
    Data   []byte  // payload
}
```

**Интерфейс трекера:**
```go
type MessageTracker interface {
    // Add добавляет сообщение в трекер
    Add(message *Message) error
    
    // Delete удаляет сообщение
    Delete(id string) error
    
    // Message возвращает сообщение по ID
    Message(id string) (*Message, error)
    
    // Messages возвращает все сообщения в порядке добавления
    Messages() []*Message
    
    // Len возвращает количество сообщений
    Len() int
    
    // Capacity возвращает максимальный размер
    Capacity() int
}
```

**Требования:**
1. **FIFO очередь** — сообщения хранятся в порядке получения
2. **Capacity** — фиксированный лимит, при переполнении удаляется старое
3. **Дедупликация** — одинаковые ID не добавляются повторно
4. **Потокобезопасность** — методы могут вызываться из многих goroutines
5. **Одинаковый ID от разных peers** — всё равно одно сообщение (дедупликация по ID)

### Нефункциональные требования

1. **Две реализации:**
   - С `sync.Mutex` + обычная map
   - С `sync.Map`
   
2. **Бенчмарки:**
   - Сравнение производительности обеих реализаций
   - `make bench`

3. **Тесты:**
   - Unit tests с race detector
   - Coverage > 80%
   - `make test`

### Ограничения
- In-memory хранилище (без БД)
- Message.ID уникален глобально
- Ограниченный размер (capacity)

### Бонусные задания
- [ ] TTL для сообщений (автоудаление старых)
- [ ] Статистика (count by peer)
- [ ] Graceful shutdown с сохранением в файл

## Критерии оценки

- [ ] Обе реализации работают корректно
- [ ] Потокобезопасность (проходит race detector)
- [ ] FIFO порядок сохраняется
- [ ] Дедупликация работает
- [ ] Бенчмарки сравнивают производительность
- [ ] Тесты с хорошим покрытием
- [ ] Makefile с командами

## Инфраструктура

Не требуется — in-memory решение.

```
chainsafe-gossamer/
├── network/
│   ├── message.go              # тип Message
│   ├── message_tracker.go      # интерфейс
│   ├── singlemap/
│   │   ├── singlemap_tracker.go    # реализация с mutex
│   │   └── singlemap_tracker_test.go
│   └── syncmap/
│       ├── syncmap_tracker.go      # реализация с sync.Map
│       └── syncmap_tracker_test.go
├── Makefile
└── README.md
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет FIFO** — используют map, порядок не гарантирован
2. **Race condition** — читают и пишут в map без синхронизации
3. **Утечка памяти** — не удаляют старые сообщения при capacity
4. **Deadlock** — вложенные блокировки
5. **Потеря дедупликации** — проверяют ID после добавления в очередь

### Ключевые паттерны

**Реализация 1: Mutex + Map + List**
```go
type SingleMapTracker struct {
    mu       sync.RWMutex
    messages map[string]*Message  // для O(1) поиска
    order    *list.List           // для FIFO
    capacity int
}
```

**Реализация 2: sync.Map + отдельный список**
```go
type SyncMapTracker struct {
    messages sync.Map  // map[string]*list.Element
    order    *list.List
    mu       sync.Mutex // только для order
    capacity int
}
```

**Почему list а не slice:**
- Удаление из начала slice — O(n) (копирование)
- Удаление из list — O(1)

### Follow-up вопросы
1. Когда `sync.Map` выигрывает у `map+mutex`? (чтение >> записи)
2. Как бы ты масштабировал это на кластер?
3. Что если нужен поиск по PeerID?
4. Как бы ты реализовал TTL?

### Уровневые подсказки
- 🟢 **Направление:** Нужна структура для O(1) доступа по ID + структура для FIFO порядка
- 🟡 **Конкретика:** Используй `map[string]*list.Element` + `container/list` для FIFO
- 🔴 **Фрагмент:** ```go
func (t *Tracker) Add(m *Message) error {
    t.mu.Lock()
    defer t.mu.Unlock()
    
    if _, exists := t.messages[m.ID]; exists {
        return ErrDuplicate
    }
    
    if t.order.Len() >= t.capacity {
        // удалить самое старое
        oldest := t.order.Front()
        oldMsg := oldest.Value.(*Message)
        delete(t.messages, oldMsg.ID)
        t.order.Remove(oldest)
    }
    
    elem := t.order.PushBack(m)
    t.messages[m.ID] = elem
    return nil
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L4 — Горутины, каналы, mutex
- L5 — sync.Map, context

### Новые навыки
- Продвинутая конкурентность
- Выбор между mutex и sync.Map
- Бенчмаркинг и профилирование
- P2P networking concepts

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

