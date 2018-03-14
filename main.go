package main

import "fmt"

func main() {
	deck := newDeck()
	// deck := newDeckFromFile("test.data")
	deck.shuffle()

	hand, deck := deck.deal(3)

	hand.print()
	fmt.Println("---")
	deck.print()
	fmt.Println("---")
	deck.saveToFile("test.data")
}