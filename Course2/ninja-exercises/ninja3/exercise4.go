package main

import "fmt"

func main() {
	yr := 2000

	for {
		if yr > 2023 {
			break
		}
		fmt.Println(yr)
		yr++
	}
}
