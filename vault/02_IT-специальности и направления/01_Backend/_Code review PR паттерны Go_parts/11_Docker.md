---
parent: "[[20_IT-специальности]]"
part_of: "[[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]]"
type: concept
description: "FROM golang:1.21"
---

## Docker

### ❌ Антипаттерн

```dockerfile
FROM golang:1.21
COPY . .
RUN go build -o app
CMD ["./app"]
# ← образ ~1GB, нет multi-stage, root user
```

### ✅ Multi-stage build

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/api

# Runtime stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates && \
    adduser -D -u 1000 appuser
WORKDIR /app
COPY --from=builder /build/app .
USER appuser
EXPOSE 8080
CMD ["./app"]
```

**Преимущества:** ~20MB вместо 1GB, без root, includes security updates.

---

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go|Code review и PR-паттерны для Go (типичные замечания)]] — документ-оглавление, частью которого является эта заметка.
<!-- auto-links:end -->

