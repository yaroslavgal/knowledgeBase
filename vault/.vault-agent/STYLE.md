---
name: vault-style
description: Стиль для контент-базы YouTube/IT-карьера. Русский, разговорный, мысль через механизм. Нейтральный старт — реальные правки приходят через LESSONS.md.
type: style
status: active
---

# Style

Базовый стиль выставлен под тематику владельца (контент про IT-карьеру на русском, YouTube). Реальные предпочтения уточняются через `LESSONS.md`.

## Voice

- Русский. Технические термины по-английски, если это реальные поисковые слова (Backend, Frontend, System Design, OOP).
- Простой язык, разговорный, без академичности.
- Механизм важнее лозунга — каждый тезис должен объяснять "почему так".
- Не использовать AI-prose, "ultimate guide", "в этой статье вы узнаете".
- В `60_Возражения_и_истории/` сохранять авторскую интонацию автобиографических историй.

## Visual Flow

Use visual structure only when it helps reading or retrieval.

| form | use |
|---|---|
| Short sections | note has several jobs or arguments |
| Table | stable comparison across 2+ dimensions |
| Mermaid | lifecycle, feedback loop, architecture, dependency, branching process |
| Code fence | checklist, simple ladder, prompt skeleton, folder tree |
| TL;DR | useful for concept/source/reference notes when it says the actual mechanism |

Do not add diagrams just to make a note look complete.

## Links

Inline first when the target term already appears naturally.

Related links are for whole-note relations:

```markdown
## Related

- [[Target note]] — why this relation matters.
```

Each related link needs a mechanism comment.

## Diagrams

Use `visual-diagram` before adding a non-trivial diagram.

Mermaid defaults:

- short node labels;
- one reading direction;
- explicit return edge for cycles;
- no custom styling unless visually checked.
