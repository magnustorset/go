package main

func main() {
	cards := newDeck()

	hand := cards.deal(5)
	hand := cards.deal(2)
	hand.print()
}

func newCard() string {
	return "Five of Diamonds"
}
