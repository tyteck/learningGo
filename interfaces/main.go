package main

import "fmt"

type langBot interface {
	getGreeting() string
}

type englishBot struct{}
type frenchBot struct{}

func main() {
	eb := englishBot{}
	fb := frenchBot{}

	printGreeting(eb)
	printGreeting(fb)
}

func (englishBot) getGreeting() string {
	return "Hello there"
}

func (frenchBot) getGreeting() string {
	return "Salut la dedans"
}

func printGreeting(lb langBot) {
	fmt.Println(lb.getGreeting())
}
