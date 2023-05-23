package main

import (
	"fmt"
	"runtime"
)

// exericse 7

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("Routines: ", runtime.NumGoroutine())
	}

	return c
}

func main() {
	x := 10
	y := 5

	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("Routines: ", runtime.NumGoroutine())
}
