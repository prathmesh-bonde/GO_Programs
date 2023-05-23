package main

// exercise 5

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	calcArea() float64
}

func (c circle) calcArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (sq square) calcArea() float64 {
	return sq.side * sq.side
}

func info(s shape) {
	fmt.Println(s.calcArea())
}

func main() {
	circle := circle{10}
	square := square{20}

	info(circle)
	info(square)
}
