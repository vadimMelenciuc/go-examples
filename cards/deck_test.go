package main

import (
	"os"
	"testing"
)

func TestNewdeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected length of cards 16, not %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected Ace of Spades as first card not: %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected Four of Clubs as first card not: %v", d[len(d)-1])
	}
}

func TestSaveDeckToFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected length of cards 16, not %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
