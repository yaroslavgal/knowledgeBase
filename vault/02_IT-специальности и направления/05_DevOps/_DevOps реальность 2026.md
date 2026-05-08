---
parent: "[[20_IT-специальности]]"
description: "DevOps/SRE 2026: дефицит специалистов, что учить, реалистичный путь входа. AI не автоматизирует инфраструктуру."
type: concept
---

# DevOps / SRE — реальность 2026

DevOps — один из немногих сегментов с растущим дефицитом. AI не автоматизирует инфраструктуру в ближайшие годы: слишком много контекста, безопасности, легаси, неочевидных взаимодействий между сервисами.

## Реальность 2026

**Вакансии:** дефицит нарастает. Особенно остро не хватает middle+ с реальным опытом продакшен-инцидентов.

**Конкуренция:** ниже, чем в backend/frontend. На middle-вакансию ~10–25 резюме (для сравнения: Python — 100–200).

**Зарплаты в Москве 2026:**
- Стажёр: 0–50к (символическая).
- Junior (вакансий мало): 100–180к.
- Middle: 250–450к.
- Senior: 450–800к.
- Lead/Staff SRE: 600–1000к+.

**Куда нанимают:**
- Bigtech (Яндекс, Сбер, Тинькофф, Авито, ВК) — крупные SRE-команды.
- Финтех — банки строят свои PaaS, нужны люди.
- E-commerce — нагрузки растут, надо.
- Облачные провайдеры (Yandex Cloud, VK Cloud, Selectel) — внутренняя разработка.

## DevOps vs SRE vs Platform Engineer

Названия часто пересекаются, но смысл разный:

**DevOps Engineer** — настраивает CI/CD, деплой, инфраструктуру для разработчиков. Чаще всего в средних компаниях. Ходит между разработкой и операциями.

**SRE (Site Reliability Engineer)** — инженер с уклоном в надёжность. Отвечает за SLA, monitoring, incident response, capacity planning. Чаще в bigtech и компаниях с высокой нагрузкой.

**Platform Engineer** — строит внутреннюю платформу для разработчиков (внутренний PaaS, dev portal, self-service инструменты). Тренд последних 2–3 лет в bigtech.

**На собесах требования похожи**, разница больше в фокусе: DevOps больше про скрипты и pipelines, SRE — про мониторинг и инциденты, Platform — про продуктовое мышление о внутренних tooling.

## Стек 2026 — что должно быть в наборе

### Базовый минимум

**Linux:**
- Bash, основные команды (find, grep, sed, awk, ps, top, netstat, ss).
- Системные сервисы (systemd), журналы (journalctl).
- Файловая система, права, владение.
- Сети: iptables/nftables basics, маршрутизация, DNS.

**Сети:**
- HTTP/HTTPS, TLS basics (как работает handshake, certificate chain).
- TCP/IP, что такое порты, sockets.
- DNS: A, CNAME, MX, NS — как резолвится домен.
- CDN basics.

**Программирование:**
- Один скриптовый язык: Python или Go (Go всё чаще в DevOps-инструментах).
- Bash — must.
- Базовый Git.

### Контейнеры и оркестрация

**Docker:**
- Dockerfile, multi-stage builds.
- docker-compose.
- Networks, volumes.
- Best practices (root vs non-root, layer caching, security scanning).

**Kubernetes:**
- Pod, Deployment, Service, Ingress.
- ConfigMap, Secret.
- Namespace, RBAC.
- StatefulSet, DaemonSet.
- Helm — пакетный менеджер.
- Helm vs Kustomize.
- Базовое понимание операторов (CRD).

### CI/CD

- GitLab CI, GitHub Actions, или Jenkins (один из).
- Pipeline as code, multi-stage, parallel jobs.
- Артефакты, кэширование.
- GitOps: ArgoCD или Flux.

### Облака

Один основной (выбираешь по компании-цели):
- **AWS** — крупнейший, нужен для международного рынка.
- **Yandex Cloud / VK Cloud** — для РФ-найма.
- **GCP** — нишево, но удобно интегрируется с k8s.

Базовые сервисы: VM (EC2), Object storage (S3), Load balancer, RDS-аналог, Managed k8s (EKS/GKE/AKS), VPC.

### Инфраструктура как код

- **Terraform** — стандарт. Знай modules, state, providers.
- Ansible — для конфигурации серверов (post-Terraform).
- Pulumi — альтернатива Terraform на полноценных языках (растущий тренд).

### Мониторинг и логирование

- **Prometheus + Grafana** — стандарт.
- ELK/EFK stack (Elasticsearch + Fluentd/Logstash + Kibana) или Loki.
- OpenTelemetry — стандарт трейсинга.
- Sentry — error tracking.
- Alertmanager — alerts из Prometheus.

### Безопасность

- Secrets management: Vault, sops, sealed-secrets.
- Image scanning: Trivy, Snyk.
- Network policies в k8s.
- Базовое понимание SBOM.

## Реалистичный путь входа в DevOps

DevOps — не первая профессия для входа в IT. Лучший путь — через **разработку → DevOps**, потому что нужно понимать, как работает приложение, чтобы его задеплоить и поддерживать.

Канон обучения — через практику с AI. Никаких лекций «впрок». AI как ментор 24/7 (намёк → декомпозиция → объяснение, **никогда готовый код**). См. [[01_Образование и обучение/06_Менторы и AI-обучение/_Философия обучения через практику|Философия обучения через практику]].

**Если уже разработчик (полгода-год бэкендера):**

**Месяц 1–2: Linux + Docker + сети**
- Установить Linux-VM (Ubuntu/Debian), жить в терминале.
- Linux Foundation course по Linux (бесплатный).
- Docker: написать Dockerfile для своего пет-проекта, поднять с docker-compose.
- Базовые сети: что такое NAT, как работает Docker network.

**Месяц 2–3: Kubernetes**
- Поднять локальный k8s (minikube / kind / k3s).
- Деплоить свой пет-проект в k8s: Pod → Deployment → Service → Ingress.
- ConfigMap, Secret, переменные окружения.
- Helm chart для проекта.

**Месяц 3–4: CI/CD + IaC**
- GitHub Actions для пет-проекта: build, test, deploy.
- Terraform: создать через него VM в одном из облаков.
- Полный цикл: код → CI → Docker image → deploy в k8s через ArgoCD.

**Месяц 4–6: продакшен-уровень**
- Развернуть полноценный стек: приложение + БД + Redis + Prometheus + Grafana + Loki.
- Настроить alerts (Slack/Telegram bot).
- Один сложный сценарий: blue-green deployment или canary.
- Disaster recovery план.

**Месяц 6–9: подготовка к собесу + поиск middle DevOps-оффера**
- Банк вопросов по DevOps (k8s internals, Linux deep, troubleshooting).
- System Design на DevOps: «как масштабировать сервис до 1М RPS», «спроектируй CI/CD для микросервисов».
- Параллельный поиск 15–20 компаний.

**Если идёшь в DevOps с нуля без опыта разработки:**
- Срок 9–12 месяцев минимум.
- Сначала Linux + Bash + Python (3 месяца).
- Один backend-фреймворк (Python + FastAPI) на уровне «могу написать REST API» (2 месяца).
- Дальше — DevOps-стек как описано выше.

## Темы, которые точно спросят на middle-собесе

### Linux и сети
- Что происходит, когда вводишь URL в браузер? (от DNS до рендеринга).
- Как работает TLS handshake.
- Что такое load average, чем отличается от CPU usage.
- Как поймать утечку памяти на сервере.
- Чем отличается soft link от hard link.

### Docker
- Что хранится в Docker image vs container.
- Как уменьшить размер image (multi-stage, distroless).
- Чем VOLUME отличается от bind mount.
- Что происходит при `docker run` под капотом (cgroups, namespaces).

### Kubernetes
- Жизненный цикл Pod (Pending → Running → Succeeded/Failed).
- Liveness vs Readiness vs Startup probe.
- Как работает Service (kube-proxy, iptables).
- Как Ingress отличается от Service LoadBalancer.
- Stateful vs Stateless workloads, когда StatefulSet.
- Resource requests/limits, OOMKilled.
- Как реализовать canary deployment.

### CI/CD
- В чём разница между CI и CD.
- Как сделать pipeline быстрым (caching, параллелизация).
- Как обезопасить secrets в pipeline.
- GitOps vs traditional deploy.

### Мониторинг и инциденты
- Four Golden Signals (latency, traffic, errors, saturation).
- SLI vs SLO vs SLA.
- Как организовать on-call.
- Как делать incident postmortem (blameless).

## AI на работе DevOps: ~10–15% автономно

В DevOps цифра ниже остальных (~10–15% автономно), потому что инфраструктура контекстная: компания, легаси, безопасность, специфика облака.

Что AI делает хорошо:
- Boilerplate Terraform-модулей.
- Базовые Helm charts.
- YAML-генерация (k8s manifests, GitHub Actions workflows).
- Скрипты для troubleshooting (поиск по логам, diff между конфигами).
- Объяснение незнакомых ошибок (k8s events, OOMKilled).

Что AI **не делает:**
- Архитектурные решения (multi-region, disaster recovery, какой облачный провайдер).
- Дебаг продакшен-инцидентов в реальном времени.
- Любая работа с production secrets.
- Capacity planning.
- Network policies в сложной топологии.

## Антипаттерны

❌ **Идти в DevOps первой профессией без разработки.** Деплоишь то, чего не понимаешь — постоянные баги, конфликты с командой.
**Решение:** хотя бы 6 месяцев на одном backend-стеке перед DevOps.

❌ **Учить только Kubernetes без Docker и Linux.** Часто видят k8s как «магию», не понимая, что под капотом контейнеры и Linux-фичи.
**Решение:** Linux → Docker → Kubernetes, в этом порядке.

❌ **Игнорировать программирование.** Современный DevOps — это IaC, скрипты, операторы. Без Python/Go/Bash не работает.
**Решение:** Python до уровня «могу написать утилиту 100–200 строк с тестами».

❌ **Учить все облака.** Распыление.
**Решение:** один облачный провайдер до middle-уровня (AWS для международного рынка, Yandex Cloud для РФ).

## Чек-лист готовности к middle-собесу DevOps

- [ ] Linux: уверенно работаю в терминале, могу troubleshoot системные проблемы.
- [ ] Docker: пишу Dockerfile, использую multi-stage, понимаю layer caching.
- [ ] Kubernetes: разворачиваю приложение через Helm, могу описать жизненный цикл Pod.
- [ ] CI/CD: написал минимум 3 pipeline (GitHub Actions / GitLab CI).
- [ ] Terraform: создал VPC + VM + БД через IaC, понимаю state.
- [ ] Prometheus + Grafana: настроил мониторинг с alerts.
- [ ] Один облачный провайдер на уровне «могу развернуть веб-приложение от 0 до прод».
- [ ] Понимаю Four Golden Signals и SLI/SLO.
- [ ] Один пет-проект — production-grade infrastructure (приложение + БД + monitoring + auto-deploy).
- [ ] Bash до уровня «могу написать deployment script на 100 строк».
- [ ] Python до уровня «могу написать утилиту с тестами».

## Связи

- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — сравнение основных стеков.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий 2026]] — сырые цифры.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — зарплаты.

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — общее сравнение стеков; DevOps — отдельный путь, не покрытый там.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий по стекам 2026]] — комплементарные цифры спроса на инфраструктурные роли.
- [[08_Будущее и тренды/01_AI и автоматизация профессий/_Заменит ли AI разработчиков|Заменит ли AI разработчиков]] — почему DevOps — одна из устойчивых к AI-замене ролей.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — DevOps-вилки в общей картине рынка.
- [[01_Образование и обучение/05_Путь обучения по специальностям/_Выбор направления|Выбор направления]] — DevOps как альтернативный путь в обучении.
<!-- auto-links:end -->
