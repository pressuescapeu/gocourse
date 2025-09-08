package basics

import "fmt"

func functions() {
	// func <name>(parameters list) returnType {} - this is if the returned value is single
	// private function - lower case
	// public function - upper case, like Println() or Equal()
	//sum := add(1, 2)
	//fmt.Println(add(1, 2))
	//
	//// anonymous function - functions without a name, defined in-line
	//out := func() {
	//	fmt.Println("anon lol")
	//}
	//out()
	//
	//operation := add
	//result := operation(100, 23232)
	//fmt.Println(result)

	// passing a function as an argument
	result := applyOperation(5, 4, add)
	fmt.Println(result)
	// creating a function
	times2 := createMultiplier(2)
	fmt.Println(times2(6))
}

func add(a, b int) int {
	return a + b
}

// function that takes a function as argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// function that returns a function
// takes factor as argument
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
