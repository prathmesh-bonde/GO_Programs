package main

// exercise 1

import "fmt"

func main() {
	v := 10
	p := &v

	fmt.Printf("%T\n", v)
	fmt.Printf("%T\n", p)
	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", *p)
}
