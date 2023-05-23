package main

import "fmt"

func main() {
	i := 10
	for i <= 20 {
		fmt.Println(i)
		i++
	}

	for {
		if i >= 20 {
			break
		}
		fmt.Println(i)
		i++
	}
}
