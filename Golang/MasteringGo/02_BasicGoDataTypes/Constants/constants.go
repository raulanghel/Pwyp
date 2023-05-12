package main

import "fmt"

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 = 1 << iota
		p2_1
		p2_2
		p2_3
		p2_4
		p2_5
		_
		_
		p2_8
	)

	fmt.Println(p2_2)
	fmt.Println(p2_4)
	fmt.Println(p2_8)
}
