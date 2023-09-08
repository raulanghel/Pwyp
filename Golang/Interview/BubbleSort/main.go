package main

import "fmt"

func main() {
	var arr = []int{19, 33, 4, 9, 5, 21, 17}

	sortedArr := bubbleSort(arr)
	fmt.Println(sortedArr)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}

	return arr
}

func swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}
