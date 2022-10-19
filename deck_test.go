package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 36 {
		t.Errorf("Expected deck lenth 36, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Nine of Clubs" {
		t.Errorf("Expected last card Nine of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("C:\\Users\\Dilanka_700422\\Desktop\\GO\\deck_test.txt")

	d := newDeck()
	d.saveToFile("C:\\Users\\Dilanka_700422\\Desktop\\GO\\deck_test.txt")

	loadedDeck := newDeckFromFile("C:\\Users\\Dilanka_700422\\Desktop\\GO\\deck_test.txt")
	if len(loadedDeck) != 36 {
		t.Errorf("Expected 36 lenth deck, got %v", len(loadedDeck))
	}

	err := os.Remove("C:\\Users\\Dilanka_700422\\Desktop\\GO\\deck_test.txt")
	if err != nil {
		fmt.Printf("Error %v", err)
	}
}
