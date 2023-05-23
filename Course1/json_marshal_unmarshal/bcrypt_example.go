package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `pass1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	// loginpw := `pass1234`  // will log in
	loginpw := `pass123` // will not log in
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpw))
	if err != nil {
		fmt.Println("Wrong pw")
		return
	} else {
		fmt.Println("Logged in")

	}
}
