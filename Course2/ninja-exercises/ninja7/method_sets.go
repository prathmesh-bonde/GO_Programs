package main

import "fmt"

// exercise 2

// method sets: methods attached to a type

type person struct {
	name string
}

func changeMe(p *person) {
	(*p).name = "Michael"
	p.name = "Joe" // works the same
}

func main() {
	p := person{"John"}

	fmt.Println(p.name)
	changeMe(&p)
	fmt.Println(p.name)

}
