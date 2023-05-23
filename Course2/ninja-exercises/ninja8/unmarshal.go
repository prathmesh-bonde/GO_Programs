package main

import (
	"encoding/json"
	"fmt"
)

// exercise 2

type User struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
}

func main() {
	users := `[ {"Name": "John", "Age": 25, "Email": "john02@gmail.com"}, 
				{"Name": "Mike", "Age": 40, "Email": "mike69@hotmail.com"}, 
				{"Name": "Brad", "Age": 32, "Email": "bradtrad32@rediffmail.com"}]`

	var usrs []User
	err := json.Unmarshal([]byte(users), &usrs)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range usrs {
		fmt.Printf("User%v: %v\n", i+1, v)
	}
}
