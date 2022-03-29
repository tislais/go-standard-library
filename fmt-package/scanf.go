package main

import "fmt"

func main() {
	var firstname string
	var lastname string
	fmt.Println("What is your name?")
	// Scanf separates values by spaces
	fmt.Scanf("%s %s", &firstname, &lastname)
	fmt.Printf("Hello %s %s! Nice to meet you!\n", firstname, lastname)

	var firstNumber int
	var secondNumber int
	fmt.Println("What two numbers would you like to add?")
	fmt.Scanf("%d %d", &firstNumber, &secondNumber)
	fmt.Printf("Your total is %d\n", firstNumber+secondNumber)
}
