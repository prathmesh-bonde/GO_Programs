package main

import "fmt"

// exercise 5

func main() {
	c := make(chan int)

	go func() {
		// putting to channel
		c <- 20
	}()

	// pulling from channel
	v, ok := <-c
	fmt.Printf("%v\t%v", v, ok)
	close(c)

	fmt.Printf("%v\t%v", v, ok)
}
