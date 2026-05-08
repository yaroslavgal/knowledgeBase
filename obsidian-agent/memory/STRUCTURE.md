---
name: vault-structure
description: Согласованная модель folders, MOC, parent и навигации для Obsidian vault.
type: structure
status: draft
---

# Structure

Этот файл заполняется во время `onboard-vault` и уточняется через `organize-vault`.

## Верхнеуровневые области

Не создавай области ради симметрии. Каждая область должна помогать реальной навигации.

| area | job | folder? | MOC? | notes |
|---|---|---:|---:|---|

## MOC Model

MOC — это карта, а не папочный список.

Хороший MOC:

- даёт маршрут в область;
- группирует заметки по задаче, роли, этапу или вопросу;
- показывает важные заметки с коротким контекстом;
- остаётся читаемым.

Плохой MOC:

- зеркалит папку alphabetically;
- существует для 2-3 заметок без навигационной пользы;
- пытается выразить все связи в одном месте.

## Parent Policy

`parent` указывает на главный MOC или map заметки.

## Link Policy

Inline link first when natural. Related links need comments.

## Открытые структурные вопросы

-
