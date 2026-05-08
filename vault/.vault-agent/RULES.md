---
name: vault-rules
description: Shared operating rules for agent work inside this Obsidian vault.
type: rules
status: active
---

# Rules

These are the default rules before onboarding. The owner can refine them through `save-lesson` and `update-rules`.

## Load Order

At the start of vault work, read:

1. `AGENTS.md`
2. `.vault-agent/RULES.md`
3. `.vault-agent/PROFILE.md`
4. `.vault-agent/STRUCTURE.md`
5. `.vault-agent/STYLE.md` when writing, improving, formatting, or linking notes

Read `.vault-agent/LESSONS.md` only inside `update-rules` or when the user explicitly asks for calibration.

## Core Rules

1. No new note without search-before-write.
2. QMD is required for semantic duplicate search, cross-cluster recall, and link discovery.
3. No link without mechanism. Read both notes before proposing a meaningful link.
4. Inline links are preferred when they fit naturally into existing prose.
5. Related links need short comments. Bare related links are not allowed.
6. Folders are for durable areas. MOCs are for routes. Links are for mechanisms.
7. Do not move, rename, split, merge, archive, delete, or rebuild structure without explicit approval.
8. Do not create a complex taxonomy before the vault has evidence for it.

## Note Types

Use the smallest set that works:

| type | use |
|---|---|
| `concept` | One durable idea, model, principle, mechanism. |
| `source` | Article, book, video, podcast, transcript, or other consumed source. |
| `reference` | Stable reference, overview, person, product, resource, or external thing. |
| `map` | MOC / navigation entry point. |
| `project` | Active work context with decisions and next actions. |
| `capture` | Raw input not yet shaped. |

## Frontmatter

Default working note:

```yaml
---
date: YYYY-MM-DD
type: concept | source | reference | map | project | capture
description: "Dense retrieval sentence for humans and agents."
parent: "[[Primary Map]]"
---
```

Rules:

- `description` is for retrieval and disambiguation, not marketing copy.
- `parent` is the primary navigation spine.
- top-level maps and raw captures may omit `parent`;
- aliases are useful for alternate names, acronyms, RU/EN variants, and likely unlinked mentions;
- tags are not the primary navigation layer.

## Search Order

Use cheap search before semantic search:

1. filename/title;
2. aliases;
3. exact text with `rg` or Obsidian search;
4. `description`;
5. `parent`, MOC, backlinks;
6. QMD `lex`;
7. QMD `vec` / `hyde`.

## Approval Gates

Safe:

- read/search;
- local typo/format fixes;
- propose structure or links;
- improve `description` during requested note improvement.

Review first:

- add/remove links;
- edit related-link comments;
- update MOC sections;
- change style rules.

Explicit approval:

- move, rename, split, merge, archive, delete;
- change `parent` or `type`;
- mass link cleanup;
- broad folder/MOC rebuild;
- promote lessons into permanent rules.
