package main

// exercise 6

import "fmt"

const (
	a = 2000 + iota
	b = 2000 + iota
	c = 2000 + iota
	d = 2000 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
