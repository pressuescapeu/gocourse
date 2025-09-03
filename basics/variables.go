package basics

import "fmt"

var firstName = "Billie"

func variables() {
	// uninitialized variable
	//var age int
	// if you initialize then specifying type is optional
	//var name = "Bobby"

	// gofer sign
	// this can only be used within functions like local, in block
	//count := 122
	//lastName := "Lee"
	var firstName = "Jean"

	fmt.Println(firstName)
	// default values
	// default values
	// numeric - 0
	// boolean - false
	// string - "" or like an empty string
	// Pointers, slices, maps, functions, structs - nil

	// variable live only within scope
	// so they're destroyed after for memory management

	// variable have a block scope
	// like they're available only within their block
}
