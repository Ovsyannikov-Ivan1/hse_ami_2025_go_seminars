package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["a"] = 1
	fmt.Println(a)
	value, ok := a["b"]
	if !ok {
		fmt.Println("key 'b' does not exist")
	}
	fmt.Println(value)
}
