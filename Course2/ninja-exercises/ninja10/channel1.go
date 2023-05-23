package main

import "fmt"

// exercise 1, 2

func main() {
	c := make(chan int)

	// c <- 50
	// fmt.Println(<-c)     // deadlock, channel block, main goroutine is asleep

	// using anon func
	go func() {
		c <- 50
	}()

	// using buffered channel
	// c := make(chan int, 1)

	fmt.Println(<-c)

	fmt.Printf("c\t%T\n", c)
}
