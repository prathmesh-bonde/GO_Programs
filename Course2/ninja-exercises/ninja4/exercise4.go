package main

import "fmt"

func main() {
	x := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	x = append(x, 32)
	fmt.Println(x)
	x = append(x, 33, 34, 35)
	fmt.Println(x)
	y := []int{36, 37, 38, 39, 40}
	x = append(x, y...)
	fmt.Println(x)
}
