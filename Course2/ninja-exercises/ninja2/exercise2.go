package main

import "fmt"

func main() {
	a := (20 == 20)
	b := (20 != 21)
	c := (20 <= 21)
	d := (20 >= 20)
	e := (20 < 21)
	f := (20 > 21)

	fmt.Println(a, b, c, d, e, f)
}
