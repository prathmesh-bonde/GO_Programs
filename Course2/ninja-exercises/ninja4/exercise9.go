package main

import "fmt"

func main() {
	m := map[string][]string{
		`scott_mike`: {
			`Shaken , not stirred`, `Martinis`, `Women`},
		`hal_jim`: {`Mike Scott`, `Literature`, `Computer Science`},
		`no_dr`:   {`Happy`, `Ice cream`, `Sunsets`},
	}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
