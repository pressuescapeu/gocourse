package intermediate

import "fmt"

func closures() {
	//// adder is called only once here
	//// returned anonymous function is stored in sequence and this function
	//// closes over i
	//sequence := adder()
	//// here, in the adder, we create a local variable i
	//
	//// so every time we call sequence we call this anonymous function
	//fmt.Println(sequence())
	//// this anonymous function has access to the local variable i
	//// even if adder already returned
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//
	//// here, the new i is declared
	//sequence2 := adder()
	//// and new i is getting incremented
	//fmt.Println(sequence2())
	//fmt.Println(sequence2())
	//fmt.Println(sequence2())
	//fmt.Println(sequence2())
	// anonymous function with no parameters that returns another function that takes int
	// and returns int as well
	subtracter := func() func(int) int {
		countdown := 101
		return func(x int) int {
			// this "remembers" the countdown variable from the outer scope
			countdown -= x
			return countdown
		}
	}() // () at the end means we called it immediately

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
}

// no argument for adder
// return value is a function that has no parameters and returns int
func adder() func() int {
	i := 0
	fmt.Println("prev value of i:", i)
	return func() int {
		i++
		fmt.Println("adding 1 to i")
		return i
	}
}
