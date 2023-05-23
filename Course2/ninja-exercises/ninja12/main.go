package main

import (
	"fmt"

	"ninja12/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	dog1 := canine{"Max", dog.Years(5)}

	fmt.Println(dog1)
}
