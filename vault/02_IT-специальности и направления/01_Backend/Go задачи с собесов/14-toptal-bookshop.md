---
parent: "[[20_IT-специальности]]"
description: "- Источник: https://github.com/evt/home-assignments/tree/main/toptal - Компания: Toptal - Уровень: Middle - Время на выполнение: 12-18 часов - Ключевые навыки: REST API, PostgreSQL, авторизация, корзина, race conditions..."
type: concept
---

# Book Shop API

## Метаданные
- **Источник:** https://github.com/evt/home-assignments/tree/main/toptal
- **Компания:** Toptal
- **Уровень:** Middle
- **Время на выполнение:** 12-18 часов
- **Ключевые навыки:** REST API, PostgreSQL, авторизация, корзина, race conditions
- **Инфраструктура:** PostgreSQL, Docker

## ТЗ

### Описание
Создать REST API для интернет-магазина книг. Поддержка анонимных и авторизованных пользователей, корзина, покупка, админ-панель. Обработка race conditions при покупке.

### Функциональные требования

**Уровни доступа:**

| Роль | Доступ |
|------|--------|
| Anonymous | Просмотр книг, фильтрация по категориям |
| Authenticated | + Корзина, покупка |
| Admin | + CRUD категорий и книг |

**Эндпоинты:**

**Публичные (Anonymous):**
| Метод | URL | Описание |
|-------|-----|----------|
| GET | `/books` | Список книг (пагинация) |
| GET | `/books?category=fiction&category=scifi` | Фильтр по категориям (multi-select) |
| GET | `/books/:id` | Детали книги |
| POST | `/register` | Регистрация (email, password) |
| POST | `/login` | Авторизация |

**Авторизованные:**
| Метод | URL | Описание |
|-------|-----|----------|
| POST | `/cart` | Добавить книгу в корзину |
| GET | `/cart` | Просмотр корзины |
| POST | `/checkout` | Купить книги из корзины |

**Админ:**
| Метод | URL | Описание |
|-------|-----|----------|
| POST/PUT/DELETE | `/admin/categories` | CRUD категорий |
| POST/PUT/DELETE | `/admin/books` | CRUD книг |

**Бизнес-правила:**
1. Книга имеет: title, year, author, price (USD), category, stock (количество)
2. Категории плоские (не nested)
3. Книга может быть только в одной категории
4. Проданные книги (stock=0) не отображаются в списке
5. В корзине можно иметь только 1 копию каждой книги (без quantity)
6. Cart expiring: книга возвращается в stock через 30 минут если не куплена
7. Race condition: два пользователя не могут купить последнюю книгу одновременно

### Нефункциональные требования

1. **API:** RESTful или GraphQL
2. **База данных:** PostgreSQL
3. **Масштабируемость:** 10,000+ книг в каталоге
4. **Race conditions:** `SELECT ... FOR UPDATE` или аналог
5. **Валидация:** входные данные проверяются
6. **Безопасность:** пароли хешированы, JWT токены

### Ограничения
- Админы создаются только через БД (нет UI регистрации админа)
- Оплата фейковая (без интеграции с платёжками)
- Stock нельзя редактировать после создания книги

### Бонусные задания
- [ ] Email уведомление при покупке
- [ ] Поиск книг по названию/автору
- [ ] Сортировка (по цене, году)
- [ ] История заказов пользователя
- [ ] Промо-коды

## Критерии оценки

- [ ] Все CRUD операции работают
- [ ] Авторизация и роли
- [ ] Корзина с expiring (30 мин)
- [ ] Race condition handling
- [ ] Тесты (unit + integration)
- [ ] Docker и docker-compose
- [ ] Postman коллекция или curl примеры
- [ ] README с инструкциями

## Инфраструктура

### Docker Compose
```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: bookshop
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
  
  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - postgres
```

### SQL-схема
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year INTEGER,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0,
    category_id INTEGER REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    book_id INTEGER REFERENCES books(id),
    expires_at TIMESTAMP NOT NULL,
    UNIQUE(user_id, book_id)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    total DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    book_id INTEGER REFERENCES books(id),
    price DECIMAL(10,2)  -- snapshot price at purchase
);
```

## Подсказки для ментора

### Типичные ошибки учеников
1. **Нет транзакций** — race condition на checkout
2. **N+1 queries** — при загрузке корзины с книгами
3. **Plain text passwords** — не хешируют
4. **Нет индексов** — медленный поиск
5. **Cart не expire** — книги "зависают" в корзинах

### Ключевые паттерны
- **Repository pattern** — отделить логику от БД
- **Unit of Work** — транзакции
- **JWT middleware** — авторизация
- **Background job** — cleanup expired cart items

### Обработка race condition
```sql
-- Вариант 1: SELECT FOR UPDATE
BEGIN;
SELECT stock FROM books WHERE id = $1 FOR UPDATE;
-- check stock > 0
UPDATE books SET stock = stock - 1 WHERE id = $1;
COMMIT;

-- Вариант 2: Optimistic locking
UPDATE books SET stock = stock - 1 
WHERE id = $1 AND stock > 0;
-- check rows affected
```

### Follow-up вопросы
1. Как бы ты добавил real оплату (Stripe)?
2. Что если пользователь добавил в корзину, цена изменилась, а потом купил?
3. Как масштабировать до 1M книг?
4. Как сделать корзину persistent между сессиями?

### Уровневые подсказки
- 🟢 **Направление:** Начни с моделей БД, потом репозитории, потом хендлеры
- 🟡 **Конкретика:** Используй `sql.Tx` для атомарного списания stock
- 🔴 **Фрагмент:** ```go
func (r *BookRepo) DecrementStock(ctx context.Context, tx *sql.Tx, bookID int) error {
    result, err := tx.ExecContext(ctx, 
        "UPDATE books SET stock = stock - 1 WHERE id = $1 AND stock > 0",
        bookID,
    )
    if err != nil { return err }
    
    rows, _ := result.RowsAffected()
    if rows == 0 { return ErrOutOfStock }
    return nil
}
```

## Маппинг на систему обучения

### Необходимые уровни
- L6 — HTTP, JSON, middleware
- L7 — Структура проекта
- L8 — PostgreSQL, миграции, Docker

### Новые навыки
- Сложная бизнес-логика (корзина, expiring)
- Race conditions и транзакции
- RBAC (role-based access control)

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
<!-- auto-links:end -->

