package basics

import "fmt"

func defer_main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("defered i value:", i) // value of i was 10 here
	// so defer treats it like a stack
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	i++
	fmt.Println("normal")
	fmt.Println(i)
}
