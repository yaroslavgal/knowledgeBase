---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]]"
type: concept
description: "var db *sql.DB // Глобальная"
---

## III. ЗАВИСИМОСТИ И DI

### Junior
```go
var db *sql.DB                            // Глобальная
var notificationChan = make(chan ...)     // Глобальный канал
var configs = map[string]Config{...}      // Глобальный конфиг

func handle(w http.ResponseWriter, r *http.Request) {
    rows, _ := db.Query(...)              // Прямой доступ
}
```

### Middle
```go
type Server struct {
    userService    UserService
    notifService   NotificationService
    logger         *slog.Logger
}

func NewServer(us UserService, ns NotificationService, log *slog.Logger) *Server {
    return &Server{userService: us, notifService: ns, logger: log}
}

func (s *Server) HandleSendNotification(w http.ResponseWriter, r *http.Request) {
    user, err := s.userService.GetByID(r.Context(), userID)
    // ...
}
```

**Ключи:** Constructor pattern (`NewServer`), методы на структуре, инжекция через интерфейсы. Логгер — `slog.Logger` (Go 1.21+, не `log.Println`).

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

