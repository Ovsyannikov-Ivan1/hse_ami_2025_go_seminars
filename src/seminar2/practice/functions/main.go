package main

import (
	"errors"
	"fmt"
	"log"
)

func calculate(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func getMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Go полностью поддерживает замыкания
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	result = a / b
	return
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func process(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("empty data")
	}
	// Обработка данных...
	return 1, nil
}

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func someFunction() error {
	return &MyError{Code: 404, Message: "Not found"}
}

func main() {
	add := func(a, b int) int {
		return a + b
	}
	result := calculate(add, 1, 3)
	fmt.Println(result)

	double := getMultiplier(2)
	fmt.Println(double(5))

	c := counter()
	fmt.Println(c())
	fmt.Println(c())

	fmt.Println(divide(1.0, 2.0))

	fmt.Println(sum(1, 2, 3))

	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
	nums = append(nums, []int{5, 6}...)
	nums = append(nums, 5, 6)

	result, err := process([]byte{})
	if err != nil {
		log.Printf("Error processing data: %v", err)
	}
}
