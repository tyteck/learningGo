package main

import "fmt"

func (p person) print() {
	fmt.Println(p.firstName, " ", p.lastName)
	fmt.Println("email :", p.email)
	fmt.Println("website :", p.website)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}
