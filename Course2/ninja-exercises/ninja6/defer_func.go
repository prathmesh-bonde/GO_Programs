package main

import "fmt"

// exercise 3

func foo() {
	// anonymous func
	defer func() {
		fmt.Println("defer foo exit")
	}()
	fmt.Println("foo exit")
}

func main() {
	defer foo()
	fmt.Println("Hello from main")
}
