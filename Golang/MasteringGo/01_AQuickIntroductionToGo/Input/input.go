package main

import "fmt"

func main() {
	// get user input
	fmt.Printf("Please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)
}
