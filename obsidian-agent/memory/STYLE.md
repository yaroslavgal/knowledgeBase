---
name: vault-style
description: Default writing, formatting, diagram, and linking style before owner calibration.
type: style
status: active
---

# Style

Start neutral. Learn the owner's real preferences through `LESSONS.md`.

## Voice

- Plain language.
- Mechanism over slogans.
- No hype, no "ultimate guide", no AI sales prose.
- Preserve the owner's language when rewriting personal notes.
- Use English technical handles when they are real search terms.

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
