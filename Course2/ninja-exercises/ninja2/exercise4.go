package main

import "fmt"

func main() {
	a := 20
	fmt.Printf("%d\t%b\t%#x", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
