package main

func main() {
	newDeck := newDeckFromFile("my_cards")

	newDeck.shuffle()
	newDeck.print()
}
