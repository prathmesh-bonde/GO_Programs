package main

import "fmt"

// exercise 3

// custom struct to implement the built in error interface

type customErr struct {
	info string
}

// this is the built in method of the error interface
func (ce customErr) Error() string {
	return fmt.Sprintf("err: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran - ", e)
	// info is not of error, so we need to use assertion i.e. asserting e is of type customErr
	fmt.Println("foo ran - ", e.(customErr).info)
}

func main() {
	c := customErr{"need more coffee"}
	foo(c)
}
