package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}

func myPrint(start, finish int) {
	for i := start; i < finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}
