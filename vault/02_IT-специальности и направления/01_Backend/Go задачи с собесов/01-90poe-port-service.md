---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/90poe - Компания: 90poe - Уровень: Middle - Время на выполнение: 10-14 часов - Ключевые навыки: DDD, Domain Service, микросервисы, PostgreSQL - Инфраструктура..."
type: concept
---

# Port Domain Service

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/90poe
- **Компания:** 90poe
- **Уровень:** Middle
- **Время на выполнение:** 10-14 часов
- **Ключевые навыки:** DDD, Domain Service, микросервисы, PostgreSQL
- **Инфраструктура:** PostgreSQL, Docker

## ТЗ

### Описание
Создать domain service для управления портами морских судов. Сервис реализует бизнес-логику работы с портами, их характеристиками и маршрутами. Реализация согласно принципам Domain-Driven Design.

### Функциональные требования

**Предметная область:**
- **Port (Порт):** Уникальная сущность с координатами, названием, кодом (UN/LOCODE)
- **Vessel (Судно):** Заходит в порты, имеет маршрут
- **Route (Маршрут):** Последовательность посещений портов

**Эндпоинты:**

| Метод | Endpoint | Описание |
|-------|----------|----------|
| POST | /ports | Создать порт |
| GET | /ports/{id} | Получить порт |
| GET | /ports | Список портов с фильтрами |
| PUT | /ports/{id} | Обновить порт |
| DELETE | /ports/{id} | Удалить порт |
| POST | /vessels/{id}/routes | Добавить маршрут судна |
| GET | /vessels/{id}/routes | Получить маршрут судна |
| GET | /ports/{id}/vessels | Суда в порту |

**Структуры данных:**

```go
type Port struct {
    ID       string    // UUID
    Name     string    // Название порта
    Code     string    // UN/LOCODE (например, "NLRTM" для Rotterdam)
    City     string
    Country  string
    Coordinates
    CreatedAt time.Time
}

type Coordinates struct {
    Latitude  float64   // -90 to 90
    Longitude float64   // -180 to 180
}

type VesselRoute struct {
    VesselID    string
    PortID      string
    ArrivalTime time.Time
    DepartureTime *time.Time // nil если ещё в порту
    Order       int         // порядок в маршруте
}
```

**Бизнес-правила:**
1. UN/LOCODE уникален глобально
2. Координаты валидны (lat: -90..90, lon: -180..180)
3. Судно не может быть в двух портах одновременно
4. Маршрут упорядочен по времени
5. Порт можно "закрыть" (soft delete) если нет активных судов

### Нефункциональные требования

1. **Архитектура:**
   - Domain-Driven Design (DDD)
   - Clean Architecture / Ports and Adapters
   - Domain слой изолирован от инфраструктуры

2. **Структура:**
   ```
   internal/
   ├── domain/          # Entities, Value Objects, Domain Services
   ├── application/     # Use Cases, DTOs
   ├── infrastructure/  # DB, HTTP, Config
   └── interfaces/      # HTTP handlers, middleware
   ```

3. **База данных:**
   - PostgreSQL
   - Миграции
   - Индексы по code, coordinates

### API Спецификация

**POST /ports:**
```json
{
    "name": "Rotterdam",
    "code": "NLRTM",
    "city": "Rotterdam",
    "country": "Netherlands",
    "coordinates": {
        "latitude": 51.9244,
        "longitude": 4.4777
    }
}
```

**GET /ports?lat=51.9&lon=4.4&radius=50:**
```json
{
    "ports": [
        {
            "id": "uuid",
            "name": "Rotterdam",
            "code": "NLRTM",
            "distance_km": 5.2
        }
    ]
}
```

### Ограничения
- Только REST API (без gRPC)
- Нет аутентификации
- Нет real-time обновлений

### Бонусные задания
- [ ] Поиск портов в радиусе (geospatial query)
- [ ] ETA расчёт между портами
- [ ] История изменений (audit log)
- [ ] Импорт портов из JSON (UN/LOCODE dataset)

## Критерии оценки

- [ ] DDD структура (domain, application, infrastructure)
- [ ] Корректные entities и value objects
- [ ] Domain service с бизнес-логикой
- [ ] Repository pattern
- [ ] Валидация координат и кодов
- [ ] PostgreSQL с индексами
- [ ] Миграции
- [ ] Тесты (unit + integration)
- [ ] Docker

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: portdomain
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DATABASE_URL=postgres://user:password@postgres:5432/portdomain
    depends_on:
      - postgres

volumes:
  postgres_data:
```

### SQL-схема
```sql
CREATE TABLE ports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(10) UNIQUE NOT NULL,
    city VARCHAR(255),
    country VARCHAR(255),
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE vessel_routes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vessel_id VARCHAR(100) NOT NULL,
    port_id UUID REFERENCES ports(id),
    arrival_time TIMESTAMP NOT NULL,
    departure_time TIMESTAMP,
    route_order INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_ports_code ON ports(code);
CREATE INDEX idx_ports_coordinates ON ports USING GIST (
    point(longitude, latitude)
);
CREATE INDEX idx_routes_vessel ON vessel_routes(vessel_id);
CREATE INDEX idx_routes_port ON vessel_routes(port_id);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Anemic domain model** — все логика в сервисах, entities пустые
2. **Нет value objects** — используют примитивы везде
3. **Смешение слоёв** — domain зависит от infrastructure
4. **Нет валидации** — невалидные координаты в БД
5. **Нет индексов** — медленный поиск

### Ключевые паттерны

**Entity с бизнес-правилами:**
```go
type Port struct {
    id          PortID
    name        string
    code        UNLOCODE
    coordinates Coordinates
}

func (p *Port) UpdateCoordinates(c Coordinates) error {
    if !c.IsValid() {
        return ErrInvalidCoordinates
    }
    p.coordinates = c
    return nil
}
```

**Value Object:**
```go
type Coordinates struct {
    Latitude  float64
    Longitude float64
}

func (c Coordinates) IsValid() bool {
    return c.Latitude >= -90 && c.Latitude <= 90 &&
           c.Longitude >= -180 && c.Longitude <= 180
}

func (c Coordinates) DistanceTo(other Coordinates) float64 {
    // Haversine formula
}
```

**Domain Service:**
```go
type RouteService struct {
    portRepo   PortRepository
    routeRepo  RouteRepository
}

func (s *RouteService) AddPortToRoute(vesselID string, portID string, eta time.Time) error {
    // Проверка: судно не в другом порту
    // Проверка: ETA после предыдущего порта
    // Сохранение
}
```

### Follow-up вопросы
1. Какие ещё bounded contexts есть в maritime logistics?
2. Как бы ты реализовал geospatial search?
3. Как интегрироваться с AIS (Automatic Identification System)?
4. Какие event-sourcing возможности здесь применимы?

### Уровневые подсказки
- 🟢 **Направление:** Выдели entities (Port, Vessel), value objects (Coordinates, UNLOCODE) и domain service для маршрутов
- 🟡 **Конкретика:** Создай методы-конструкторы с валидацией для value objects
- 🔴 **Фрагмент:** ```go
func NewCoordinates(lat, lon float64) (Coordinates, error) {
    if lat < -90 || lat > 90 || lon < -180 || lon > 180 {
        return Coordinates{}, ErrInvalidCoordinates
    }
    return Coordinates{Latitude: lat, Longitude: lon}, nil
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L3 — Структуры, методы, интерфейсы
- L6 — HTTP API
- L7 — Clean architecture, структура проекта
- L8 — PostgreSQL, Docker

### Новые навыки
- Domain-Driven Design (DDD)
- Entities vs Value Objects
- Domain Services
- Repository pattern в контексте DDD

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

