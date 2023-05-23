package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.height * tr.base
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println("Area is ", s.getArea())
}

func main() {
	t := triangle{10, 3}
	s := square{20}

	printArea(t)
	printArea(s)
}
