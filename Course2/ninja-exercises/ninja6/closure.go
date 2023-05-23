package main

import "fmt"

// exercise 10

// enclosing the scope of a variable in a func

func foo() func() int {
	x := 0
	// enclosing scope of x in a anonymous func
	return func() int {
		x++
		return x
	}
}

func main() {
	// so everytime we call foo, we get an increment in x
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	f = foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
