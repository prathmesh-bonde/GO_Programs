package main

import (
	"encoding/json"
	"fmt"
)

// type person struct {   // will not work as fieldnames are not capitalized
// 	first string
// 	last  string
// 	age   int
// }

type Person struct { // this will work with marshalling
	First string
	Last  string
	Age   int
}

func main() {
	p1 := Person{
		"John", "Doe", 24,
	}
	p2 := Person{
		"Michael", "Scott", 40,
	}

	people := []Person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}
}
