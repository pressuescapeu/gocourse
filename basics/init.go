package basics

import "fmt"

func init() {
	fmt.Println("initializing package 1...")
}

func init() {
	fmt.Println("initializing package 2...")
}

func init() {
	fmt.Println("initializing package 3...")
}

func main() {
	fmt.Println("inside the main")
}
