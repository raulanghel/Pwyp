// iota
package main

import "fmt"

func main() {
	const (
		A1 = iota // 0: Starts at 0
		B1 = iota // 1: Increment by 1
		C1 = iota // 2: Incremen by 2
	)
	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0: Starts at 0
		B2        // 1: Increment by 1
		C2        // 2: Increment by 1
	)
	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1: Starts at 0 + 1
		B3            // 2: Increment by 1
		C3            // 3: Increment by 1
	)
	fmt.Println("3:", A3, B3, C3)

	const (
		A4 = 2 // 2: Starts at 2
		B4     // 2:
		C4     // 2:
	)
	fmt.Println("4:", A4, B4, C4)

	const (
		Ldate         = 1 << iota // 1: Shift 1 to the left 0. 0000 0001
		Ltime                     // 2: Shift 1 to the left 1. 0000 0010
		Lmicroseconds             // 4: Shift 1 to the left 2. 0000 0100
		Llongfile                 // 8: Shift 1 to the left 3. 0000 1000
		Lshortfile                // 16: Shift 1 to the left 4. 0001 0000
		LUTC                      // 32: Shift 1 to the left 5. 0010 0000
	)
	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
