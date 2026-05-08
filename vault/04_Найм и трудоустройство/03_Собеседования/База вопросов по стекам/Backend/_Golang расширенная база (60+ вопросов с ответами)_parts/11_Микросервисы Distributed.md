---
parent: "[[40_Найм и трудоустройство]]"
part_of: "[[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]]"
type: concept
description: "Таблица `outbox(id, type, payload, processed_at)`. Запись в outbox в той же tx, что и основные данные. Worker polling + `FOR UPDATE SKIP LOCKED + LIMIT N` → обработка → `UPDATE processed_at`. Retry с exponential backoff."
---

## Микросервисы / Distributed

### 73. Outbox pattern `M`
Таблица `outbox(id, type, payload, processed_at)`. Запись в outbox в той же tx, что и основные данные. Worker polling + `FOR UPDATE SKIP LOCKED + LIMIT N` → обработка → `UPDATE processed_at`. Retry с exponential backoff.

### 74. Saga pattern `M`
Распределённый бизнес-процесс с компенсациями. Если шаг 3 из 5 упал — откатываем 1 и 2 через компенсирующие действия. Альтернатива 2PC.

### 75. Graceful shutdown `S`
```go
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
<-sigChan
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
server.Shutdown(ctx)
```
Без graceful: при SIGTERM (Kubernetes) — обрыв соединений, текущие запросы прерываются, corrupted data.

⚡ Буфер `sigChan` размером 1 нужен — без него возможна (редкая) потеря сигнала.

### 76. Circuit Breaker `M`
3 состояния: Closed (норма) → Open (после N ошибок, отклоняем все запросы) → Half-Open (через timeout, пропускаем один — успех? → Closed). Защищает от каскадных падений.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[40_Найм и трудоустройство]] — родительская карта области.
- [[04_Найм и трудоустройство/03_Собеседования/База вопросов по стекам/Backend/_Golang расширенная база (60+ вопросов с ответами)|Golang расширенная база — 60+ вопросов с ответами]] — документ-оглавление, частью которого является эта заметка.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/12_Graceful shutdown|Graceful shutdown]] — детальный практический разбор того же sigChan + server.Shutdown(ctx).
<!-- auto-links:end -->

