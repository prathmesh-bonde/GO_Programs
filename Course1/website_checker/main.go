package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", err, link, "is down")

		c <- "Yup its down" // sending data thru channel

		return
	}

	fmt.Println(link, "is up and running")

	c <- "Yup its up"
}

func main() {
	links := []string{
		"https://google.com",
		"https://geeksforgeeks.org",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	mychannel := make(chan string) // creating a channel of type chan and data transfer is of type string

	for _, link := range links {
		go checkLink(link, mychannel) // creating a goroutine and passing channel as argument
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-mychannel)

	}
}
