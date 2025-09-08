package basics

import "fmt"

func range_main() {
	message := "Hello World"
	// using range to iterate over string
	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c \n", i, v)
	}
	// for arrays, slices and strings range iterates in order from the first element to the last
	// but for maps it iterates over key-value pairs
	// so order is unspecified
	// for channels it's until the channel is closed

}
