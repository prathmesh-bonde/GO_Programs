package main

import "fmt"

func main() {
	x := "ABC"
	if x == "ABC" {
		fmt.Println(x)
	} else if x == "abc" {
		fmt.Println("abc", x)
	} else {
		fmt.Println("none")
	}

}
