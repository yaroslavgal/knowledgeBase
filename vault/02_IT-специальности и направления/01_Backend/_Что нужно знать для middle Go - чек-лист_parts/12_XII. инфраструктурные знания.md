---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]]"
type: concept
description: "ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)"
---

## XII. ИНФРАСТРУКТУРНЫЕ ЗНАНИЯ

- Docker — пишешь Dockerfile с multi-stage build, scratch/distroless final image
- Понимаешь docker-compose
- Базовые знания Kubernetes (Pod, Service, Deployment, ConfigMap, Secret)
- CI/CD пайплайн (GitHub Actions / GitLab CI)
- Знаешь, что такое graceful shutdown:
  ```go
  ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
  defer cancel()
  
  go func() { server.ListenAndServe() }()
  <-ctx.Done()
  
  shutdownCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
  server.Shutdown(shutdownCtx)
  ```

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Чек-лист: что отличает middle Go от junior]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

