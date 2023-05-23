package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U", i)
	}
}
