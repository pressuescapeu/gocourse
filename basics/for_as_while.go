package basics

import "fmt"

func for_as_while() {
	i := 1
	// just condition
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop:
	//for {
	//	fmt.Println("inf")
	//}
}
