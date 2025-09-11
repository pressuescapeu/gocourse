package main

import "fmt"

func main() {
	process(9)
	process(-1)
}

func process(input int) {
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")
	if input < 0 {
		fmt.Println("before panic:")
		// so it goes before panic, deferred 2, 1, panic
		// so only after all the stuff done -> panic
		panic("input must be a non-negative number")
		// no code is reachable here, function ends here
	}
	fmt.Println("ok input", input)
}
