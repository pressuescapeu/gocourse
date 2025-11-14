package intermediate

import "fmt"

// see, we're not defining these or methods in the main()
// structs and methods should be defined at package levels - not in main()
// by design in Go
// Go requires types and their associated methods to be declared in the global scope

// main is execution logic!

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

// a struct inside a struct

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

// methods associated with structs
// they are also declared outside the struct

// FROM THE METHODS lecture:
// methods can be associated with any specific type

// RECEIVERS:
// value receiver - like fullName()
// pointer receiver - like incrementAgeByOne()

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// p is instance of pointer of type Person
// pointer make an actual original value / variable accessible
func (p *Person) incrementAgeByOne() {
	p.age++
}

func structs() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       27,
		// these 2 are at the same level of accessing
		address: Address{
			city:    "London",
			country: "UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "323232",
			cell: "23872932938203",
		},
	}

	// BUT we can access home or cell directly
	fmt.Println(p.home)
	// that's what anonymous fields do - promote it to the outer struct!

	// if we omit any field, it will be initialized with default zero value

	p1 := Person{
		firstName: "Ddott",
	}

	p1.address.city = "Compton"
	p1.address.country = "USA"

	p2 := Person{
		firstName: "Ddott",
		address: Address{
			city:    "Compton",
			country: "USA",
		},
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

	// we can also check the equality of structs if they are from the same type, same fields all that
	fmt.Println(p1 == p)
	fmt.Println(p1 == p2)
}
