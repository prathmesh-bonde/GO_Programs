package main

import "fmt"

func main() {
	x := []int{22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
