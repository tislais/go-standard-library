package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		searchTerm := os.Args[1]

		file, _ := os.Open("log.txt")
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			result := strings.Contains(line, searchTerm)
			if result {
				fmt.Println(line)
			}
		}
	}
}

// we can even pipe these searches into new files
// try this command:
// $ go run strings-package/logparser.go ERROR: > error.txt
