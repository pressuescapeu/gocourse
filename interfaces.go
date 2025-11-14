package main

import (
	"fmt"
	"math"
)

// here, you can't export geometry bc it's a lowercase
// we declared geometry interface with 2 unimplemented functions

type geometry interface {
	// IMPORTANT:
	// ALL methods here should be implemented by a struct
	// for that struct to be associated with this interface
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

// now, rect struct implements a function called area()
// area was defined in the geometry interface
func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64
}

// same thing with the rest of the functions

func (c circle) area() float64 {
	// any method or function that needs to be exported should start with an uppercase
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return c.radius * math.Pi * 2
}

func (c circle) diameter() float64 {
	return c.radius * 2
}

// here, measure takes both rect or circle because of the interface geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{
		width:  3,
		height: 4,
	}

	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)

	myPrinter(1, "John", 45.9, true, r)
}

// interface in Go is a way of defining a set of methods that other types must implement
// in order for them to be considered the type that interface is of

// in myPrinter we accept any type, that's why we wrote interface
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
		printType(v)
	}
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type Int")
	case string:
		fmt.Println("Type String")
	case bool:
		fmt.Println("Type Bool")
	default:
		fmt.Println("idk")
	}

}
