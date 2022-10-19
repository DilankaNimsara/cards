package main

import "fmt"

func main() {
	//cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//cards.saveToFile("C:\\Users\\Dilanka_700422\\Desktop\\GO\\my_card.txt")
	deckFromFile := newDeckFromFile("C:\\Users\\Dilanka_700422\\Desktop\\GO\\my_card.txt")
	deckFromFile.print()
	fmt.Println("#######################")
	deckFromFile.shuffle()
	deckFromFile.print()
}
