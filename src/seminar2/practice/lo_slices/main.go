package main

import (
	"fmt"
	"slices"

	"github.com/samber/lo"
)

func main() {
	result := lo.Ternary(len(loUniq()) > 3, "A", "B")
	fmt.Println(result)

	fmt.Println(slicesEqualFunc([]string{"A"}, []string{"a"}))

	s := []int{4, 2, 5, 1, 3}
	slices.Sort(s)
	fmt.Println(s)

	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{Name: "Alice", Age: 30},
		{"Bob", 25},
		{"Charlie", 40},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return a.Age - b.Age
		// Правило: отрицательное -> a < b, 0 -> a == b, положительное -> a > b
	})
	fmt.Println(people)

	s = []int{1, 2, 3, 2}
	fmt.Println(slices.IsSorted(s))

	sorted := []int{10, 20, 30, 40, 50}
	pos, found := slices.BinarySearch(sorted, 30)
	fmt.Println(pos, found)

	pos, found = slices.BinarySearch(sorted, 35)
	fmt.Println(pos, found)

	numbers := []int{4, 2, 5, 1, 3}
	fmt.Println(slices.Min(numbers))
	fmt.Println(slices.Max(numbers))

	people = []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 40},
	}
	oldest := slices.MaxFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	fmt.Println(oldest)
}
