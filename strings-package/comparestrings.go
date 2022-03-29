package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "this is a string!"
	string2 := "this is a string!"

	if strings.Compare(string1, string2) == 0 {
		fmt.Println("The strings are identical!")
	} else {
		fmt.Println("The strings do not match.")
	}

	stooges := []string{"Larry", "Curly", "Moe"}

	for _, stooge := range stooges {
		fmt.Println(strings.Compare("Larry", stooge))
	}
	// output
	// 0 (match)
	// 1 (not a match)
	// -1 (less than, in lexigraphical order)
}
