package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("when %v is divided by 4, rem is %v", i, i%4)
	}
}
