package intermediate

import "fmt"

//func swap[T any](a, b T) (T, T) {
//	return b, a
//}

// Stack T any - generic type that can be any type in Stack
type Stack[T any] struct {
	elements []T
	// all elements should be of the same type
}

// receiver - s Stack[T]
// parameter - element T - element to be added to the Stack, can be of any type T
// pointer receiver because we change stuff with push
func (s *Stack[T]) push(element T) {
	// accessing the original value, and appending the parameter to the original elements
	s.elements = append(s.elements, element)
}

// bool in the returned element is for whether values was deleted or not
// pop deletes the last element of the stack and returns it
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, true
}

func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if len(s.elements) == 0 {
		fmt.Println("Empty stack!")
	}
	fmt.Println("Stack elements: ")
	for _, element := range s.elements {
		fmt.Print(element)
	}
	fmt.Println()
}

func generics() {
	//x, y := 1, 2
	//x, y = swap(x, y)
	//fmt.Println(x, y)
	//
	//name, surname := "Chris", "Paul"
	//name, surname = swap(name, surname)
	//fmt.Println(name, surname)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.push(4)
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	fmt.Println(intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("Chris")
	stringStack.push("Muhammed")
	stringStack.push("Alexander")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.isEmpty())

}
