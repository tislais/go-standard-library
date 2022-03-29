package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "    This is our text    "
	fmt.Printf("%q\n", sampleString)

	newString := strings.TrimSpace(sampleString)
	fmt.Printf("%q\n", newString)

	newString = strings.TrimLeft(sampleString, " ")
	fmt.Printf("%q\n", newString)

	urlString := "https://www.pluralsight.com"
	domainName := strings.TrimPrefix(urlString, "https://")
	fmt.Printf("%q\n", domainName)

	fileString := "error.txt"
	fileName := strings.TrimSuffix(fileString, ".txt")
	fmt.Printf("%q\n", fileName)

}
