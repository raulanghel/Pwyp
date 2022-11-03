// Sharing Data
package main

import "fmt"

func main() {
	count := 10

	// Display the "value of" and "address of" variable count
	fmt.Printf("count: \tValue[%v] \tAddress[%v]\n", count, &count)

	// Pass the Address Of count
	increment(&count)

	fmt.Printf("count: \tValue[%v] \tAddress[%v]\n", count, &count)
}

func increment(inc *int) {
	// Increment the "Value Of" count that the pointer points to
	*inc++
	fmt.Printf("inc: \tValue[%v] \tAddress[%v]\n", inc, &inc)
}
