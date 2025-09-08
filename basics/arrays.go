package basics

import "fmt"

func arrays() {
	var numbers [5]int // blank array with 0s
	fmt.Println(numbers)

	numbers[0] = 1
	numbers[4] = 5
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Orange", "Banana", "Grapes"}

	// iterate over array using RANGE
	// %d for int
	// %s for string
	// little c shit over here huh
	for index, value := range fruits {
		fmt.Printf("index: %d value: %s \n", index, value)
	}
	// ts would've been different like this
	for i := 0; i < len(fruits); i++ {
		fmt.Println("index:", i, "value:", fruits[i])

	}

	originalArray := [3]int{1, 2, 3}
	// so this is an assignment
	// here, the whole orignalArray thing is considered as a value ([1, 2, 3])
	// we get the entire block of memory that is the originalArray's block of memory
	// and copy it into copiedArray
	copiedArray := originalArray
	// after this, these 2 live a separate life

	copiedArray[0] = 101

	fmt.Println("Original array: ", originalArray)
	fmt.Println("Copied array: ", copiedArray)

	// if I wanted to make them both change with any operation
	// I'd create a slice
	original := []int{1, 2, 3} // see, ts is dynamic bc we ain't got no predefined length
	copied := original         // they share the same backing array

	copied[0] = 101 // changes both

	fmt.Println(original) // [101 2 3]
	fmt.Println(copied)   // [101 2 3]

	// we might wanna go pointers on this one
	orig := [3]int{1, 2, 3}
	// declare a POINTER to an array of 3 integers
	var copy *[3]int

	// copy points to orig's memory
	copy = &orig
	// change through the pointer
	copy[0] = 10 // same as if you'd dereference and then did it like
	// (*copy)[0] = 10
	// Go automatically does that (yay)
	fmt.Println(orig)
	// dereference the pointer
	fmt.Println(*copy)

	// UNDERSCORE
	// no errors here bc we don't want to use whatever is in the underscore
	a, _ := someFunction()
	fmt.Println(a)

	// comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 4}

	fmt.Println("Arrays are equal?", array1 == array2)

	var matrix = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
}

func someFunction() (int, int) {
	return 1, 2
}
