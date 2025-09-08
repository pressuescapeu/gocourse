package basics

import (
	"fmt"
	"slices"
)

func slices_main() {
	slices_basics()
	// two dimensional slice
	twoD := make([][]int, 3) // outer capacity is 3
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d and in inner slice of index of %d\n", i+j, i, j)
		}
	}
	fmt.Println(twoD)

	// slice operator
	// looks like this -- slice[low:high]
	slice1 := []int{3, 5, 7, 1, 0, 9}
	slice2 := slice1[2:4]
	fmt.Println(slice2)
	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The len of slice is", len(slice2))
	// somewhere here we got a brand new array
	slice2 = append(slice2, 4, 5, 6, 7, 8)
	fmt.Println(slice2)
	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The len of slice is", len(slice2))
}

func slices_basics() {
	// empty slice, nil slice, length 0, capacity 0
	var numbers []int
	// empty but not nil slice, length 0
	nonNilNumbers := []int{}
	// slice with predefined values and var
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{1, 2, 3}
	// length 5, capacity 5, all zeroes
	slice := make([]int, 5)
	a := [5]int{2, 3, 4, 5, 6}
	// slice from an array - includes elements 1, 2, 3 but NOT 4 and NOT 0
	s := a[1:4] // 3, 4, 5
	fmt.Println(s)
	// append function - add to a slice
	s = append(s, 6, 7)
	sCopy := make([]int, len(s))
	copy(sCopy, s)
	s[2] = 100
	fmt.Println(s)
	fmt.Println(sCopy)
	fmt.Println(numbers)
	fmt.Println(nonNilNumbers)
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println(slice)
	fmt.Println(a)

	for i, v := range s {
		fmt.Println(i, v)
	}

	// check if slices are equal
	if slices.Equal(s, sCopy) {
		fmt.Println("s is equal to sCopy")
	} else {
		fmt.Println("nah")
	}
}
