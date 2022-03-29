package main

import "fmt"

type point struct {
	x, y int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := point{2, 3}
	fmt.Printf("%v\n", p)

	newPerson := Person{"Jeremy", "Morgan", 42}

	// type
	fmt.Printf("%T\n", newPerson)

	// bool
	fmt.Printf("%t\n", false)

	// integer
	fmt.Printf("%d\n", 4567)

	// binary
	fmt.Printf("%b\n", 4567)

	// ascii character from number
	fmt.Printf("%c\n", 42)

	// hex
	fmt.Printf("%x\n", 4567)
}
