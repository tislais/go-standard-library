package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	string1 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	string2 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id ast laborum."

	// basic compare
	start1 := time.Now()
	if string1 == string2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
	timeTrack(start1, "Basic compare")
	// output: 2022/03/29 12:54:16 Basic compare took 22.914µs

	// using strings compare
	start2 := time.Now()
	if strings.Compare(string1, string2) == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
	timeTrack(start2, "Basic compare")
	// output: 022/03/29 12:54:16 Basic compare took 2.727µs
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
