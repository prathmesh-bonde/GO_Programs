package main

import "fmt"

// exercise 8
// func returning another function

func foo() func() {
	fmt.Println("Hello from foo")
	// return func() int {
	// 	return 40
	// }
	return bar
}

func bar() {
	fmt.Println("Hello from bar")
}

func main() {
	f := foo()

	f()
	// fmt.Println(f())
}
