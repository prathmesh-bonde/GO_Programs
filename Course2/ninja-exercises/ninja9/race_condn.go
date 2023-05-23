package main

// exercise 3

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	x := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			y := x
			runtime.Gosched()
			y++
			x = y
			fmt.Println(x)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x:", x)
}
