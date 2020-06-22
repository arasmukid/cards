package main

import "fmt"

// create new data type of 'deck' wich is a slice of string

type deck []string

// create function specific for this deck data type

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}
