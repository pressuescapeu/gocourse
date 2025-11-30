package main

import "fmt"

func main() {
	num := 42232320
	// minimum width is fixed to 5 characters
	// if less - populate with 0s
	// 42 would be 00042
	// 42232320 would be 42232320
	fmt.Printf("%05d\n", num)

	// string alignment
	message := "message"
	// we use |    | just to see the alignment
	// fixed width is the minimum fixed width - 10 here
	// right-aligned
	fmt.Printf("|%10s|\n", message)
	// left-aligned
	fmt.Printf("|%-10s|\n", message)

	// string interpolation
	message = "Hello \nWorld!"
	message_with_backticks := `Hello \nWorld!`

	fmt.Println(message)
	fmt.Println(message_with_backticks)

	// we will use backticks for SQL queries
}
