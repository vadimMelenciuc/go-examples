package main

func main() {
	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// cards.saveToFile("my_deck")
	// cards := newDeckFromFile("my_deck")

	cards.print()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}
