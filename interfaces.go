package main

import "fmt"

type Animal interface { // "An interface type is defined as a set of method signatures"
	Species() string
}

type Cat struct {
	Name  string
	Color string
}

func main() {
	boris := Cat{"Boris", "gray"}
	PrintInfo(&boris)
}

func PrintInfo(a Animal) {
	fmt.Println(a.Species())
}

func (c *Cat) Species() string { // "A type implements an interface by implementing its methods"
	return "Felis catus"
}
