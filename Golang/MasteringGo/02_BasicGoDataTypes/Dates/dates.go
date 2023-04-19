package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
	}

	dateString := os.Args[1]

	// is this a date only?
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
	}

	// is this date + time?
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())
	}

	// is this date + time with month represented as a number?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Date: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Hour(), d.Minute())
	}

	// Is this time only?
	d, err = time.Parse("15:04:05", dateString)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), d.Minute(), d.Second())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time: ", t)

	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time: ", duration)
}
