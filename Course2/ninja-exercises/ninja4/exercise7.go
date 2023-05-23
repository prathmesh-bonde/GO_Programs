package main

import "fmt"

func main() {
	xs1 := []string{"Mike", "Scott", "Shaken , not stirred"}
	xs2 := []string{"Jim", "Hal", "Hellooooo, Mike."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index: %v \t value: %v \n", j, val)
		}
	}
}
