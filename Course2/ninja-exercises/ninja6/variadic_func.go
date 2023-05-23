package main

import "fmt"

// exercise 2

// variadic func
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	sint := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// unfurling the slice
	fmt.Printf("%v", foo(sint...))
}
