package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "Never trust a Programmer who carries a screwdriver\n"
	fmt.Println("Before: " + sampleString)

	strLowerCase := strings.ToLower(sampleString)
	fmt.Println("Lower case: " + strLowerCase)

	strUpperCase := strings.ToUpper(sampleString)
	fmt.Println("Upper case: " + strUpperCase)

	strTitleCase := strings.Title(sampleString)
	fmt.Println("Title case: " + strTitleCase)

	newString := "welcome to the dollhouse\n"
	fmt.Println(strings.Title(newString))
	// output: Welcome To The Dollhouse

	fmt.Println(properTitle(newString))
	// output: Welcome to the Dollhouse

}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallWords := " a an on the to "

	for i, w := range words {
		if strings.Contains(smallWords, " "+w+" ") {
			words[i] = w
		} else {
			words[i] = strings.Title(w)
		}
	}
	return strings.Join(words, " ")
}
