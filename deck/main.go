package main

import "fmt"

func main() {
	cards := newDeck()

	handDeck, _ := deal(cards, 3)

	handDeck.print()

	fmt.Println(toString(handDeck))

	/* fmt.Println("------")

	otherCards.print() */
}
