---
parent: "[[20_IT-специальности]]"
description: "Mobile 2026: iOS, Android, кросс-платформа. Зарплаты, конкуренция, выбор пути."
type: concept
---

# Mobile — реальность 2026

Mobile — стабильная ниша со средней конкуренцией. Меньше вакансий, чем в backend, но и меньше кандидатов. Подходит тем, кто готов специализироваться и любит видеть результат на экране устройства.

## Реальность 2026

**Вакансии:**
- Kotlin / Android: ~324 в крупных компаниях.
- Swift / iOS: ~148 — заметно меньше.
- React Native / Flutter — серая зона: где-то нанимают, где-то категорически нет.

**Конкуренция:** ниша. На одну Kotlin-вакансию приходит ~10–25 резюме (для сравнения: Frontend — 18–40).

**Зарплаты в Москве 2026:**
- Стажёр: 0–50к (символическая).
- Junior (вакансий мало): 80–160к.
- Middle: 180–380к.
- Senior: 350–650к.
- Lead: 450–800к.

**Где работают mobile-разработчики:**
- Bigtech: Яндекс, Авито, Озон, Сбер, Тинькофф — у каждого свои приложения.
- Финтех и e-commerce — основные потребители native-разработки.
- Стартапы — чаще берут React Native / Flutter из-за экономии.
- Аутсорс — разное.

## iOS vs Android — что выбрать

### iOS (Swift / SwiftUI / UIKit)

**Плюсы:**
- Зарплаты чуть выше Android в среднем (Apple-экосистема даёт премию ~10–20%).
- Меньше фрагментация устройств — один тестируешь как-то и работает на всех.
- Swift как язык приятнее в работе для многих.

**Минусы:**
- Меньше вакансий (148 vs 324 у Kotlin).
- Нужен Mac (~150к₽ на старте).
- Apple-экосистема жёстче: жёсткие гайдлайны, медленный апрув в App Store.

### Android (Kotlin / Jetpack Compose)

**Плюсы:**
- Больше вакансий, особенно в Bigtech и Tinkoff.
- Не нужен Mac — любой ноут.
- Open-source инструменты, проще найти материалы.

**Минусы:**
- Зарплаты чуть ниже iOS.
- Фрагментация: тысячи устройств, разные версии Android, разные оболочки производителей.
- Java-наследие в legacy-коде (если попадёшь на старый проект).

### Кросс-платформа (React Native / Flutter / KMP)

**Плюсы:**
- Один код на iOS и Android.
- Проще войти если уже знаешь React (для RN) или хочешь Dart (Flutter).

**Минусы:**
- В Bigtech на чистый кросс-платформ редко берут — там native.
- Native-проблемы всё равно вылезают (камера, нотификации, deep links) и приходится писать платформенный код.
- Зарплаты ниже native (~10–15%).

**Стратегия выбора:**
- Хочу Bigtech → Native (Kotlin или Swift).
- Хочу стартап / агентство → React Native или Flutter.
- Уже React-разработчик и хочу мобильную разработку → React Native.
- С нуля и больше вакансий → **Kotlin (Android)** — оптимум.
- С нуля и есть Mac + любишь Apple → Swift (iOS).

## Стек 2026 — что должно быть в наборе

### Android (Kotlin)

**Язык:**
- Kotlin: coroutines, flows, scope functions, data classes, sealed classes.
- Базовое понимание Java (legacy в проектах встречается).

**UI:**
- Jetpack Compose — стандарт в новых проектах с 2023.
- XML Views — для legacy кода (остаётся в продакшене).
- Material Design 3.

**Архитектура:**
- MVVM или MVI.
- Clean Architecture: domain / data / presentation слои.
- Single Activity + Navigation Component.

**Андроид-стек:**
- Jetpack: ViewModel, LiveData/StateFlow, Room, DataStore.
- Dependency Injection: Hilt или Koin.
- Networking: Retrofit + OkHttp.
- Корутины + Flow.
- WorkManager для фоновых задач.

**Тесты:**
- Unit: JUnit, MockK.
- UI: Espresso или Compose Test.

**Инфра:**
- Firebase basics (Crashlytics, Analytics, Cloud Messaging).
- Git, gradle build system.
- Play Store releases (signed APK/AAB, staged rollouts).

### iOS (Swift)

**Язык:**
- Swift: optionals, generics, protocols, async/await.
- Базовое Objective-C (legacy).

**UI:**
- SwiftUI — стандарт в новых проектах.
- UIKit — для legacy и сложных экранов.
- AutoLayout / SwiftUI declarative layout.

**Архитектура:**
- MVVM или TCA (The Composable Architecture).
- Clean Architecture.
- Coordinators для навигации в UIKit.

**iOS-стек:**
- Combine или async/await — реактивность.
- Core Data / SwiftData — локальная база.
- URLSession для сети, Alamofire опционально.
- DI: вручную через Composition Root или Factory.
- Push-нотификации (APNs).

**Тесты:**
- Unit: XCTest, Quick/Nimble.
- UI: XCUITest.

**Инфра:**
- Xcode, Swift Package Manager (или CocoaPods для legacy).
- TestFlight, App Store Connect.
- Provisioning profiles (боль).

## Реалистичный путь входа в Mobile

Канон обучения — через практику с AI. Лекции «впрок» не работают. AI как ментор 24/7 (намёк → декомпозиция → объяснение, **никогда готовый код**). См. [[01_Образование и обучение/06_Менторы и AI-обучение/_Философия обучения через практику|Философия обучения через практику]].

**Месяц 1–2: язык + базовый UI**
- Kotlin: основы синтаксиса, coroutines, классы. Книга «Kotlin in Action» или Kotlin Koans.
- Или Swift: основы, optionals, protocols. Apple Swift Tour + Stanford CS193p (бесплатный курс).
- Первый экран: список + детальный просмотр (todo, заметки, погода).

**Месяц 2–4: фреймворк + архитектура**
- Compose / SwiftUI до уровня «могу собрать 5–7 экранов с состоянием».
- MVVM на практике.
- Networking: REST API + JSON parsing.
- Локальное хранение: Room / Core Data.

**Месяц 4–6: продакшен-уровень + пет-проект**
- Полноценный пет-проект: список + детали + поиск + offline-режим + push.
- Деплой в Play Store / TestFlight (внутреннее тестирование, не нужен публичный релиз).
- Тесты, CI (GitHub Actions для Android, fastlane для iOS).

**Месяц 6–9: подготовка к собесу + поиск middle-оффера**
- Банк вопросов по платформе (Compose internals, lifecycle, мемори-менеджмент).
- 10–15 типовых mobile-задач (lifecycle bugs, retain cycles, recomposition, navigation) на Kotlin/Swift. LeetCode-гринд под mobile-собес почти бесполезен — гоняют lifecycle и memory, а не алгоритмы.
- Дополнительный пет-проект с интересной фичей (offline-first, WebSocket, ML-on-device).
- Параллельный поиск 15–20 компаний.

## Темы, которые точно спросят на middle-собесе

### Android
- Lifecycle: Activity vs Fragment vs Compose lifecycle, что когда вызывается.
- Memory leaks: как возникают (anonymous Handler, static refs), как ловить (LeakCanary).
- Coroutines: scope, dispatcher, structured concurrency, что когда отменяется.
- Compose: рекомпозиция, remember, derivedStateOf, оптимизация.
- Navigation: deep links, back stack, аргументы.
- Background work: WorkManager vs Service vs JobScheduler.
- Build: Gradle, Variants, Flavors, R8 / ProGuard.

### iOS
- Memory management: ARC, strong/weak/unowned references, retain cycles.
- SwiftUI: state, binding, observable object, environment.
- Concurrency: GCD vs async/await, actor, MainActor.
- Navigation: NavigationStack vs Coordinator pattern.
- Background: BGTaskScheduler, push wakeups.
- App lifecycle: scene, foreground/background, state restoration.
- Build: Provisioning, code signing, бandle id, entitlements.
## Антипаттерны

❌ **Учить и Android и iOS параллельно.** Распыление = ноль глубины.
**Решение:** одну платформу до middle-собеса, вторую — после оффера, если захочется.

❌ **Игнорировать тесты.** На mobile тесты обязательны (UI тестирование особенно болезненно без них).
**Решение:** unit-тесты ViewModel/UseCase с первого пет-проекта.

❌ **Не публиковать в TestFlight / Play Console.** Деплой в проде — отдельный болезненный навык. Не пройдёшь раз — провалишь его на работе.
**Решение:** хотя бы один пет-проект задеплой во внутренний тест.

❌ **Кросс-платформ как стартовая точка для Bigtech.** Bigtech нанимает native.
**Решение:** native язык первым, кросс-платформ второй стек.

## Чек-лист готовности к middle-собесу Mobile

### Android
- [ ] Kotlin: coroutines, flows, scope functions, sealed classes — без гугла.
- [ ] Compose: state, recomposition, side effects, оптимизация.
- [ ] MVVM с ViewModel + StateFlow.
- [ ] Networking: Retrofit + Moshi/Kotlinx Serialization.
- [ ] Room: миграции, suspending queries, Flow-возвраты.
- [ ] DI: Hilt или Koin.
- [ ] Navigation Component с back stack.
- [ ] Один пет-проект в Play Console (даже internal).
- [ ] Lifecycle и memory leaks — могу объяснить и продемонстрировать.
- [ ] Unit-тесты с MockK.

### iOS
- [ ] Swift: protocols, generics, async/await, optionals — без гугла.
- [ ] SwiftUI: state/binding/observable, environment.
- [ ] MVVM или TCA.
- [ ] URLSession + JSON Decoder.
- [ ] Core Data или SwiftData с миграциями.
- [ ] Push-нотификации (хотя бы локальные + APNs basics).
- [ ] Один пет-проект в TestFlight.
- [ ] ARC и retain cycles — могу объяснить и продемонстрировать.
- [ ] Unit-тесты с XCTest.

## Связи

- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — сравнение основных стеков.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий 2026]] — сырые цифры.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — детальные зарплаты.

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — общее сравнение стеков, в котором Mobile — узкая ниша.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий по стекам 2026]] — сырые цифры спроса по Kotlin/Swift, на которых строится разбор.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — сравнение Mobile-зарплат с другими специализациями.
- [[01_Образование и обучение/05_Путь обучения по специальностям/_Выбор направления|Выбор направления]] — выбор Mobile среди прочих направлений обучения.
- [[08_Будущее и тренды/01_AI и автоматизация профессий/_Заменит ли AI разработчиков|Заменит ли AI разработчиков]] — какие части mobile-разработки AI закрывает первыми.
<!-- auto-links:end -->
