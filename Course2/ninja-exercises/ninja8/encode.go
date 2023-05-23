package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// exercise 3

type User struct {
	Name   string
	Email  string
	UserID string
}

func main() {

	user1 := User{"Jon", "jon@gmail.com", "jon123"}
	user2 := User{"Max", "max4@gmail.com", "max45"}
	user3 := User{"Dany", "dany20@gmail.com", "dany09"}
	user4 := User{"Joe", "joe@gmail.com", "joe69"}

	users := []User{user1, user2, user3, user4}

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
