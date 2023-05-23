package main

// exercise 9

import "fmt"

func foo(f func(xi []int) int) int {
	n := f([]int{1, 2, 3, 4, 5})
	n++
	return n
}

func main() {
	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}

		tot := 0
		for _, v := range xi {
			tot += v
		}
		return tot
	}
	// callback
	x := foo(g)

	fmt.Println(x)
}
