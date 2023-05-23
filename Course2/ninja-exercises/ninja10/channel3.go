package main

import "fmt"

// exercise 4

// making c2 bidirectional
func gen(c2 chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		c2 <- 1 // sending data on channel c2
		close(c)
	}()

	return c
}

func receive(c, c2 <-chan int) {
	for {
		select {
		case i := <-c:
			fmt.Println(i)

		case <-c2:
			return
		}
	}
}

func main() {
	c2 := make(chan int)
	c := gen(c2)

	receive(c, c2)

	fmt.Println("about to exit")
}
