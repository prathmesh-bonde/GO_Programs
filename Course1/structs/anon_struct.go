package main

import "fmt"

func main() {
	person := struct {
		first int
		last  int
		nm    string
	}{
		first: 1,
		last:  10,
		nm:    "hello",
	}

	fmt.Println(person.first, person.last, person.nm)
}
