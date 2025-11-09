package main

import "fmt"

// see, we're not defining these or methods in the main()
// structs and methods should be defined at package levels - not in main()
// by design in Go
// Go requires types and their associated methods to be declared in the global scope

// main is execution logic!

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       27,
	}

	// if we omit any field, it will be initialized with default zero value

	p1 := Person{
		firstName: "Ddott",
	}

	// see, this prints out:
	// {John Doe 27}
	// {Ddott  0}
	// bc the rest are default 0 values

	fmt.Println(p)
	fmt.Println(p1)

	// let's change the age now
	p.incrementAgeByOne()
	fmt.Println(p)

	fmt.Println(p.fullName())

	// anonymous structs
	user := struct {
		email    string
		username string
	}{ // we pass the values immediately
		username: "ddott11",
		email:    "dotdotdotdot@dotcom",
	}

	fmt.Println(user)
}

// methods associatted with structs
// they are also declared outside the struct

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// p is instance of pointer of type Person
// pointer make an actual original value / variable accessible
func (p *Person) incrementAgeByOne() {
	p.age++
}
