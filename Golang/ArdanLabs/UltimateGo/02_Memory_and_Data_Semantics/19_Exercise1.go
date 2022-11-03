// Exercise 1
// Part A: Declare three variables that are initialized to their zero value
// and three declared with a literal value. Declare variables of type string, int and bool.
// Display the values of those variables.

// Part B: Declare a new variable of type float32 and initialize the variable by converting
// the literal value of Pi (3.14).

package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool

	aa := 1
	bb := "hello"
	cc := true

	fmt.Println(a, b, c)
	fmt.Println(aa, bb, cc)

	pi := float32(3.14159)
	fmt.Println(pi)
}
