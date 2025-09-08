package basics

import (
	"fmt"
	"strconv"
)

func variadic() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(223, 23232, 3545, 56372, 3546053))

	// slice as a variadic parameter
	numbers := []int{34, 57, 92, 46, 239}
	fmt.Println(sum(numbers...))
}

func sum(nums ...int) (string, int) {
	total := 0
	return_string := "Sum of "
	for _, value := range nums {
		total += value
		return_string += strconv.Itoa(value) + ", "
	}

	return return_string, total
}
