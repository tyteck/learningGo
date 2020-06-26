package main

import "fmt"

func main() {
	cards := newDeck()

	handDeck, _ := deal(cards, 3)

	handDeck.print()

	fmt.Println(handDeck.toString())

	handDeck.saveIt("./foo.txt")

	loadedDeck := loadIt("./foo2.txt")
	loadedDeck.print()

	(loadedDeck.shuffle()).print()
}
