// Part A Declare and initialize a variable of type int with the value of 20.
// Display the address of and value of the variable.

// Part B Declare and initialize a pointer variable of type int that points to the last variable you just created.
// Display the address of , value of and the value that the pointer points to.

package main

import "fmt"

func main() {
	// Part A
	i := 20
	fmt.Printf("%v, %v\n", &i, i)

	// Part B
	p := &i
	fmt.Printf("%v, %v, %v\n", &p, p, *p)
}
