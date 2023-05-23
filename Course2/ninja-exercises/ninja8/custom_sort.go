package main

import (
	"fmt"
	"sort"
)

// exercise 5

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s: %d", p.First, p.Last, p.Age)
}

// ByAge implements sort.Interface methods for []Person based on Age field
type ByAge []Person
type ByLast []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (l ByLast) Len() int {
	return len(l)
}

func (l ByLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ByLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func main() {
	p1 := Person{"Joe", "Root", 20}
	p2 := Person{"Mike", "Scott", 14}
	p3 := Person{"Jim", "Hal", 34}
	p4 := Person{"Pam", "Bees", 24}

	people := []Person{p1, p2, p3, p4}
	// custom sort methods
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByLast(people))
	fmt.Println(people)

}
