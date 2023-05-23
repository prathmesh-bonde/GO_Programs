package main

import "fmt"

func main() {
	evenOdd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range evenOdd {
		if num%2 == 0 {
			fmt.Println(num, "Even")
		} else {
			fmt.Println(num, "Odd")
		}
	}
}
