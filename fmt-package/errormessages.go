package main

import (
	"bufio"
	"fmt"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
)

func main() {
	fileName := "before.txt"
	showMessage(INFO, fmt.Sprintf("About to open %s", fileName))
	showMessage(WARNING, "If the file is not present, this application will fail.")
	file, err := os.Open(fileName)
	if err != nil {
		showMessage(ERROR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func showMessage(messagetype messageType, message string) {
	switch messagetype {
	case INFO:
		printMessage := fmt.Sprintf("\nInformation: %s\n", message)
		fmt.Printf(InfoColor, printMessage)
	case WARNING:
		printMessage := fmt.Sprintf("\nWarning: %s\n", message)
		fmt.Printf(WarningColor, printMessage)
	case ERROR:
		printMessage := fmt.Sprintf("\nError: %s\n", message)
		fmt.Printf(ErrorColor, printMessage)
	}
}
