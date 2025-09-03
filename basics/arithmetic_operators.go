package basics

import (
	"fmt"
	"math"
)

func arithmetic_operators() {
	// variable declaration
	var a, b int = 10, 3
	var result int
	var flo_res float64
	var smallFloat float64 = 1.0e-323

	result = a + b
	fmt.Println(result)

	result = a - b
	fmt.Println(result)

	result = a * b
	fmt.Println(result)
	// division between 2 integers is also an integer
	// any fraction part is discarded
	result = a / b
	fmt.Println(result)

	// why does this work? bc 22 is still an untyped constant, like not yet int
	// so since 7.0 is float, whole 22 / 7.0 is float
	flo_res = 22 / 7.0
	fmt.Println(flo_res)
	// same shit
	flo_res = 22.0 / 7
	fmt.Println(flo_res)

	// but like variable a / some float wouldn't work
	// bc operands must be the same type

	result = a % b
	fmt.Println(result)

	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)
	maxInt++
	fmt.Println(maxInt)

	// unsigned: 0 to 18446744073709551615
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)
	uMaxInt++
	fmt.Println(uMaxInt)
	uMaxInt++
	fmt.Println(uMaxInt)

	fmt.Println(smallFloat)
	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)
}
