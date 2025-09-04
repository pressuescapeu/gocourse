package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess")
		// if we use guess then the copy of guess is created
		// meaning that the original guess in line 22 is not updated
		// but if we use the address of the guess, e.g. &guess
		fmt.Scanln(&guess)

		// now we check if the guess is correct
		if guess == target {
			fmt.Println("congrats bro")
			break
		} else if guess > target {
			fmt.Println("too high")
		} else {
			fmt.Println("too low")
		}
	}
}
