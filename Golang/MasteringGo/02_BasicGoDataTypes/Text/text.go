package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First character: ", string(aString[0]))

	// Runes
	r := '€'
	fmt.Println("As an integer: ", r)
	// Concert rune to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	// print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
	// print an existing string as characters
	for _, v := range aString {
		fmt.Printf("%c ", v)
	}
	fmt.Println()

}
