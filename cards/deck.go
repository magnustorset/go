package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) deal(handSize int) deck {
	hand := d[:handSize]
	// d=[handSize:]
	return hand
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
