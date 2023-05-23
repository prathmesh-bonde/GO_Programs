package main

// exercise 4

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	x := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			y := x
			y++
			x = y
			fmt.Println(x)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("x:", x)
}
