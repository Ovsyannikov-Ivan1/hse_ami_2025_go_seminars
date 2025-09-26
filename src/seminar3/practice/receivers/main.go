package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

type Scalable interface {
	Scale(float64)
}

func main() {
	var s Scalable

	myRect := Rectangle{width: 10, height: 5}

	// Это не работает
	// s = myRect

	s = &myRect
	s.Scale(2)
	fmt.Println(myRect)
}
