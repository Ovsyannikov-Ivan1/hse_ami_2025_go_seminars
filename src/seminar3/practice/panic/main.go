package main

import (
	"fmt"
	"runtime/debug"
)

func deepFunction() {
	panic("deep panic")
}

func middleFunction() {
	deepFunction()
}

func topFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
			debug.PrintStack()
		}
	}()
	middleFunction()
}

func main() {
	topFunction()
}
