package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	arguments := os.Args
	for _, k := range arguments[1:] {
		// is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		// is it a float?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		// then it is invalid
		invalid = append(invalid, k)
	}

	fmt.Println("Ints:", nInts)
	fmt.Println("Floats:", nFloats)
	fmt.Println("Invalids:", len(invalid))
}
