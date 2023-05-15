package main

import "fmt"

func main() {
	// only length is defined; capacity = length
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))

	// initialize slice; capacity = langth
	b := []int{1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// same length and capacity
	aSlice := make([]int, 4)
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	// add an element
	aSlice = append(aSlice, 5)
	// the capacity is doubled
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	// now add four elements
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	// the capacity is doubled again
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
