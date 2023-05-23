package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface methods for []Person based on Age field
type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func main() {
	xi := []int{2, 6, 7, 5, 12, 3}
	s := []string{"asdjsd", "asdsads", "abc", "ghd", "def", "bcv"}

	// inbuilt sort methods
	sort.Ints(xi)
	sort.Strings(s)
	fmt.Println(xi)
	fmt.Println(s)

	p1 := Person{"A", 20}
	p2 := Person{"sdsd", 14}
	p3 := Person{"defef", 34}
	p4 := Person{"G", 24}

	people := []Person{p1, p2, p3, p4}
	// custom sort methods
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
