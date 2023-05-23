package main

import (
	"fmt"
	"runtime"
	"sync"
)

// adding waitgroup to sync and run the goroutine created
var wg sync.WaitGroup

func foo() {
	for i := 0; i < 3; i++ {
		fmt.Println("foo")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 3; i++ {
		fmt.Println("bar")
	}
}

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GORoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	// creating a goroutine
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GORoutines\t", runtime.NumGoroutine())
	wg.Wait()
}
