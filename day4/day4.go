package main

import (
	"fmt"
	"unicode/utf8"
)

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge

	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console

}

func (p person) yearPasses() person {
	//Increment the age of the person in here

	return p
}

func main() {
	var x = []byte("Ḁ")
	//var a =  "Ḁ"
	r, _ := utf8.DecodeRune(x)
	fmt.Printf("%#U \n", r)
	fmt.Printf("%.8b \n", r)
	fmt.Printf("%#U", x)
	fmt.Printf("%.8b", x)
}
