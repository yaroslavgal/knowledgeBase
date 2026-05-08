---
name: save-lesson
description: "Записать одну reusable correction владельца в LESSONS.md без продвижения в постоянные правила."
---

# save-lesson

Добавь одну компактную запись в `.vault-agent/LESSONS.md`.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/LESSONS.md` только для entry template и pending-entry count

Pending lessons не являются правилами.

Записывай только reusable signal:

- владелец отклонил style, structure или link pattern;
- владелец принял non-obvious pattern;
- search/QMD пропустил важное;
- повторяющееся поведение агента нужно исправить.

Пропускай опечатки и one-off wording polish.

## Entry Shape

```markdown
## YYYY-MM-DD — short lesson title
path: note or area that triggered the lesson
agent did: what happened
owner fixed: what the owner corrected or rejected
lesson candidate: reusable lesson in one sentence
domain: rules | style | structure | links | retrieval | patterns
status: pending
```

Не редактируй `RULES.md`, `STYLE.md`, `STRUCTURE.md` или `PATTERNS.md`.
