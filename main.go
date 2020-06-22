package main

import "fmt"

func ()  {
	// card := "Ace of Spades" //inisialisai dan deklarasi
	card := newCard() //ambil value variable dari fungsi
	
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
 