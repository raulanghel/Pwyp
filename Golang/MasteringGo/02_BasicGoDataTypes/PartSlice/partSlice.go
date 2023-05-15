package main

import "fmt"

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	// first five elements
	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	// last two elements
	fmt.Println(aSlice[l-2 : l])
	fmt.Println(aSlice[l-2:])

	// first 5 elements
	t := aSlice[0:5:10]
	fmt.Println(t)
	fmt.Println(len(t), cap(t))

	// elements at indexes 2, 3, 4
	t = aSlice[2:5:10]
	fmt.Println(t)
	fmt.Println(len(t), cap(t))

	// elements at indexes 0,1,2,3,4
	// new capacity will be 6-0
	t = aSlice[:5:6]
	fmt.Println(len(t), cap(t))
}
