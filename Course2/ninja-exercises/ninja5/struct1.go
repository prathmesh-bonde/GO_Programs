package main

import "fmt"

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {
	p1 := person{"Michael", "Scott", []string{"chocolate", "mango", "mint"}}
	p2 := person{"Karen", "Filipelli", []string{"coffee", "red velvet", "vanilla"}}

	fmt.Println(p1)
	fmt.Println(p2)
}
