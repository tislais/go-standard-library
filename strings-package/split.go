package main

import (
	"fmt"
	"strings"
)

func main() {
	ourString := "This is a sentence."
	stringCollection := strings.Split(ourString, "|")
	for i := range stringCollection {
		fmt.Println(stringCollection[i])
	}
	// output:
	// This
	// is
	// a
	// sentence.

	ourString2 := "This is a new sentence.|This is a banana.|Banan a jamma??"
	// SplitAfter splits on the character, but still includes it
	stringCollection2 := strings.SplitAfter(ourString2, "|")
	for i := range stringCollection2 {
		fmt.Println(stringCollection2[i])
	}

	ourString3 := "This is a banana string!"
	// SplitAfterN splits on the character, but stops when its reached the number (i think?)
	stringCollection3 := strings.SplitAfterN(ourString3, " ", 2)
	for i := range stringCollection3 {
		fmt.Println(stringCollection3[i])
	}
	// output:
	// This
	// is a banana string!
}
