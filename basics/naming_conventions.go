package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func naming_conventions() {
	// PascalCase
	// types in go
	// structs, interfaces, enums

	// snake_case
	// user_id, first_name, etc
	// variables, constants, filenames !!!!!

	// UPPERCASE (all caps)
	// for constants

	// mixedCase (starts with a lower case letter)
	// javaScript, htmlDocument, isValid
	// when dealing with external libraries, external language code

	const MAXRETIRES = 5
	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)
}
