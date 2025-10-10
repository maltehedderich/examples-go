package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// When
	d := newDeck()

	// Then
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Given
	testfile := "_decktesting"
	os.Remove(testfile)

	// When
	deck := newDeck()
	deck.saveToFile(testfile)

	loadedDeck := newDeckFromFile(testfile)

	// Then
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// Cleanup
	os.Remove(testfile)
}
