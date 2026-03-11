package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")
	newCards.print()
}
