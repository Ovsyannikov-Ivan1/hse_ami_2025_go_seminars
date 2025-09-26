package main

import "fmt"

func DescribeAnything(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	DescribeAnything(42)
	DescribeAnything("hello")
	DescribeAnything(3.14)
	DescribeAnything([]int{1, 2, 3})
}
