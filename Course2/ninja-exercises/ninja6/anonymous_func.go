package main

// exercise 6

import (
	"fmt"
)

func main() {
	// anonymous func
	func() {
		tot := 0
		for x := 0; x <= 10; x++ {
			tot += x
		}
		fmt.Println(tot)
	}()
}
