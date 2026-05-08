---
name: check-vault
description: "Read-only audit vault или кластера: structure, schema, links, duplicates, descriptions и QMD recall gaps."
---

# check-vault

Проверяй read-only. Не редактируй заметки или memory.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/PROFILE.md`
- `.vault-agent/STRUCTURE.md`
- `.vault-agent/STYLE.md`, если проверяешь prose или visual flow

## Проверяй

- missing/weak `description`;
- missing `parent`;
- weak/missing MOCs;
- duplicate candidates;
- broken/weak links;
- related links without comments;
- folder/MOC/parent mismatch;
- области, где нужен onboarding follow-up.

Сначала exact search. QMD используй для semantic duplicates, cross-cluster recall и link-discovery candidates.

## Output

Findings по severity:

- path;
- evidence;
- proposed action;
- approval needed.
