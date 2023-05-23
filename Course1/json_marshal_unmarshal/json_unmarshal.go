package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Michael", "Last":"Scott", "Age":40}, {"First":"Jim", "Last": "Halpert", "Age":25}]`
	bs := []byte(s)
	fmt.Printf("%T", s)
	fmt.Printf("%T", bs)

	var people []Person

	if err := json.Unmarshal(bs, &people); err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

	for i, v := range people {
		fmt.Printf("Person %v: %v\n", i+1, v)
	}
}
