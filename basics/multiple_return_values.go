package basics

import (
	"errors"
	"fmt"
)

func functions_multiple_main() {
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d Remainder: %d \n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

// before named:

//func divide(a, b int) (int, int) {
//	quotient := a / b
//	remainder := a % b
//	return quotient, remainder
//}

// after named:
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// already needs to know what to return
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b < a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
