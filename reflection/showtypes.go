package main

import (
	"fmt"
	"reflect"
)

func main() {

	type person struct {
		personId  int
		firstName string
		lastName  string
	}

	newPerson := person{0, "Fred", "Flintstone"}
	fmt.Printf("Our person is %s %s with an id of %d\n", newPerson.firstName, newPerson.lastName, newPerson.personId)
	// output: Our person is Fred Flintstone with an id of 0

	// type
	fmt.Printf("The type is %v \n", reflect.TypeOf(newPerson))
	// output: The type is main.person

	// value
	fmt.Printf("The value is %v \n", reflect.ValueOf(newPerson))
	// output: The value is {0 Fred Flintstone}

	// kind
	fmt.Printf("The kind of type it is: %v\n", reflect.ValueOf(newPerson).Kind())
	// output: The kind of type it is: struct

	// name
	fmt.Printf("The name of the type is: %q\n", reflect.TypeOf(newPerson).Name())

	type employee struct {
		personId  int
		firstName string
		lastName  string
	}

	type customer struct {
		customerId int
		firstName  string
		lastName   string
		company    string
	}

	newEmployee := employee{0, "Wilma", "Flintstone"}
	newCustomer := customer{1, "Barney", "Rubble", "Slate Rock and Gravel"}

	addPerson(newCustomer)
	addPerson(newEmployee)

}

func addPerson(p interface{}) bool {
	if reflect.ValueOf(p).Kind() == reflect.Struct {
		v := reflect.ValueOf(p)

		// reflect.TypeOf(p).Name() should return 'employee' or 'customer'
		switch reflect.TypeOf(p).Name() {
		case "employee":
			fmt.Println("It's an employee!")
			empSqlString := "INSERT INTO employees (personId, firstName, lastName) VALUES (?,?,?)"
			fmt.Printf("SQL: %s\n", empSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
		case "customer":
			custSqlString := "INSERT INTO customers (personId, firstName, lastName, company) VALUES (?,?,?,?)"
			fmt.Printf("SQL: %s\n", custSqlString)
			fmt.Printf("Added %v\n", v.Field(1))
			fmt.Println("It's a customer!")
		}
		return true
	} else {
		return false
	}
}
