package main

import "fmt"

func main() {

	ourString := "\x47\x6f\x20\x69\x73\x20\x41\x77\x65\x73\x6f\x6d\x65\x21"
	anotherString := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println(ourString)
	// output: Go is Awesome!

	// print out the hex values
	for i := 0; i < len(anotherString); i++ {
		fmt.Printf("%x ", anotherString[i])
	}
	fmt.Print("\n")

	// print out quoted byte sequence
	for i := 0; i < len(anotherString); i++ {
		fmt.Printf("%q ", anotherString[i])
	}
	fmt.Print("\n")

	fmt.Println(anotherString + "\n")

	newString := "This is a string!"
	fmt.Print(newString[3])
	// output: 115
	// the byte value at that index.
	fmt.Print(newString[0:5])

	// Strings are readonly slices of bytes
	// the bytes are arbitrary and not normalized in any way
}
