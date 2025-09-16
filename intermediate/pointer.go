package intermediate

import "fmt"

func pointer() {
	// pointer stores another variable's address
	// just like we have types for variables, we have types for pointers
	// USE CASES:
	// - modify the value of a variable indirectly
	// - pass large data structures efficiently between functions
	// - manage memory directly for performance reasons

	// zero value of pointer is - <nil>

	// no support of pointer arithmetic in Go
	// so there is only referencing and dereferencing

	var pointer_instance *int
	var number int = 10
	// referencing operator - &
	pointer_instance = &number
	// prints the address as in 0x14000102020 or some other hex
	// btw every time we execute this .go file, pointer changes
	fmt.Println(pointer_instance)
	// prints the number itself, 10 in our case
	fmt.Println(*pointer_instance) // dereferencing operator - *

	modifyValue(pointer_instance)

	fmt.Println(number)
}

func modifyValue(ptr *int) {
	*ptr += 2
}
