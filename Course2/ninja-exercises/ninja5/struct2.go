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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)
		}
		fmt.Println()
	}
}
