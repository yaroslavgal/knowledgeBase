---
name: improve-note
description: "Улучшить текст, структуру, description, visual flow и локальное оформление одной заметки без смены её identity."
---

# improve-note

Улучшай заметку, не меняя молча то, чем она является.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/PROFILE.md`
- `.vault-agent/STRUCTURE.md`
- `.vault-agent/STYLE.md`

Не читай `LESSONS.md`.

## Можно

- прояснить текст;
- улучшить `description`;
- улучшить порядок секций;
- добавить TL;DR, если он реально помогает;
- использовать `visual-diagram` для реальной visual job;
- предложить link/structure follow-ups.

## Нельзя молча

- добавлять или удалять links;
- менять `parent`, `type`, title или path;
- split, merge, archive, delete;
- превращать одну коррекцию владельца в постоянное правило.

## Workflow

1. Прочитай заметку и найди настоящую проблему.
2. Определи тип работы: content, structure, links или style calibration.
3. Если изменение структурное, переключись на `organize-vault`.
4. Если нужны link changes, предложи `find-links`.
5. Редактируй только внутри approved note boundary.
6. Проверь frontmatter, links и visual syntax.

## Output

- changed path;
- edit type;
- что намеренно не менялось;
- follow-up вопросы или proposals.
