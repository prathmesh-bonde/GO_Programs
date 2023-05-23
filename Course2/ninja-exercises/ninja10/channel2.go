package main

import "fmt"

// exercise 3
// pulling values of channel using for range loop

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}
