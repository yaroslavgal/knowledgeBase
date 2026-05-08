---
name: visual-diagram
description: "Выбрать и набросать Mermaid, table, code-fence ladder, checklist или решить, что диаграмма не нужна."
---

# visual-diagram

Добавляй visual structure только когда она снижает cognitive load.

## Читать сначала

- `.vault-agent/RULES.md`
- `.vault-agent/STYLE.md`

## Decision Table

| Задача | Форма |
|---|---|
| lifecycle, feedback loop, architecture, dependency, branching process | Mermaid |
| стабильное сравнение | table |
| простая лестница, checklist, prompt skeleton, folder tree | code fence |
| one-line relation | prose |
| нет ясной visual job | no diagram |

## Mermaid Rules

- короткие labels;
- одно направление чтения;
- явный return edge для cycles;
- без custom styling by default.

## Output

- visual job;
- selected form;
- draft diagram/table/checklist, если полезно;
- почему диаграмма не нужна, если skipped.
