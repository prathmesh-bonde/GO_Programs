package main

// exercise2

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func say(h human) {
	h.speak()
}

func main() {
	p := person{"ABC"}

	// say(p)    // will not work

	// both will work
	say(&p)
	p.speak()
}
