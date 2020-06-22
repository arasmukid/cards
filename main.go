package main

import "fmt"

func main() {
	// card := "Ace of Spades" //inisialisai dan deklarasi
	// card := newCard() //ambil value variable dari fungsi
	cards := []string{"Ace of Diamonds", newCard()} //inisialisasi slice dg type string dg 2 parameter value, satu string dan fungsi

	cards = append(cards, "six of spades") //append data ke dalam slice yg telah di-inisialisasi

	// fmt.Println(cards) // print isi card ke layar

	for i, card := range cards { //loop slice cards ke variable baru card dan print ke layar dg indexnya
		fmt.Println(i, card) //indexs i harus diprint/ digunakan juga utk menghindari error
	}
}

func newCard() string {
	return "Five of Diamonds"

}
