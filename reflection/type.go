package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	id        int
	firstName string
	lastName  string
}

func main() {
	// employees := make([]employee, 3)
	var employees []employee

	employees = append(employees, employee{0, "Lloyd", "Christmas"})
	employees = append(employees, employee{1, "Harry", "Dunne"})
	employees = append(employees, employee{0, "Sea", "Bass"})

	// name
	fmt.Printf("The name of the type is %v\n", reflect.TypeOf(employees).Name())

	// type
	fmt.Printf("The type is %v\n", reflect.TypeOf(employees))

	// kind
	fmt.Printf("The kind is %v\n", reflect.TypeOf(employees).Kind())

	// value
	fmt.Printf("The value is %v\n", reflect.ValueOf(employees))
}
