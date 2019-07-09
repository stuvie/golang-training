package main

import "fmt"

var deckSize int

// A deck of cards
func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()

	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	loaded := newDeckFromFile("my_cards")
	loaded.shuffle()
	loaded.print()

	newDeckFromFile("oops")

	// var deckSize int
	deckSize = 52
	fmt.Println(deckSize)
}
