package main

// exercise 1

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}

func main() {
	// at begin
	fmt.Println("numCPU: ", runtime.NumCPU())
	fmt.Println("numGoRoutines: ", runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()

	// before goroutines
	fmt.Println("numCPU: ", runtime.NumCPU())
	fmt.Println("numGoRoutines: ", runtime.NumGoroutine())

	fmt.Println("main")
	wg.Wait()

	// at end
	fmt.Println("Exit")
	fmt.Println("numCPU: ", runtime.NumCPU())
	fmt.Println("numGoRoutines: ", runtime.NumGoroutine())
}
