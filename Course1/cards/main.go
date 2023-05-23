package main

func main() {
	// cards := newCard()
	// creating a slice DS of type string
	cards := newDeck()

	// to add elements to slide
	cards = append(cards, "King of Diamond")

	// iterate over slice
	// for index, ele := range cards {
	// 	fmt.Println(index, ele)
	// }
	// fmt.Println(cards)

	// calling receiver function
	// cards.print()

	// fmt.Println()
	// deal cards
	// hand, remaining := deal(cards, 3)
	// hand.print()
	// fmt.Println()
	// remaining.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards = newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

// newdeck, print, deal, shuffle, saveToFile, newDeckFromFile
