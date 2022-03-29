package main

import "fmt"

func main() {
	var age = 42
	var name = "Jeremy"

	// %v tells printf to infer the value from the variable
	// %s string, %d decimal integer
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	var pi float32 = 3.141592
	// %f float format
	fmt.Printf("Pi is %f\n", pi)

	// only show two digits of pi
	fmt.Printf("Pi is %.2f\n", pi)

	// print a row of data
	fmt.Printf("\n|%d|%d|%d|\n", 666, 420, 69)

	// print a row of data with varied numbers
	// pads with 0s to make sure they line up
	// %4.2f will ensure four digits before decimal and 2 after
	fmt.Printf("\n|%4.2f|%4.2f|%4.2f|\n", 666.69, 420.666, 69.42)
	fmt.Printf("|%4.2f|%4.2f|%4.2f|\n", 6.66, 420.69, 690.666)

	// use %7.2f to make it look more like cells in a spreadsheet
	fmt.Printf("\n|%7.2f|%7.2f|%7.2f|\n", 666.69, 420.666, 69.42)
	fmt.Printf("|%7.2f|%7.2f|%7.2f|\n", 6.66, 420.69, 690.666)

	// left justify: %-7.2f
	fmt.Printf("\n|%-7.2f|%-7.2f|%-7.2f|\n", 666.69, 420.666, 69.42)
	fmt.Printf("|%-7.2f|%-7.2f|%-7.2f|\n", 6.66, 420.69, 690.666)

	// you can do this with strings too
	fmt.Printf("\n|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	fmt.Printf("|%-7s|%-7s|%-7s|\n", "a", "ab", "abc")

	// for quoted strings use %q
	fmt.Printf("\n|%-7q|%-7q|%-7q|\n", "foo", "bar", "go")
	fmt.Printf("|%-7q|%-7q|%-7q|\n", "a", "ab", "abc")

	// If you want to format the output without printing it
	output := fmt.Sprintf("\n|%-7s|%-7s|%-7s|\n", "foo", "bar", "go")
	print(output)
}
