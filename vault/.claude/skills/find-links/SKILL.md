---
name: find-links
description: "Найти и проверить meaningful note links по механизму. Использует QMD для semantic recall и остаётся proposal-first."
---

# find-links

Ищи связи по evidence, не по vibes.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/PROFILE.md`
- `.vault-agent/STRUCTURE.md`
- `.vault-agent/STYLE.md`

Не читай `LESSONS.md`.

## Workflow

1. Прочитай source note.
2. Ищи exact/title/alias/description/parent/MOC/backlinks.
3. Используй QMD `lex`, `vec` или `hyde` для semantic candidates.
4. Ранжируй candidates, но не считай rank доказательством.
5. Прочитай top candidate notes.
6. Сформулируй механизм связи в одно предложение.
7. Выбери placement:
   - inline, если target естественно есть в тексте;
   - related link, если связь на уровне всей заметки или нужен комментарий.
8. Предложи изменения. Редактируй только если пользователь явно попросил применить links.

## Reject

- topical similarity;
- generic hubs без local mechanism;
- QMD-only confidence;
- links без чтения обеих заметок;
- related links без comments.

## Output

```markdown
## Proposed links

- [[Target]] — mechanism: ...
  - evidence: exact / parent-MOC / backlink / QMD lex / QMD vec / QMD hyde
  - placement: inline / related
  - approval: needed

## Rejected

- [[Target]] — reason: ...
```
