package main

import (
	"fmt"
	"math"
)

// Methods on Structs
func methodOnStruct() {
	fmt.Println("\nMethod on struct")
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.details()
	fmt.Printf("Salary: %d\n", emp.getSalary())
}

func changeReceiverFieldFromMethod() {
	fmt.Println("\nSet field from method")
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.setNewName("John")
	fmt.Printf("Name: %s\n", emp.name)

	(&emp).setNewName("Mike")
	fmt.Printf("Name: %s\n", emp.name)

	emp2 := &employee{name: "Sam", age: 31, salary: 2000}
	emp2.setNewName("John")
	fmt.Printf("Name: %s\n", emp2.name)
}

func callMethodUsingStructName() {
	fmt.Println("\nCalling method using struct name")
	emp := employee{name: "Sam", age: 20, salary: 2000}
	employee.details(emp)

	(*employee).setNewName(&emp, "John")
	fmt.Printf("Name: %s\n", emp.name)
}

// Method on anonymous nested struct fields
func methodAnonymousNested() {
	fmt.Println("\nMethod on anonymous nested struct fields")
	address := address{city: "London", country: "UK"}
	emp := employee{name: "Sam", age: 31, salary: 3000, address: address}
	emp.details()
	emp.address.details()
}

// Method Chaining
func methodChaining() {
	fmt.Println("\nMethod Chaining")
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}

//Method on Non-Struct Types
//Method can also be defined on a non-struct custom type.
//Non-struct custom types can be created through type definition.

// Custom type
type myFloat float64

func (m myFloat) ceil() float64 {
	return math.Ceil(float64(m))
}

func methodOnNonStructTypes() {
	fmt.Println("\nMethod on non struct types")
	num := myFloat(1.4)
	fmt.Println(num.ceil())
}
