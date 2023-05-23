package main

// exercise 6

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
