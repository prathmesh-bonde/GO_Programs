package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected last card Four of Club, but got %v", d[len(d)-1])
	}
}

// create file -> save to file -> file created -> load file -> crash
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	/*
		testing saveToDeck and newDeckFromFile:
			del any files in curr working dir with name _decktesting
			create a deck
			save to file _decktesting
			load from file
			assert deck length
			del any files in curr working dir with name _decktesting
	*/

	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
