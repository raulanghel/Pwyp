package main

import "fmt"

// example represents a type with different fields
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of type exmple and set it to its zero value
	var e1 example

	// Display the value
	fmt.Printf("%v\n", e1)
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init it using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.12159,
	}

	// Display the field values
	fmt.Println("flag: ", e2.flag)
	fmt.Println("counter: ", e2.counter)
	fmt.Println("pi: ", e2.pi)
}
