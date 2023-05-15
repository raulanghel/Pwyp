package main

import "fmt"

func main() {
	// Create an empty slice
	aSlice := []float64{}
	// Length and Capacity are 0 because the slice is empty
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// Add elements to a slice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// A slice with length 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// Now we will need to use append
	t = append(t, -5)
	fmt.Println(t)

	// A 2D slice
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	// Visiting all the elements of a twoD slice with a double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4, 5}
	make2D[1] = []int{-1, -2, -3, -4, -5}
	fmt.Println(make2D)
}
