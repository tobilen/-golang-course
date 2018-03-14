package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck length of 16,  but got %v", len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected first card to be One of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card to be Ace of Clubs, but got %v", d[0])
	}
}