package main

import "fmt"

type bot interface { // to refer to all types of bot
	getGreeting() string // generalized method which is to be implemented by all types of bot
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	return "Hola amigo"
}

// above as receiver func is diff dtype, doeesn't give error

// using interface as a parameter
func printGreet(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreet(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreet(sb spanishBot) {
// 	sb.getGreeting()
// }                    // here gives an error as it the same function of same dtype

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreet(eb)
	printGreet(sb)
}
