package main

import "fmt"

func main() {
	var name string
	fmt.Println("What is your name?")
	// %q will only read in a quoted string
	inputs, _ := fmt.Scanf("%q", &name)

	switch inputs {
	case 0:
		fmt.Printf("You must enter a name!\n")
	case 1:
		fmt.Printf("Hello %q! You input %d\n", name, inputs)
	}
}
