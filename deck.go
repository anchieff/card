package main

import "fmt"

// Create a new type 'deck'

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
