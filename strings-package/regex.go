package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString := "This is a song. There are many strings like it, but this one is mine"

	// when doing a regex search you compile the search when initializing
	r, _ := regexp.Compile(`s(\w[a-z]+)g\b`)
	fmt.Println(r.MatchString(sampleString))

	fmt.Println(r.FindAllString(sampleString, -1))

	fmt.Println(r.FindStringIndex(sampleString))

	newText := r.ReplaceAllString(sampleString, "apple")
	fmt.Println(newText)
}
