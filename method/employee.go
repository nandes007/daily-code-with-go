package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
	address
}

type address struct {
	city    string
	country string
}

func (e employee) details() {
	fmt.Printf("Name: %s\n", e.name)
	fmt.Printf("Age: %d\n", e.age)
}

func (e employee) getSalary() int {
	return e.salary
}

func (e *employee) setNewName(name string) {
	e.name = name
}

// Method on anonymous nested struct fields
func (a address) details() {
	fmt.Printf("City: %s\n", a.city)
	fmt.Printf("Country: %s\n", a.country)
}

// This method will call with chaining
// For method chaining to be possible the methods in the chain should be return the receiver.
// Returnig receiving for the last method in the chain is optional.
func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}
