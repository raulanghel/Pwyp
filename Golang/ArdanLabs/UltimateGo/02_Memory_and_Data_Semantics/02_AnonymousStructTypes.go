package main

import "fmt"

type raul struct {
	flag    bool
	counter int16
	pi      float32
}
type gabriel struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of an anonymous type and set to its zero value
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init using a struct literal
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.14159,
	}

	// Display the value
	fmt.Printf("%+v\n", e2)

	var r raul
	var g gabriel
	// although identical in structure, raul and gabriel are different types
	// r = g
	// we need an explict conversion, on compatible types
	r = raul(g)
	// this is permitted because e2 is a compatible, litteral type
	r = e2
	fmt.Println(r, g)
}
