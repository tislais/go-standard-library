package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Print(t, "\n")

	// time.Now() returns a struct
	// and you can break it apart
	fmt.Print(t.Year(), "\n")
	fmt.Print(t.Month(), "\n")
	fmt.Print(t.Day(), "\n")

	// break them out into variables
	Year := t.Year()
	Month := t.Month()
	Day := t.Day()

	fmt.Printf("Today is %d/%d/%d\n", Month, Day, Year)
}
