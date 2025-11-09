package intermediate

import "fmt"

func fmt_package() {
	// printing functions
	fmt.Print("Hello ") // no newline
	fmt.Print("World!")
	fmt.Print(12, 456) // no space

	fmt.Println("whatever bro") // there is a newline and space, etc

	name := "John"
	age := 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// formatting functions
	// sprintf - takes any number of args, converts each to string, and concatenates all
	s := fmt.Sprint("Hello ", "World!", 123, 456)
	fmt.Print(s)

	// sprintln - same, but adds spaces and newline at the end
	sln := fmt.Sprintln("Hello", "World!", 123, 456, true, false)
	fmt.Print(sln)

	// just a formatted string with placeholders yk
	sf := fmt.Sprintf("Hello, %s. Are you %d?", name, age)
	fmt.Println(sf)

	// scanning functions
	var name_input string
	var age_input int

	fmt.Println("Enter your name and age")
	// you need direct addressing bc otherwise it would be a copy
	fmt.Scan(&name_input, &age_input)
	// waits for the exact format
	//fmt.Scanf("%s %d", &name_input, &age_input)
	// waits for the newline character
	//fmt.Scanln(&name_input, &age_input)
	fmt.Printf("Name: %s, Age: %d\n", name_input, age_input)
}
