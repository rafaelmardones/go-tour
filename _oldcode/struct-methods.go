package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m myStruct) getFirstName() string { // method = function with a receiver argument
	return m.FirstName
}

func (m *myStruct) forgetFirstName() { // pointer receiver!
	m.FirstName = "unknown"
}

func main() {
	myVar := myStruct{ // shorthand
		FirstName: "Rafael",
	}

	fmt.Println("first name is", myVar.FirstName)
	fmt.Println("first name is", myVar.getFirstName())
	myVar.forgetFirstName()
	fmt.Println("first name is", myVar.FirstName)
	fmt.Println("first name is", myVar.getFirstName())
}
