package main

import (
	"fmt"
)

func main() {
	var evenOrOdd string
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		evenOrOdd = "odd"

		if n%2 == 0 {
			evenOrOdd = "even"
		}
		fmt.Println(n, " is ", evenOrOdd)
	}
}
