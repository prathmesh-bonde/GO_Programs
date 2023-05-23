package main

import "fmt"

func main() {
	x := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	fmt.Println(x)

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
