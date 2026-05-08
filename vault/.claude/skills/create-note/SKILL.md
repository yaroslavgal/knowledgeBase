---
name: create-note
description: "Создать долговечную Obsidian-заметку после поиска дубликатов, QMD semantic check, выбора type, parent и места."
---

# create-note

Создавай заметку только если в vault ещё нет той же идеи.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/PROFILE.md`
- `.vault-agent/STRUCTURE.md`
- `.vault-agent/STYLE.md`

Не читай `LESSONS.md`.

## Workflow

1. Определи identity: `concept`, `source`, `reference`, `map`, `project` или `capture`.
2. Найди title, aliases, точные фразы, `description`, parent/MOC cluster.
3. Используй QMD, если есть semantic duplicate risk, cross-language recall или неясная формулировка.
4. Реши: обновить существующую заметку, предложить merge или создать новую.
5. Запиши минимальный frontmatter: `date`, `type`, `description`, `parent`, если это не top-level map или raw capture.
6. Напиши тело в текущем стиле владельца.
7. Добавляй links только после чтения обеих сторон и механизма.

## Output

- созданный или переиспользованный path;
- evidence поиска дубликатов;
- type и parent;
- links proposed/deferred;
- вопросы по структуре.
