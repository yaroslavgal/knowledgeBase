---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/snake-validator - Компания: Не указана (X backend challenge) - Уровень: Middle - Время на выполнение: 8-12 часов - Ключевые навыки: HTTP API, валидация, state..."
type: concept
---

# Snake Game Validator

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/snake-validator
- **Компания:** Не указана (X backend challenge)
- **Уровень:** Middle
- **Время на выполнение:** 8-12 часов
- **Ключевые навыки:** HTTP API, валидация, stateless сервис
- **Инфраструктура:** Нет (stateless)

## ТЗ

### Описание
Создать stateless HTTP-сервис для валидации игры Snake. Клиенты отправляют ходы, сервис проверяет корректность и возвращает новое состояние игры. Используется для защиты от читерства в глобальном лидерборде.

### Упрощения игры

| Правило | Значение |
|---------|----------|
| Длина змейки | Всегда 1 (не растёт) |
| Фруктов на поле | Всегда 1 |
| Game over | При столкновении со стеной |
| Стартовая позиция | (0, 0) |
| Стартовая скорость | (1, 0) — движение вправо |

### Функциональные требования

**Эндпоинты:**

| Метод | URL | Query Params | Response |
|-------|-----|--------------|----------|
| GET | `/new` | `w` (width), `h` (height) | JSON state с случайной позицией фрукта |

**State структура:**
```go
type State struct {
    GameID    string    // уникальный ID игры
    Width     int       // ширина поля
    Height    int       // высота поля
    Score     int       // текущий счёт
    Snake     Position  // позиция головы змейки
    Fruit     Position  // позиция фрукта
    Direction Velocity  // текущее направление
}

type Position struct {
    X int
    Y int
}

type Velocity struct {
    VelX int  // -1, 0, 1
    VelY int  // -1, 0, 1
}
```

| Метод | URL | Body | Response |
|-------|-----|------|----------|
| POST | `/validate` | State + ticks | Новый state или ошибка |

**Body POST /validate:**
```json
{
    "game_id": "uuid",
    "width": 10,
    "height": 10,
    "score": 0,
    "snake": {"x": 0, "y": 0},
    "fruit": {"x": 5, "y": 5},
    "direction": {"vel_x": 1, "vel_y": 0},
    "ticks": [
        {"vel_x": 1, "vel_y": 0},
        {"vel_x": 1, "vel_y": 0},
        {"vel_x": 0, "vel_y": 1}
    ]
}
```

**Валидация ходов:**
1. Каждый tick — смещение змейки по velocity
2. Новая позиция = старая + velocity
3. Если змейка вышла за границы (x<0 || x>=width || y<0 || y>=height) → 400 Invalid
4. Если змейка достигла фрукта → score++, новый фрукт (рандом), 200 OK
5. Если не достигла → текущий state, 200 OK

**Требования:**
- Сервис **stateless** — не хранит состояние между запросами
- Генерация fruit — случайная, но не на змейке
- Thorough testing — это ключевое требование

### Ограничения
- Stateless (нет БД, сессий)
- Змейка не растёт (упрощение)
- Один фрукт на поле

### Бонусные задания
- [ ] Валидация game_id (что он существовал)
- [ ] Replay protection (нельзя отправить старый state)
- [ ] Rate limiting

## Критерии оценки

- [ ] Корректность валидации движения
- [ ] Обработка столкновений со стеной
- [ ] Правильный подсчёт очков
- [ ] Генерация нового фрукта
- [ ] Stateless архитектура
- [ ] Thorough testing (граничные случаи)
- [ ] Обработка ошибок (400 для invalid)

## Инфраструктура

Не требуется — stateless сервис.

```
snake-validator/
├── main.go           # HTTP сервер
├── game.go           # логика игры
├── validator.go      # валидация
├── types.go          # структуры
├── game_test.go      # тесты
└── README.md
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Stateful сервис** — пытаются хранить игры в памяти
2. **Неправильная валидация границ** — off-by-one ошибки
3. **Генерация фрукта** — может появиться на змейке
4. **Не проверяют velocity** — можно телепортироваться
5. **Нет тестов на граничные случаи** — что при (0,0) и velocity (-1,0)?

### Ключевые паттерны
- **State machine** — переходы между состояниями
- **Pure functions** — валидация без side effects
- **Table-driven tests** — много сценариев

### Алгоритм валидации
```go
func Validate(state State, ticks []Velocity) (State, error) {
    snake := state.Snake
    
    for _, vel := range ticks {
        // применяем velocity
        snake.X += vel.VelX
        snake.Y += vel.VelY
        
        // проверка границ
        if snake.X < 0 || snake.X >= state.Width || 
           snake.Y < 0 || snake.Y >= state.Height {
            return State{}, ErrOutOfBounds
        }
    }
    
    // проверка съедания фрукта
    if snake == state.Fruit {
        state.Score++
        state.Fruit = randomFruit(state.Width, state.Height, snake)
    }
    
    state.Snake = snake
    return state, nil
}
```

### Follow-up вопросы
1. Как бы ты сделал stateful версию с лидербордом?
2. Как защититься от читерства (подделка state)?
3. Как бы ты добавил растущую змейку?
4. Почему stateless лучше для этого use case?

### Уровневые подсказки
- 🟢 **Направление:** Для каждого tick применяй velocity к позиции и проверяй границы
- 🟡 **Конкретика:** Создай функцию `ApplyTick(state State, tick Velocity) (State, error)`
- 🔴 **Фрагмент:** ```go
func (s *Position) Move(v Velocity) {
    s.X += v.VelX
    s.Y += v.VelY
}

func IsValid(s State) bool {
    return s.Snake.X >= 0 && s.Snake.X < s.Width &&
           s.Snake.Y >= 0 && s.Snake.Y < s.Height
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L3 — Структуры, методы
- L6 — HTTP сервер, JSON

### Новые навыки
- Stateless архитектура
- Валидация бизнес-логики
- Game logic в коде

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

