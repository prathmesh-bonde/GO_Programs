package main

import (
	"fmt"
	"sort"
)

// exercise 4

func main() {
	xi := []int{2, 6, 7, 5, 12, 3}
	s := []string{"asdjsd", "asdsads", "abc", "ghd", "def", "bcv"}

	// inbuilt sort methods
	sort.Ints(xi)
	sort.Strings(s)
	fmt.Println(xi)
	fmt.Println(s)
}
