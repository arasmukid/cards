package main

import "fmt"

// create new data type of 'deck' wich is a slice of string

type deck []string

// create function specific for this deck data type

func (d deck) print() { // (d deck) adalah referrence variable ke data typenya shg si variable
	// bisa mengakses print method,
	for i, card := range d {
		fmt.Println(i, card)
	}

}
