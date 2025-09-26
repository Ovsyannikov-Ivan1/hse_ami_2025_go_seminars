package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

func main() {
	john := Person{Name: "John", Age: 30}
	fmt.Println(john)
}
