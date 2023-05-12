package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need input in __02 January 2006 15:04 MST__ format")
		return
	}

	input := os.Args[1]
	now, err := time.Parse("02 January 2006 15:04 MST", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// local time
	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current location: %s\n", now.In(loc))

	// NY
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York: %s\n", now.In(loc))

	// Bucuresti
	loc, _ = time.LoadLocation("Europe/Bucharest")
	fmt.Printf("Iasi: %s\n", now.In(loc))

	// Tokyo
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo: %s\n", now.In(loc))
}
