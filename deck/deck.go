package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveIt(filename string) error {
	content := []byte(d.toString())

	return ioutil.WriteFile(filename, content, 0644)
}

func loadIt(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPos := r.Intn(len(d) - 1)
		d[newPos], d[index] = d[index], d[newPos]
	}

	return d
}
