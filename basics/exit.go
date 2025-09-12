package basics

import (
	"fmt"
	"os"
)

func exit() {
	// os.Exit() exits immediately without any cleanup
	// any deferred functions won't be executed
	// 0 - successful completion
	// bypasses defer, recover, panic mechanisms

	// this won't be executed either
	defer fmt.Println("anything")
	// this will be executed
	fmt.Println("Starting the main function")
	// exit with status of code of 1
	os.Exit(1)
	// this will never be executed
	fmt.Println("anything")
}
