package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

// we have p person and not just person because we are accessing the fields with the introduce() function
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

type Employee struct {
	person // anonymous field
	// so if it was not anonymous, we couldn't do instance.name
	empId  string
	salary float64
}

// methods can be overwritten by redefining them in the outer struct

func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old, I have an employee ID of %s and I earn %.2f.\n", e.name, e.age, e.empId, e.salary)
}

func struct_embedding() {
	emp := Employee{
		person: person{name: "John", age: 30},
		empId:  "E001", salary: 50000,
	}

	// employee can directly access fields of person as if they were fields of Employee by default
	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	fmt.Println(emp.salary)

	// since person was embedded in Employee struct, we can directly access introduce from an instance of employee
	// but then we override the introduce() that was in the person, and use the one that is in Employee
	emp.introduce()
}
