# Obsidian Agent

`obsidian-agent` — runtime для работы AI-агента внутри Obsidian vault.

Он ставит в vault инструкции, skills, read-only agents и QMD MCP config для OpenCode и Claude Code.

Он не ставит локальные dev-сервисы, caches, project tooling и чужую конфигурацию. Это минимальный setup для Obsidian vault.

```text
ObsidianVault/
  AGENTS.md
  CLAUDE.md
  .mcp.json
  .vault-agent/
    RULES.md
    STYLE.md
    PROFILE.md
    STRUCTURE.md
    PATTERNS.md
    LESSONS.md
    examples/
  .opencode/
    opencode.json
    README.md
    agents/
    commands/
    skills/
  .claude/
    settings.json
    README.md
    agents/
    skills/
```

## Зачем

С Obsidian можно работать вручную, но со временем это становится медленно.

Обычно начинают мешать не сами заметки, а обслуживание системы:

- нужно помнить, куда положить новую заметку;
- сложно понять, когда нужна папка, а когда MOC;
- сложно найти старые мысли другими словами;
- сложно увидеть связи между заметками;
- сложно поддерживать структуру, когда vault растёт;
- легко начать линковать всё подряд или, наоборот, ничего не связывать.

`obsidian-agent` снимает часть этой ручной работы:

- проводит onboarding;
- помогает выбрать стартовую структуру;
- проверяет дубликаты через обычный поиск и QMD;
- предлагает связи с объяснением механизма;
- помогает улучшать заметки и MOC;
- запоминает коррекции владельца.

## Установка

Нужно:

- Obsidian;
- OpenCode или Claude Code;
- `qmd`, потому что QMD MCP здесь обязательный.
- `python3` и `rsync` для install/validation scripts.

Полезно, но не обязательно на первом шаге:

- Git/GitHub для истории изменений и backup;
- Obsidian plugins для ручной навигации и удобства. Список ниже.

```bash
./install.sh "/path/to/ObsidianVault"
```

После установки открой OpenCode или Claude Code из корня vault.

Установщик трогает только свои runtime-файлы:

- `AGENTS.md`;
- `CLAUDE.md`;
- `.mcp.json`;
- `.vault-agent/`;
- `.opencode/`;
- `.claude/`.

Он не редактирует root `.gitignore` и не удаляет чужие dotfolders.
При повторной установке он не перезаписывает `.vault-agent/*.md` и owner examples.

Он не ставит Obsidian plugins. Агент может объяснить, что включить вручную, но сам setup остаётся минимальным.

## Obsidian plugins

`obsidian-agent` не устанавливает plugins автоматически.

Plugins — это ручной слой Obsidian. Агент только помогает выбрать, что включить.

### Базовый минимум

Включить в Obsidian Core plugins:

| Plugin | Зачем |
|---|---|
| Backlinks | Видеть входящие ссылки и unlinked mentions. |
| Outgoing Links | Видеть исходящие ссылки и потенциальные links. |
| Page Preview | Быстро смотреть заметку по hover, не открывая её. |
| Search | Обычный точный поиск по vault. |
| Quick Switcher | Быстро открывать заметки по имени. |
| Command Palette | Быстро запускать команды Obsidian и plugins. |
| File Recovery | Локальная страховка от случайных правок. |

### Полезно после onboarding

Community plugins:

| Plugin | Когда нужен |
|---|---|
| Obsidian Git | Если владелец хочет backup/version history через GitHub из интерфейса Obsidian. |
| Editing Toolbar | Если неудобно писать Markdown руками. |
| Recent Files | Если владелец часто возвращается к последним заметкам. |
| Homepage | Если нужен стартовый экран/MOC при открытии Obsidian. |
| Templater | Если появились стабильные шаблоны заметок. |
| QuickAdd | Если появились повторяемые capture/create workflows. |
| Image Zoom & Drag | Если vault активно использует Mermaid/изображения/диаграммы. |

### Не ставить по умолчанию

| Plugin | Почему не в минимуме |
|---|---|
| Auto Note Mover | Может перемещать заметки автоматически. Это конфликтует с approval-first структурой. |
| Linter | Может массово форматировать заметки. Нужен только после согласования style rules. |
| Dataview | Мощный query layer, но добавляет сложность. Лучше сначала попробовать Bases или обычные MOC. |
| Manual Sorting | Личный UI-порядок, не фундамент системы. |
| Book Search | Нужен только если vault активно ведёт книги. |
| Contribution Graph | Метрика активности, не навигация. |
| Style Settings / Minimal Theme Settings | Внешний вид, не структура знаний. |

Первый запрос:

```text
/onboard
```

Или обычным языком:

```text
Помоги настроить этот vault.
```

## Что будет на onboarding

Сначала агент проверит базовый runtime:

- QMD доступен;
- QMD MCP прописан для OpenCode/Claude Code;
- runtime-файлы на месте;
- Obsidian plugins понятны владельцу как optional ручная настройка.

Агент сначала объяснит базовые вещи:

| Термин | Что значит |
|---|---|
| Vault | Папка с Markdown-файлами. |
| MOC | Map of Content: карта входа в область. Не список всех файлов. |
| `parent` | Главная карта, к которой относится заметка. |
| `description` | Короткая строка, по которой заметку легче найти и понять. |
| Inline link | Ссылка прямо в тексте. |
| Related link | Связь в отдельном блоке с коротким комментарием. |
| QMD | Семантический поиск по vault. Нужен для похожих заметок, дубликатов и связей. |

Потом спросит:

- зачем нужен vault;
- какие материалы будут внутри;
- как владелец хочет искать заметки;
- какие области и MOC нужны на старте;
- какие связи полезны;
- какой стиль раздражает или подходит.

Результат onboarding:

- заполненный `PROFILE.md`;
- черновик `STRUCTURE.md`;
- базовый `STYLE.md`;
- черновик `PATTERNS.md`, если появились повторяемые паттерны;
- первые MOC, если они нужны;
- 3-5 следующих действий.

## Skills

Основа — skills. Команды только короткие входы.

Skill — это рабочий сценарий: создать заметку, найти связи, проверить vault, обновить правила.
Command — короткая OpenCode-команда для частого сценария. Например, `/onboard` просто запускает `onboard-vault`.
Agent — read-only помощник для проверки или поиска кандидатов. Он не редактирует заметки, а возвращает evidence и предложения.

| Skill | Для чего |
|---|---|
| `onboard-vault` | Первый запуск и проектирование vault. |
| `create-note` | Создать заметку после поиска дубликатов. |
| `improve-note` | Улучшить существующую заметку без смены её смысла. |
| `organize-vault` | Папки, MOC, `parent`, структура. |
| `find-links` | Найти связи и дубликаты через поиск + QMD. |
| `visual-diagram` | Mermaid, таблицы, checklist, простые схемы. |
| `setup-git` | Настроить Git/GitHub backup после отдельного подтверждения. |
| `check-vault` | Read-only проверка vault или кластера. |
| `save-lesson` | Записать коррекцию владельца. |
| `update-rules` | Превратить накопленные уроки в правила. |

OpenCode commands:

```text
/onboard
/check
/learn
/reflect
```

В Claude Code используй те же сценарии обычным языком: “проведи onboarding”, “проверь vault”, “запиши lesson”, “обнови rules”.

Read-only agents:

| Agent | Для чего |
|---|---|
| `connection-scout` | Найти кандидаты связей и проверить механизм. |
| `structure-scout` | Посмотреть folders, MOC, `parent` и topology. |
| `reviewer` | Свежим взглядом проверить заметку или proposal. |
| `vault-scanner` | Проверить vault или область на schema, links, duplicates и QMD gaps. |

## Структура

Базовая модель:

```text
folder      -> устойчивая область
MOC         -> маршрут по области
parent      -> главный навигационный указатель
link        -> смысловая связь
alias       -> альтернативное имя
description -> строка для поиска
QMD         -> поиск по смыслу
```

Папки не должны заменять мышление. MOC не должен быть копией папки. Ссылки не должны ставиться просто потому, что две заметки "про одно и то же".

## Ссылки

Правило:

```text
нет ссылки без механизма
```

Приоритет:

1. inline link, если ссылка естественно встраивается в текст;
2. related link с комментарием, если связь относится ко всей заметке;
3. никакой ссылки, если есть только общая тема.

## Калибровка

Сначала runtime работает по базовым правилам.

Потом владелец поправляет агента:

```text
Пиши короче.
Не делай такие широкие MOC.
Не связывай заметки только по похожей теме.
```

Агент записывает это в `LESSONS.md`.

`/reflect` предлагает обновить:

- `RULES.md`;
- `STYLE.md`;
- `STRUCTURE.md`;
- `PATTERNS.md`.

## Проверка

```bash
runtime/scripts/validate-runtime.sh
runtime/scripts/validate-vault-runtime.sh "/path/to/ObsidianVault"
```
