package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("names.csv")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		fmt.Println("-- new record --")
		for i := range items {
			fmt.Println(items[i])
		}
		fmt.Print("\n")
	}
}
