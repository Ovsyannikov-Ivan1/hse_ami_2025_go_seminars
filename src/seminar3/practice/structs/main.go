package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Introduce() {
	fmt.Printf("Hi, I'm %s.\n", p.Name)
}

func Birthday(p Person) {
	p.Age++ // Изменяется КОПИЯ структуры! Исходная не поменяется.
}

func RealBirthday(p *Person) { // Принимаем указатель
	p.Age++ // Изменяем оригинальную структуру через указатель
	// Автоматическое разыменование: не нужно писать (*p).Age
}

// Employee "наследует" от Person
type Employee struct {
	Person   // Встраивание (Embedding). Не просто поле, а безымянное.
	JobTitle string
	Company  string
}

type Counter struct {
	value int
}

// GetValue Pointer receiver (работает с копией)
func (c *Counter) GetValue() int {
	return c.value
}

// Increment Pointer receiver (работает с оригиналом)
func (c *Counter) Increment() {
	c.value++
}

type User struct {
	ID        int       `json:"id" db:"user_id"`                // Разные теги для разных случаев
	Name      string    `json:"name" validate:"required,min=3"` // Валидация
	Email     string    `json:"email,omitempty"`                // omitempty - не выводить поле, если оно пустое
	CreatedAt time.Time `json:"created_at"`
}

type Config struct {
	apiKey  string
	timeout time.Duration
}

// NewConfig Конструктор. Полезен для валидации и установки значений по умолчанию.
func NewConfig(apiKey string, timeoutSeconds int) (*Config, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey cannot be empty")
	}
	if timeoutSeconds <= 0 {
		timeoutSeconds = 30 // Значение по умолчанию
	}
	return &Config{
		apiKey:  apiKey,
		timeout: time.Duration(timeoutSeconds) * time.Second,
	}, nil
}

// GetAPIKey Геттер для apiKey (чтобы сделать приватное поле доступным для чтения)
func (c *Config) GetAPIKey() string {
	return c.apiKey
}

type Point struct {
	X, Y int
}

func main() {
	// Инициализация
	// 1. Полная инициализация (порядок полей важен!)
	p1 := Person{"Alice", 30, "Moscow"}
	fmt.Println(p1)

	// 2. Инициализация по именам полей (рекомендуется!)
	p2 := Person{
		Name: "Bob",
		Age:  25,
		City: "Berlin", // Запятая в конце обязательна
	}
	fmt.Println(p2)

	// 3. Создание с нулевым значением
	var p3 Person // Name: "", Age: 0, City: ""
	fmt.Println(p3)

	// 4. Создание указателя на структуру с инициализацией
	p4 := &Person{Name: "Charlie"}
	fmt.Println(p4)

	bob := Person{Name: "Bob", Age: 25}
	Birthday(bob)
	fmt.Println(bob.Age) // 25

	RealBirthday(&bob)   // Передаем указатель
	fmt.Println(bob.Age) // 26

	c := Counter{}
	c.Increment()             // Go автоматически преобразует &c в *Counter
	fmt.Println(c.GetValue()) // 1

	u := User{ID: 1, Name: "Alice", CreatedAt: time.Now()}
	data, _ := json.Marshal(u) // Сериализуем в JSON
	fmt.Println(string(data))

	emp := Employee{
		Person:   Person{Name: "Tomoa", Age: 30, City: "Tokyo"}, // Инициализация встроенной структуры
		JobTitle: "Developer",
		Company:  "Google",
	}

	// Мы можем обращаться к полям и методам Person напрямую!
	emp.Introduce()       // "Hi, I'm Tomoa." (метод унаследован)
	fmt.Println(emp.Name) // "Tomoa" (поле унаследовано)

	// Также доступно через полное имя
	fmt.Println(emp.Person.Name) // "Tomoa"

	conf, err := NewConfig("my-secret-key", 10)
	if err != nil {
		panic(err) // Не стоит пробрасывать панику в коде в принципе
	}
	fmt.Println(conf.GetAPIKey())

	point1 := Point{1, 2}
	point2 := Point{1, 2}
	point3 := Point{2, 3}

	fmt.Println(point1 == point2)
	fmt.Println(point1 == point3)
}
