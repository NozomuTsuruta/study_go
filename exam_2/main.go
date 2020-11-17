package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	bigSquare := square{sideLength: 10}
	smallTriangle := triangle{height: 2, base: 3}
	printArea(bigSquare)
	printArea(smallTriangle)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
