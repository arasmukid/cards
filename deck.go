package main

import "fmt"

// create new data type of 'deck' wich is a slice of string

type deck []string

// create newDeck functin semacam constructor

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// create function specific for this deck data type

func (d deck) print() { // (d deck) adalah referrence variable ke data typenya shg si variable
	// bisa mengakses print method
	for i, card := range d {
		fmt.Println(i, card)
	}

}
