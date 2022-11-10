// Declare and initialize constants

package main

import "fmt"

func main() {
	// Constants live within the compiler
	// They have a parallel type system
	// Compiler can perform implicit conversions of untyped constants

	// Untyped constants
	const ui = 12345   // kind: integer
	const uf = 3.14159 // kinf: floating-point

	// Typed Constants still use the constant type system but their precision is restricted
	const ti int = 12345       // type: int
	const tf float32 = 3.14159 // type: float 32

	// constant overflows uint8
	// const myUint uint8 = 1000

	// Constant arithmetic supports different kinds
	// Kind Promotion is used to determine kind in these scenarios

	// Variable answer will be of type float64
	var answer = 3 * 0.333
	fmt.Printf("answer -> Type: %T, Value: %v", answer, answer)

	// Constant third will be of kind floating point.
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

	// Constant zero will be of kind integer.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)

	// This is an example of constant arithmetic between typed and
	// untyped constants. Must have like types to perform math.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
}
