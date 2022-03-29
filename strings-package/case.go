package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Hey this is a string!"
	string2 := "hey this is a string!"

	fmt.Println(CompareCaseInsensitive(string1, string2))
}

func CompareCaseInsensitive(a, b string) bool {
	if len(a) == len(b) {
		if strings.ToLower(a) == strings.ToLower(b) {
			return true
		}
	}
	return false
}
