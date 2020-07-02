package main

import "fmt"

type contactInfo struct {
	email   string
	website string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	john := person{firstName: "john", lastName: "connor"}
	fmt.Println(john)
	fmt.Printf("%+v", john)
	fmt.Println("")

	var alex person
	alex.firstName = "alex"
	alex.lastName = "mortimer"
	alex.contactInfo.email = "alex@mortimer.com"

	fmt.Println(alex)

	jim := person{
		firstName: "jim",
		lastName:  "conney",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			website: "http://jimconney.com",
		},
	}

	fmt.Println("------")
	(&jim).updateFirstName("jimmy")

	jim.print()
}
