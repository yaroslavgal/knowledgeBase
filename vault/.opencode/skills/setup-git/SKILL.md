---
name: setup-git
description: "Помочь владельцу настроить Git/GitHub backup для Obsidian vault без скрытых изменений и без лишней конфигурации."
---

# setup-git

Помоги владельцу настроить version history и backup vault через Git/GitHub.

Это не обязательная часть структуры заметок. Это защитный слой: история изменений, откат, синхронизация с GitHub.

## Принципы

- Ничего не меняй без явного approval.
- Не редактируй `.git/` вручную.
- Не добавляй `.gitignore` "на всякий случай".
- Не делай commit, push или remote setup без отдельного approval.
- Не включай secrets, tokens, private keys, large binary dumps.
- Объясняй простыми словами, что будет сделано и зачем.

## Проверить read-only

Сначала проверь:

```bash
git status --short
git remote -v
git branch --show-current
```

Если vault ещё не Git repo, `git status` может упасть. Это нормально.

Также проверь, есть ли уже root `.gitignore`, но не создавай и не редактируй его без approval.

## Вопросы владельцу

Задай 3-5 вопросов:

1. Нужен только локальный Git или ещё GitHub remote?
2. Vault приватный или его можно хранить в публичном репозитории?
3. Нужно ли коммитить настройки Obsidian и agent runtime вместе с заметками?
4. Есть ли файлы, которые точно нельзя отправлять в GitHub?
5. Как часто владелец хочет сохранять изменения: вручную, после сессии с агентом, по крупным блокам?

## Предложить план

Покажи план до изменений:

```markdown
## Git Setup Plan

- current state:
- target state:
- files to track:
- files to ignore, if any:
- commands I will run:
- approval needed:
```

Если нужен `.gitignore`, предложи минимальный список и объясни каждую строку.

## Типовой минимальный сценарий

После отдельного approval на `git init`, `git add` и `git commit`:

```bash
git init
git status --short
git add .
git commit -m "Initial vault setup"
```

GitHub remote добавляй только если владелец уже создал репозиторий или явно попросил помочь.

## После настройки

Объясни владельцу:

- как проверить изменения;
- как сделать commit;
- как понять, что всё отправлено на GitHub;
- какие файлы агент не должен трогать без approval.
