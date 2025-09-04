package basics

import "fmt"

func for_loop() {
	// simple iteration over range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// iteration over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d \n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i == 9 {
			break
		}
	}

	rows := 5

	// exercise - ain't nothing fun in it tho
	for i := 1; i <= rows; i++ {
		// spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	// iterating over range
	for i := range 10 {
		fmt.Println(i)
	}
}
