---
parent: "[[20_IT-специальности]]"
type: concept
description: "Типовые замечания на code review Go-проектов и тестовых заданий: структура проекта, errors, тесты, Docker, CI. Чек-лист «до/после» по 16 разделам."
---

# Code review и PR-паттерны для Go (типичные замечания)

Реальные комментарии с code review реальных проектов и собеседований (включая 950К ₽ оффер с domashnim заданием на гексагональную архитектуру).

## Разделы

- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/01_Контекст Code review PR паттерны Go|Контекст]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/02_Структура проекта|Структура проекта]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/03_Модели данных — разделение|Модели данных — разделение]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/04_Receiver type — value vs pointer|Receiver type — value vs pointer]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/05_Валидация в конструкторе|Валидация в конструкторе]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/06_Application слой — нужен ли|Application слой — нужен ли]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/07_CQRS — когда оверкилл|CQRS — когда оверкилл]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/08_Тесты — пирамида|Тесты — пирамида]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/09_Errors — паттерны|Errors — паттерны]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/10_Логирование|Логирование]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/11_Docker|Docker]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/12_Graceful shutdown|Graceful shutdown]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/13_Streaming JSON parsing для больших файлов|Streaming JSON parsing для больших файлов]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/14_CI GitHub Actions для Go-проекта|CI: GitHub Actions для Go-проекта]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/15_Conventional Commits|Conventional Commits]]
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/16_Чек-лист перед PR|Чек-лист перед PR]]

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/01_Backend/_5 ошибок не берут на middle Go|5 ошибок не берут на middle Go]] — типовые замечания на ревью пересекаются с этими 5 ошибками.
- [[02_IT-специальности и направления/01_Backend/_Что такое middle Go - чек-лист|Что такое middle Go]] — навыки middle, проверяемые именно на ревью.
- [[04_Найм и трудоустройство/03_Собеседования/Технические/_Live coding Go реальные задачи|Live coding Go реальные задачи]] — те же паттерны проверяются на лайвкодинге собеса.

### Исходящие
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/01_Контекст Code review PR паттерны Go|01_Контекст__Code_review_PR_паттерны_Go]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/02_Структура проекта|02_Структура_проекта]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/03_Модели данных — разделение|03_Модели_данных_—_разделение]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/04_Receiver type — value vs pointer|04_Receiver_type_—_value_vs_pointer]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/05_Валидация в конструкторе|05_Валидация_в_конструкторе]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/06_Application слой — нужен ли|06_Application_слой_—_нужен_ли]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/07_CQRS — когда оверкилл|07_CQRS_—_когда_оверкилл]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/08_Тесты — пирамида|08_Тесты_—_пирамида]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/09_Errors — паттерны|09_Errors_—_паттерны]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/10_Логирование|10_Логирование]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/11_Docker|11_Docker]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/12_Graceful shutdown|12_Graceful_shutdown]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/13_Streaming JSON parsing для больших файлов|13_Streaming_JSON_parsing_для_больших_файлов]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/14_CI GitHub Actions для Go-проекта|14_CI_GitHub_Actions_для_Go-проекта]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/15_Conventional Commits|15_Conventional_Commits]] — связанная заметка из той же области.
- [[02_IT-специальности и направления/01_Backend/_Code review PR паттерны Go_parts/16_Чек-лист перед PR|16_Чек-лист_перед_PR]] — связанная заметка из той же области.
<!-- auto-links:end -->

