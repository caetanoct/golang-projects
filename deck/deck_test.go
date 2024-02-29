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
	if d[0] != "ace of spades" {
		t.Errorf("Expected first card to be ace of spades but got %v", d[0])
	}
}

func TestSaveFileAndLoadFile(t *testing.T) {
	filename := "deck_test"
	os.Remove(filename)
	deck := newDeck()
	deck.saveToFile(filename)
	loaded, err := newDeckFromFile(filename)
	if err == nil {
		if len(loaded) != len(deck) {
			t.Errorf("Expected decks to have same size")
		}
	} else {
		t.Errorf("File error: %v", err)
	}
	os.Remove(filename)
}
