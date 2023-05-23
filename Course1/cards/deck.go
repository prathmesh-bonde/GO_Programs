package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string
 
// creating new deck
func newDeck() deck {
	// creating an empty deck
	cards := deck{}

	// using 2 diff slices each for suits and values
	cardSuits := deck{"Spade", "Club", "Heart", "Diamond"}
	cardValues := deck{"Ace", "Two", "Three", "Four"}

	// putting cards in deck
	// we use _ to represent unused variables
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/* receiver function
this method is defined for the datatype 'deck', thus any variable/instance
of this datatype can access this method
eg. cards is a variable of dtype deck which can access the func print()

*/

// below func, d is the actual copy of deck we're working with, i.e. d refers to cards here
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// coverting slice of string to string
func (d deck) toString() string {
	// []string(d) type conversion

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// ioutil.WriteFile() is now deprecated, it is now os.WriteFile
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)

	// if err
	if err != nil {
		// log error and quit prog
		fmt.Println("Error:", err)
		os.Exit(1) // Exit(0) means success, other int represents error
	}

	// string(byteSlice) // convert to string
	// convert to slice of strings
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPostn := r.Intn(len(d) - 1)

		d[i], d[newPostn] = d[newPostn], d[i]
	}
}
