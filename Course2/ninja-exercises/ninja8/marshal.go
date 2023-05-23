package main

import (
	"encoding/json"
	"fmt"
)

// exercise 1

type User struct {
	Name     string
	Email    string
	UserID   string
	Password string
}

func main() {
	user1 := User{"Jon", "jon@gmail.com", "jon123", "jon123"}
	user2 := User{"Max", "max4@gmail.com", "max45", "max123"}
	user3 := User{"Dany", "dany20@gmail.com", "dany09", "dany20"}
	user4 := User{"Joe", "joe@gmail.com", "joe69", "joe69"}

	users := []User{user1, user2, user3, user4}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
