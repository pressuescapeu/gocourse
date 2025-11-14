package intermediate

import "fmt"

// embedding struct inside another struct - that is how we promote methods to the outer structs

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// won't modify anything

// Area - Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// pointer receiver if the method needs to modify the receiver instance
// or we are trying to avoid copying a large struct

// Scale - Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

func methods() {
	rect := Rectangle{length: 10, width: 9}

	fmt.Println(rect.Area())
	rect.Scale(3)
	fmt.Println(rect.Area())

	num := MyInt(-5)

	fmt.Println(num.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}

	// Area() is a method of Rectangle, but we promoted it to Shape
	fmt.Println(s.Area())
	// can also write s.Rectangle.Area()
}

type MyInt int

// Method on a user-defined type

func (m MyInt) IsPositive() bool {
	return m > 0
}

// here, MyInt - is what this is associated with
// m is the instance of MyInt
// so if we're NOT accessing any value, we can do this

func (MyInt) welcomeMessage() string {
	return "any message slkdj;lkfqw;fe"
}
