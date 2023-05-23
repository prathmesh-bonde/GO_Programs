package main

import "fmt"

func main() {
	// creating a channel
	c := make(chan int)
	// setting value as 20 to channel
	c <- 20

	fmt.Println(c) // will not working as all goroutines are asleep
}
