package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp := employee{name: "Sam", age: 31, salary: 20000}

	//Accessing a struct field
	n := emp.name
	fmt.Printf("Current name is: %s\n", n)

	//Assigning a new value
	emp.name = "John"
	fmt.Printf("New name is: %s\n", emp.name)
}
