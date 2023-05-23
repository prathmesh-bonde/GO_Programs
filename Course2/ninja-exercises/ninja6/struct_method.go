package main

import "fmt"

// exercise 4

type person struct {
	first string
	last  string
	age   int
}

// attaching method to struct
func (p person) speak() {
	fmt.Println("Name of person: ", p.first, p.last)
	fmt.Println("Age of person: ", p.age)
}

func main() {
	p1 := person{
		first: "John",
		last:  "Doe",
		age:   25,
	}

	// calling method
	p1.speak()
}
