package main

import "fmt"

func main() {
	//var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
	}
	//fmt.Println(colors)

	print(colors)
}

func print(m map[string]string) {

	for key, value := range m {
		fmt.Println(key, "is", value)
	}

}
