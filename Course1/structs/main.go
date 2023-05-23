package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	person1 := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email: "jimhalpert01@hotmail.com",
			zip:   6969,
		},
	}
	// person1.updateName("Michael", "Scott")
	person1.print()

	// using pointer pass by value
	// personPointer := &person1

	// & -> address of variable in memory
	// .pointer -> value at memory
	// personPointer.updateName("Michael", "Scott")

	// pointer shortcut, even if receiver func defined of type pointer, works on root dtype struct
	person1.updateName("Michael", "Scott")
	person1.print()
}

// func (p person) updateName(newFirstName string, newLastName string) {
// 	p.firstName = newFirstName
// 	p.lastName = newLastName
// }

func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
