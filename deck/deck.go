package main

import "fmt"

type deck []string

func deal(d deck, number int) (deck, deck) {
	return d[:number], d[number:]
}

func newDeck() deck {

	result := deck{}
	colors := []string{"coeur", "carreau", "pique", "trèfle"}
	values := []string{"as", "roi", "dame", "valet"}

	for _, color := range colors {
		for _, value := range values {
			result = append(result, value+" de "+color)
		}
	}
	return result
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println("-", card)
	}
}

func (d deck) toString string {
	var result string
	for i, card := range d {
		if i > 0 {
			result = result + ' '
		}

		result = result + card
	}
	return result
}
