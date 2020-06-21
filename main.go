package main

import "fmt"

func main() {
	// card := "Ace of Spades"
	// card := newCard()
	//card = "Five of Diamonds"
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
