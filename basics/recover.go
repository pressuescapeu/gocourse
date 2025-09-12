package basics

import "fmt"

func recover_main() {
	// main calls process
	// if process doesn't recover - program just crashes
	process()
	// if program crashes - we never get to this point in code
	fmt.Println("Returned from process")
}

func process() {
	// we run this function at the end of the process, no matter what happens
	defer func() { // this is an anonymous, deferred function
		// btw this if is basically:
		// r := recover()
		// if r != nil { and the rest of the code
		if r := recover(); r != nil {
			// here, recover() checks if there was a panic, and returns the panic value as r
			// so if a panic happened, we catch it and print it instead of letting the program die
			fmt.Println("Recovered:", r)
		}
	}() // we put () as to call this function immediately when defer runs

	fmt.Println("Start Process")
	panic("Something went wrong")
	// this never runs, as execution stops at panic, and we go to the deferred anonymous function
	fmt.Println("End Process")
}
