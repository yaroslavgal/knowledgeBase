package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"
)

// ============================================================
// ВОПРОС 1: Зачем выделять интерфейс если одна реализация?
// ============================================================

type Notifier interface {
	Send(to, msg string) error
}

type SMSNotifier struct {
	APIKey string
}

func (s *SMSNotifier) Send(to, msg string) error {
	fmt.Printf("[SMS] отправлено %q на %s\n", msg, to)
	return nil
}

type MockNotifier struct {
	Sent []string
}

func (m *MockNotifier) Send(to, msg string) error {
	m.Sent = append(m.Sent, fmt.Sprintf("%s: %s", to, msg))
	return nil
}

type OrderService struct {
	notifier Notifier
}

func (o *OrderService) PlaceOrder(userPhone string) {
	fmt.Println("Заказ создан")
	o.notifier.Send(userPhone, "Ваш заказ принят!")
}

func question1() {

	// Продакшн:
	svc := &OrderService{notifier: &SMSNotifier{APIKey: "abc123"}}
	svc.PlaceOrder("+7999")

	// Тест (без реальных SMS!):
	mock := &MockNotifier{}
	svcTest := &OrderService{notifier: mock}
	svcTest.PlaceOrder("+7111")
	fmt.Printf("Mock записал: %v\n\n", mock.Sent)
}

// ============================================================
// ВОПРОС 2: Что если BoardManager не реализует метод интерфейса?
// ============================================================

type TaskManager interface {
	Add(title string) int
	Delete(id int) error
	List() []string
}

type BoardManager struct {
	tasks  map[int]string
	nextID int
}

func (b *BoardManager) Add(title string) int {
	b.nextID++
	b.tasks[b.nextID] = title
	return b.nextID
}

func (b *BoardManager) Delete(id int) error {
	if _, ok := b.tasks[id]; !ok {
		return fmt.Errorf("task %d not found", id)
	}
	delete(b.tasks, id)
	return nil
}

// List() — ЗАКОММЕНТИРУЙ чтобы увидеть ошибку компиляции:
// cannot use &bm (value of type *BoardManager) as TaskManager value
// in variable declaration: *BoardManager does not implement TaskManager
// (missing method List)
func (b *BoardManager) List() []string {
	result := make([]string, 0, len(b.tasks))
	for id, title := range b.tasks {
		result = append(result, fmt.Sprintf("#%d: %s", id, title))
	}
	return result
}

func question2() {
	fmt.Println("=== ВОПРОС 2: Неполная реализация интерфейса ===")

	bm := &BoardManager{tasks: make(map[int]string)}
	var tm TaskManager = bm // ← компилятор проверяет ВСЕ методы здесь

	tm.Add("Фикс бага")
	tm.Add("Рефакторинг")
	fmt.Println(tm.List())
	fmt.Println()
}

// ============================================================
// ВОПРОС 3: Почему sentinel errors — var, а не const?
// ============================================================

// ✅ Правильно — var:
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
)

// ❌ Не скомпилируется — const:
// const ErrBad = errors.New("bad") // ← ошибка: errors.New() не константа

// Почему: errors.New() возвращает *errorString (указатель на структуру).
// const в Go может хранить только примитивы: int, string, bool, float.
// Указатели, структуры, слайсы — не могут быть const.

func question3() {
	fmt.Println("=== ВОПРОС 3: Sentinel errors — var, не const ===")

	// Sentinel работает с errors.Is через сравнение указателей:
	err := fmt.Errorf("repo.GetUser(42): %w", ErrNotFound)

	fmt.Println("err == ErrNotFound?       ", err == ErrNotFound)            // false (разные объекты)
	fmt.Println("errors.Is(err, ErrNotFound)?", errors.Is(err, ErrNotFound)) // true (разворачивает цепочку)

	// Цепочка ошибок:
	err2 := fmt.Errorf("service: %w", err)
	fmt.Println("двойная обёртка — Is?     ", errors.Is(err2, ErrNotFound)) // true
	fmt.Println("полная ошибка:            ", err2)
	fmt.Println()
}

// ============================================================
// ВОПРОС 4: Валидация переходов статусов через map
// ============================================================

// Плохо — switch на каждый переход:
func canTransitionBad(from, to string) bool {
	switch from {
	case "todo":
		return to == "in_progress"
	case "in_progress":
		return to == "done" || to == "todo"
	case "done":
		return false
	}
	return false
}

// Хорошо — map переходов (расширяется без правки логики):
var validTransitions = map[string][]string{
	"todo":        {"in_progress"},
	"in_progress": {"done", "todo", "review"}, // легко добавить новый статус
	"review":      {"done", "in_progress"},
	"done":        {},
}

func canTransition(from, to string) bool {
	allowed, exists := validTransitions[from]
	if !exists {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

func question4() {
	fmt.Println("=== ВОПРОС 4: Валидация переходов статусов ===")

	transitions := [][2]string{
		{"todo", "in_progress"},   // ✅
		{"todo", "done"},          // ❌ нельзя пропустить
		{"in_progress", "done"},   // ✅
		{"in_progress", "todo"},   // ✅ откат
		{"in_progress", "review"}, // ✅ новый статус
		{"done", "todo"},          // ❌
	}

	for _, t := range transitions {
		ok := canTransition(t[0], t[1])
		mark := "❌"
		if ok {
			mark = "✅"
		}
		fmt.Printf("  %s  %s → %s\n", mark, t[0], t[1])
	}
	fmt.Println()
}

// ============================================================
// ВОПРОС 5: Выдели минимальный интерфейс
// ============================================================

// ДО рефакторинга — App зависит от конкретного FileLogger:
type FileLogger struct {
	Path string
}

func (l *FileLogger) Log(msg string) { fmt.Printf("[FILE:%s] %s\n", l.Path, msg) }
func (l *FileLogger) Close() error   { fmt.Println("[FILE] closed"); return nil }
func (l *FileLogger) Rotate() error  { fmt.Println("[FILE] rotated"); return nil }

// ПОСЛЕ рефакторинга — минимальный интерфейс (только то что App использует):
type AppLogger interface {
	Log(msg string) // App только логирует. Close/Rotate — не его забота.
}

// Тестовый логгер:
type TestLogger struct {
	Messages []string
}

func (t *TestLogger) Log(msg string) {
	t.Messages = append(t.Messages, msg)
}

type App struct {
	logger AppLogger // ← интерфейс из 1 метода, не *FileLogger
}

func (a *App) Start() {
	a.logger.Log("app started")
	a.logger.Log("loading config")
}

func question5() {
	fmt.Println("=== ВОПРОС 5: Минимальный интерфейс ===")

	// Продакшн:
	app1 := &App{logger: &FileLogger{Path: "/var/log/app.log"}}
	app1.Start()

	// Тест (проверяем ЧТО логируется, без файлов):
	tl := &TestLogger{}
	app2 := &App{logger: tl}
	app2.Start()
	fmt.Printf("Тест — залогировано: %v\n\n", tl.Messages)
}

// ============================================================
// ВОПРОС 6: Что выведет? (горутина может не успеть)
// ============================================================

func question6() {
	fmt.Println("=== ВОПРОС 6: Горутина может не успеть ===")

	// Вариант 1 — БЕЗ ожидания (недетерминированный результат):
	fmt.Println("--- Без ожидания (5 запусков): ---")
	for run := 0; run < 5; run++ {
		done := make(chan bool, 1)
		go func() {
			done <- true
		}()

		// Даём 1 микросекунду — горутина может не успеть
		time.Sleep(1 * time.Microsecond)
		select {
		case <-done:
			fmt.Printf("  Запуск %d: горутина УСПЕЛА\n", run+1)
		default:
			fmt.Printf("  Запуск %d: горутина НЕ УСПЕЛА\n", run+1)
		}
	}

	// Вариант 2 — С WaitGroup (всегда стабильный):
	fmt.Println("--- С WaitGroup (гарантированно): ---")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  hello из горутины")
	}()
	wg.Wait()
	fmt.Println("  world из main")
	fmt.Println()
}

// ============================================================
// ВОПРОС 7: Найди две ошибки (Add внутри горутины + замыкание)
// ============================================================

func question7() {
	fmt.Println("=== ВОПРОС 7: Две классические ошибки ===")

	// ❌ ПЛОХОЙ КОД (обе ошибки):
	fmt.Println("--- Багованный вариант: ---")
	var wg1 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg1.Add(1) // фикс #1: Add ДО go (иначе Wait может завершиться раньше)
		go func() {
			// wg1.Add(1) ← было тут — гонка!
			defer wg1.Done()
			fmt.Printf("  i=%d ", i) // ← ошибка #2: i захвачена по ссылке
		}()
	}
	wg1.Wait()
	fmt.Println("\n  ^ скорее всего все 5 или почти все = 5")

	// ✅ ПРАВИЛЬНЫЙ КОД (оба фикса):
	fmt.Println("--- Исправленный вариант: ---")
	var wg2 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg2.Add(1)       // фикс #1: Add ДО go
		go func(n int) { // фикс #2: передаём i как аргумент
			defer wg2.Done()
			fmt.Printf("  n=%d ", n) // каждая горутина имеет свою копию
		}(i)
	}
	wg2.Wait()
	fmt.Println("\n  ^ все числа 0-4 (в случайном порядке)")
	fmt.Println()
}

// ============================================================
// ВОПРОС 8: Горутина vs поток ОС (демонстрация дешевизны)
// ============================================================

func question8() {
	fmt.Println("=== ВОПРОС 8: Горутины дешёвые — запустим 100 000 ===")

	const N = 100_000
	var wg sync.WaitGroup
	wg.Add(N)

	start := time.Now()

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			// Каждая горутина просто спит 10ms
			time.Sleep(10 * time.Millisecond)
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("  Запущено горутин: %d\n", N)
	fmt.Printf("  Общее время: %v\n", elapsed.Round(time.Millisecond))
	fmt.Printf("  Каждая спала 10ms, но все параллельно → ~10-50ms а не %v\n",
		time.Duration(N)*10*time.Millisecond)
	// 100_000 потоков ОС — невозможно. 100_000 горутин — легко.
	fmt.Println()
}

// ============================================================
// ВОПРОС 9: WaitGroup по значению = deadlock
// ============================================================

func question9() {
	fmt.Println("=== ВОПРОС 9: WaitGroup по значению vs указателю ===")

	// ❌ ПО ЗНАЧЕНИЮ — deadlock (закомментировано, иначе программа зависнет):
	fmt.Println("--- По значению (deadlock, закомментировано): ---")
	fmt.Println(`  var wg sync.WaitGroup
  wg.Add(1)
  go func(wg sync.WaitGroup) {  // ← КОПИЯ
      defer wg.Done()            // Done на КОПИИ
      fmt.Println("done")
  }(wg)
  wg.Wait()  // ← ждёт вечно, оригинальный счётчик = 1`)
	fmt.Println("  // fatal error: all goroutines are asleep - deadlock!")

	// ✅ ПО УКАЗАТЕЛЮ — работает:
	fmt.Println("\n--- По указателю (работает): ---")
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) { // ← УКАЗАТЕЛЬ
		defer wg.Done() // Done на ОРИГИНАЛЕ
		fmt.Println("  горутина: done")
	}(&wg) // ← передаём &wg
	wg.Wait()
	fmt.Println("  main: горутина завершилась")
	fmt.Println()
}

// ============================================================
// ВОПРОС 10: Зачем defer wg.Done() а не просто wg.Done()
// ============================================================

func riskyWork(id int) {
	if id == 3 {
		panic("что-то пошло не так в задаче 3!")
	}
	fmt.Printf("  Задача %d выполнена\n", id)
}

func safeWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // ✅ сработает ДАЖЕ при panic

	// recover ловит panic чтобы не уронить всю программу
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Задача %d: PANIC пойман: %v\n", id, r)
		}
	}()

	riskyWork(id)
}

func unsafeWorker(id int, wg *sync.WaitGroup) {
	// без recover panic пробрасывается наверх
	riskyWork(id)
	wg.Done() // ❌ если panic выше — эта строка НЕ выполнится → deadlock
}

func question10() {
	fmt.Println("=== ВОПРОС 10: defer wg.Done() спасает от deadlock при panic ===")

	// ✅ С defer — все горутины корректно завершаются:
	fmt.Println("--- С defer wg.Done() (безопасно): ---")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go safeWorker(i, &wg)
	}
	wg.Wait()
	fmt.Println("  Все задачи обработаны (даже паникнувшая)")

	// ❌ Без defer — deadlock при panic (закомментировано):
	fmt.Println("\n--- Без defer (deadlock, закомментировано): ---")
	fmt.Println(`  go func(wg *sync.WaitGroup) {
      riskyWork(3)   // ← panic!
      wg.Done()      // ← НИКОГДА не выполнится
  }(&wg)
  wg.Wait() // ← ждёт вечно, счётчик не уменьшился`)
	fmt.Println()
}

// ============================================================
// БОНУС: Shape — полиморфизм через интерфейс (из knowledge base)
// ============================================================

type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

type Circle struct{ Radius float64 }

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }
func (c Circle) Name() string       { return "Круг" }

type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }
func (r Rectangle) Name() string       { return "Прямоугольник" }

func printShapeInfo(s Shape) {
	fmt.Printf("  %s: площадь=%.2f, периметр=%.2f\n", s.Name(), s.Area(), s.Perimeter())
}

func bonus() {
	fmt.Println("=== БОНУС: Полиморфизм через интерфейс ===")
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 1},
	}
	for _, s := range shapes {
		printShapeInfo(s)
	}
	fmt.Println()
}

// ============================================================
// MAIN — запуск всех вопросов
// ============================================================

func main() {
	question1()
	question2()
	question3()
	question4()
	question5()
	question6()
	question7()
	question8()
	question9()
	question10()
	bonus()

	// Подавляем unused import
	_ = strings.ToLower
}
