package basics

import "fmt"

func switch_case() {
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("oh wow apple")
	case "banana":
		fmt.Println("oh wow banana")
	default:
		fmt.Println("this ain't no fruit my guy")
	}

	checkType(10)
	checkType(2.22)
	checkType("sdskdksh")
	checkType(false)
}

// gets any value
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
		// can't use fallthrough with data types
		// so like int and int32 are different for compiler
		//fallthrough
	case int32:
		fmt.Print(" ... and it's 32")
	case float64:
		fmt.Println("It's an float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("idk")
	}
}
