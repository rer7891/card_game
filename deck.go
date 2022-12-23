package main

import "fmt"

// create a new type of deck which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card:= range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// returns a list of playing cards
	cards  := deck{}

	cardSuits :=[]string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues :=[]string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}