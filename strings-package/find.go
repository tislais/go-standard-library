package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go language. It's one of my favorites!"

	// searchTerm := "Go"
	// result := strings.Contains(sampleString, searchTerm)
	// fmt.Printf("The sample text includes %s: %t\n", searchTerm, result)

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		// previously .Contains, then .HasPrefix
		result := strings.HasSuffix(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string has the prefex '%s'!\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT have the prefix '%s'!\n", searchTerm)
		}
	} else {
		fmt.Println("You must enter a search term")
	}
}
