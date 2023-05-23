package main

// exercise 5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var x int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&x, 1)
			fmt.Println(atomic.LoadInt64(&x))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x:", x)
}
