---
parent: "[[20_IT-специальности]]"
description: "Data Science / ML Engineering / Data Engineering 2026: три разные роли, разные стеки, разные зарплаты."
type: concept
---

# Data / ML — реальность 2026

«Я хочу в ML» — самая частая фраза новичков, которая обычно означает три разные профессии. Перед выбором — определись, что именно хочешь делать.

## Три роли — три профессии

| Роль | Что делает | Главные инструменты |
|---|---|---|
| **Data Scientist** | A/B тесты, аналитика, статистика, прототипы ML-моделей | Python (pandas, numpy), SQL, Jupyter, scikit-learn |
| **ML Engineer** | Деплоит модели в продакшен, MLOps, оптимизация инференса | Python, PyTorch/TF, Docker, k8s, MLflow, FastAPI |
| **Data Engineer** | Строит pipelines данных, хранилища, ETL | Python, SQL, Airflow, Spark, dbt, Snowflake/Clickhouse |

**Главное отличие на собесе:**
- Data Scientist — спросят про статистику, A/B тесты, базовый ML.
- ML Engineer — спросят про деплой моделей, latency, batch vs online inference, мониторинг моделей.
- Data Engineer — спросят про SQL глубоко, оркестрацию, форматы данных (Parquet, Avro), partitioning.

## Реальность 2026

**Спрос вырос в 3–5 раз за 2023–2025.** Но это маркетинг. По факту:
- **ML Engineer** — реально дефицит, особенно с опытом продакшен-деплоя.
- **Data Engineer** — стабильно нужны, особенно в финтехе и e-commerce.
- **Data Scientist** — рынок переполнен junior-ами после курсов. Реальная цель — middle+ с опытом конкретных продуктовых задач.

**Вакансии:** ~458 в крупных компаниях (объединённый сегмент Data/ML).

**Конкуренция:** на middle ML Engineer ~30–60 резюме (умеренная). На Data Scientist junior — 200–300 резюме (запредельная).

**Зарплаты в Москве 2026:**

| Роль | Стажёр | Junior (мало вакансий) | Middle | Senior | Lead |
|---|---|---|---|---|---|
| Data Scientist | 0–50к | 100–180к | 250–400к | 400–700к | 600–1000к |
| ML Engineer | 0–50к | 120–200к | 280–500к | 500–900к | 700–1200к |
| Data Engineer | 0–50к | 100–180к | 250–450к | 450–800к | 600–1000к |

ML Engineer платят выше из-за дефицита и сложности роли (нужны и DS-знания, и DevOps).

## Какую роль выбрать

### Бери Data Engineer, если:
- Любишь данные, не любишь модели.
- Сильный SQL, готов углубляться.
- Хочешь стабильный спрос без хайпа.
- ОК с инфраструктурой (Airflow, k8s).

### Бери ML Engineer, если:
- Готов учить и DS, и DevOps.
- Любишь видеть, как модель работает в продакшене.
- Хочешь самые высокие зарплаты в сегменте.
- Готов к 9–12 месяцам подготовки минимум.

### Бери Data Scientist, если:
- Любишь статистику, гипотезы, эксперименты.
- Готов к продуктовым задачам, не только ML.
- Имеешь матбэкграунд (ВУЗ + сильная теорвер/линал).
- Готов к высокой конкуренции на входе.

## Стек 2026 — что должно быть в наборе

### Data Scientist

**Базовый минимум:**
- Python (pandas, numpy, matplotlib/seaborn).
- SQL глубоко (window functions, CTE, оптимизация).
- Статистика (распределения, тесты гипотез, A/B тестирование).
- Базовый ML: scikit-learn (regression, classification, clustering).

**Продвинутый:**
- A/B тестирование на практике (CUPED, peeking problem, multi-arm bandits).
- Causal inference (uplift modeling, propensity scoring).
- Time series (ARIMA, Prophet, SARIMA).
- Базовый deep learning (PyTorch / TensorFlow).
- Bayesian basics.

**Soft:**
- Презентация результатов нетехническим стейкхолдерам.
- Дашборды (Tableau / Metabase / Superset).

### ML Engineer

**Базовый минимум:**
- Python на продакшен-уровне (типизация, тесты, async).
- ML фреймворки: PyTorch (популярнее) или TensorFlow.
- Классический ML (sklearn) + базовый deep learning.
- Docker, базовые Linux/CI/CD.

**Продвинутый:**
- Деплой моделей: TorchServe, BentoML, FastAPI + ONNX/TensorRT.
- Оркестрация ML: MLflow, Kubeflow, Metaflow.
- Feature stores: Feast, Tecton.
- Vector databases: Pinecone, Qdrant, Weaviate.
- LLM-стек: LangChain/LlamaIndex, RAG, fine-tuning, vLLM/TGI inference.
- Мониторинг моделей: data drift, prediction drift, concept drift.

**Кросс-домен:**
- DevOps базовый (k8s, Terraform).
- Backend (REST API для модельных endpoint).

### Data Engineer

**Базовый минимум:**
- SQL глубоко (это основа).
- Python (для pipelines и обработки).
- Один облачный data warehouse: Snowflake, BigQuery, Redshift, Clickhouse.
- Airflow или аналог (Prefect, Dagster).

**Продвинутый:**
- Apache Spark (PySpark) — обработка больших данных.
- Streaming: Kafka, Flink.
- dbt — современный стандарт трансформаций в DWH.
- Форматы: Parquet, Avro, ORC, Delta Lake / Iceberg.
- Data quality: Great Expectations, dbt tests.
- DataOps: lineage tracking (Marquez, OpenLineage).

**Инфра:**
- Docker, k8s basics.
- Terraform для DWH-провайдеров.

## Реалистичный путь входа

Канон обучения — через практику с AI. Никаких лекций «впрок». AI как ментор 24/7. См. [[01_Образование и обучение/06_Менторы и AI-обучение/_Философия обучения через практику|Философия обучения через практику]].

### Data Scientist (с нуля 9–12 мес)

**Месяц 1–2:** Python + SQL + базовая статистика. Курс «Python для анализа данных» (Pandas) + любой SQL-задачник на реальных кейсах (StrataScratch, DataLemur). LeetCode SQL — слабее, потому что задачи там оторваны от реальной работы аналитика.

**Месяц 2–4:** Классический ML — scikit-learn. 5–7 учебных задач на Kaggle (House Prices, Titanic, и реалистичные tabular).

**Месяц 4–6:** A/B тесты, статистика глубже. Проектируешь и анализируешь A/B эксперимент на синтетических данных.

**Месяц 6–9:** Один продакшен-уровневый проект (например, churn prediction для синтетического телекома) с feature engineering, валидацией, презентацией результатов.

**Месяц 9–12:** Подготовка к собесу, параллельный поиск.

### ML Engineer (если уже разработчик: 6–9 мес; с нуля: 12–15 мес)

**Если уже backend-разработчик:**

**Месяц 1–2:** Python для ML, классический ML (sklearn), базовая статистика.

**Месяц 2–4:** PyTorch — deep learning основы. Один компьютер-vision или NLP проект (например, классификация изображений или sentiment analysis).

**Месяц 4–6:** MLOps. Тренируешь модель → упаковываешь в Docker → деплоишь в FastAPI → мониторинг (Prometheus + Grafana).

**Месяц 6–9:** Один продакшен-проект полного цикла: данные → feature engineering → train → deploy → A/B тест → ретрейн pipeline. Подготовка к собесу.

### Data Engineer (с нуля 9–12 мес)

**Месяц 1–2:** SQL до сильного уровня (window functions, оптимизация запросов на 1М+ строк), Python.

**Месяц 2–4:** Один DWH (BigQuery / Snowflake free tier / Clickhouse локально). Загружаешь данные, пишешь сложные запросы.

**Месяц 4–6:** Airflow + dbt. Пишешь pipeline: API → raw stage → processed stage → mart → BI-таблицы.

**Месяц 6–9:** PySpark для batch-обработки. Один streaming pipeline (Kafka → Flink/Spark Streaming → DWH).

**Месяц 9–12:** Production-grade pipeline с тестами, monitoring, CI/CD. Подготовка к собесу.

## Темы, которые точно спросят на middle-собесе

### Data Scientist
- Чем отличается Pearson от Spearman корреляции.
- Что такое p-value, чем плохо «p-hacking».
- Как спроектировать A/B тест: метрики, размер выборки, длительность.
- Bias-variance trade-off, как с этим работать.
- Cross-validation: K-fold, stratified, time series split.
- Imbalanced classes: oversampling vs undersampling vs class weights.
- Регуляризация L1 vs L2, чем отличаются результаты.

### ML Engineer
- Train/val/test split — почему именно так.
- Online vs batch inference, когда что выбирать.
- Feature store: зачем нужен, что хранится.
- Как мониторить модель в продакшене (data drift, prediction drift).
- ONNX, TensorRT — зачем, когда применять.
- Deployment patterns: shadow, canary, blue-green для ML.
- LLM: как сделать RAG, прокидывать context window, оценивать quality (BLEU, BERTScore, human eval).

### Data Engineer
- Star schema vs Snowflake schema.
- OLTP vs OLAP, какие БД для каких задач.
- Partitioning vs sharding.
- Window functions: ROW_NUMBER, RANK, DENSE_RANK — разница.
- Slowly Changing Dimensions Type 1/2/3.
- Apache Spark: RDD vs DataFrame, lazy evaluation, shuffle.
- CAP theorem применительно к БД.
- Как обработать 1ТБ данных, у тебя 1 сервер на 64GB RAM.

## AI на работе: ~20% автономно

В Data/ML AI закрывает ~20% рутины (boilerplate pandas, типовые SQL-запросы, базовые plot-ы). Остальное — постановка гипотез, выбор метрик, дебаг моделей, презентация — делает человек.

Что AI делает хорошо:
- Pandas / SQL запросы по описанию задачи.
- Boilerplate PyTorch (модели, training loop).
- Базовые plotting-сниппеты.
- Объяснение библиотечных API (sklearn, pandas).

Что AI **не делает:**
- Выбор метрик под задачу.
- Постановка гипотез.
- Дебаг «почему модель не учится» (всегда контекст).
- Архитектурные решения (какой DWH, как партиционировать).

LLM-стек особо: знание LangChain / LlamaIndex / vLLM в 2026 — must для ML Engineer (растёт спрос на RAG-приложения и собственный inference).

## Антипаттерны

❌ **Учить «всё DS» одновременно (статистика + ML + DL + NLP + CV).** Распыление.
**Решение:** одно направление до middle-уровня (например, tabular ML), потом расширение.

❌ **Идти DS первым стеком без матбазы.** Курсы не закрывают пробелы в линале/теорвере.
**Решение:** или прокачать матбазу (теорвер + линал) параллельно, или идти Data Engineer (там SQL важнее, чем матан).

❌ **Не делать end-to-end проект.** Тренируешь модели в Jupyter без деплоя — на собесе не пройдёшь как ML Engineer.
**Решение:** обязательный пет-проект полного цикла (data → train → deploy → monitor).

❌ **Игнорировать SQL.** В DS думают «у нас pandas». На работе 50% времени — SQL.
**Решение:** SQL до сильного middle-уровня (window functions, оптимизация).

❌ **Переоценивать LLM-хайп.** Не каждый ML-проект про LLM. Классический ML — 80% продакшена.
**Решение:** сначала классика (sklearn, PyTorch), потом LLM-стек.

## Чек-лист готовности к middle-собесу

### Data Scientist
- [ ] Python pandas/numpy без гугла.
- [ ] SQL window functions и оптимизация.
- [ ] Спроектировал A/B тест end-to-end.
- [ ] Знаю scikit-learn до уровня feature engineering + кросс-валидация.
- [ ] Один продакшен-уровневый проект (с метриками, валидацией, презентацией).
- [ ] Базовая статистика без подсказок (тесты, ошибки I/II рода, мощность).

### ML Engineer
- [ ] Python на продакшен-уровне.
- [ ] PyTorch — могу написать training loop с нуля.
- [ ] Деплой модели в FastAPI + Docker + k8s.
- [ ] MLflow или альтернатива для экспериментов.
- [ ] Один LLM-проект (RAG минимум).
- [ ] Один продакшен ML-pipeline с README и деплоем — для рассказа на собесе. На GitHub-профиль никто не зайдёт, важен факт работающего pipeline.

### Data Engineer
- [ ] SQL глубоко (window functions, оптимизация на 1М строк, EXPLAIN).
- [ ] Python для pipelines.
- [ ] Один DWH (BigQuery / Snowflake / Clickhouse) хорошо.
- [ ] Airflow: написал DAG с conditional logic.
- [ ] dbt: модели, тесты.
- [ ] PySpark basics.
- [ ] Один полный pipeline в работе — для собесного рассказа, не для GitHub-витрины.

## Связи

- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — сравнение основных стеков.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий 2026]] — сырые цифры.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — зарплаты.

## Связи

<!-- auto-links:start -->

### Связанные
- [[20_IT-специальности]] — родительская карта области.
- [[02_IT-специальности и направления/_Какой стек учить в 2026 (сравнение)|Какой стек учить в 2026]] — общее сравнение стеков, в котором Python+ML упомянут как путь.
- [[02_IT-специальности и направления/_Статистика рынка вакансий по стекам 2026|Статистика рынка вакансий по стекам 2026]] — сырые цифры спроса по ML/MLOps на рынке.
- [[01_Образование и обучение/05_Путь обучения по специальностям/_Выбор направления|Выбор направления]] — выбор Data/ML среди прочих направлений обучения.
- [[03_Зарплаты и грейды/01_По специализациям/_Вилки по стекам|Вилки по стекам]] — детализация Data/ML-вилок в сравнении с другими специальностями.
- [[08_Будущее и тренды/01_AI и автоматизация профессий/_Заменит ли AI разработчиков|Заменит ли AI разработчиков]] — Data/ML как роль, делающая AI, а не съедаемая им.
<!-- auto-links:end -->
