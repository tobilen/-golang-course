package main

import "fmt"

func main() {
	deck := newDeckFromFile("test.data")

	hand, deck := deck.deal(3)

	hand.print()
	fmt.Println("---")
	deck.print()
	fmt.Println("---")
	deck.saveToFile("test.data")
}