package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// exercise 1

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p := person{"Mike", "Scott", []string{"Hello", "That's what she said"}}

	bs, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("Json didnt marshal, err: ", err)
	}

	fmt.Println(string(bs))
}
