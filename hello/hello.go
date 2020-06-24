package hello

func main() {
	catNames := cats{"felix", "jaguar"}
	catNames = append(catNames, "the cat")

	catNames.show()
}
