package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// exercise 2

type person struct {
	First   string
	Last    string
	Sayings []string
}

// also needs to return an error
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("Error in toJSON, err: %v", err))
	}

	return bs, nil
}

func main() {
	p := person{"Mike", "Scott", []string{"Hello", "That's what she said"}}

	bs, err := toJSON(p)
	if err != nil {
		log.Fatalln("Json didnt marshal, err: ", err)
	}

	fmt.Println(string(bs))
}
