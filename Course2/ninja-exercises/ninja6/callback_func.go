package main

import "fmt"

// exercise 8

// var x = 7     // allowed
var x int
var g func()

// x = 7     // not allowed

func main() {
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)
	fmt.Printf("%T", g)
	g = f
	fmt.Println("--")
}
