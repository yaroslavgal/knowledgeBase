---
name: update-rules
description: "Разобрать LESSONS.md и предложить owner-approved обновления RULES, STYLE, STRUCTURE или PATTERNS."
---

# update-rules

Превращай raw lessons в стабильные правила. Не применяй изменения без approval владельца.

## Читать

- `.vault-agent/LESSONS.md`
- `.vault-agent/RULES.md`
- `.vault-agent/STYLE.md`
- `.vault-agent/STRUCTURE.md`
- `.vault-agent/PATTERNS.md`

## Workflow

1. Сгруппируй pending lessons по pattern.
2. Отклони слабые или one-off entries.
3. Выбери destination: `RULES`, `STYLE`, `STRUCTURE` или `PATTERNS`.
4. Предложи small diffs.
5. Спроси approval.
6. Применяй только approved changes.

## Output

- evidence entries;
- suspected pattern;
- proposed destination;
- proposed wording;
- approval question.
