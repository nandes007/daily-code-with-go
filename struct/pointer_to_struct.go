package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	//There are two ways of creating a pointer to the struct
	//- Using the & operator
	//- Using the new keyword

	fmt.Println("Using & operator")
	emp := employee{name: "Sam", age: 31, salary: 2000}
	empP := &emp
	fmt.Printf("Emp: %+v\n", empP)
	empP = &employee{name: "John", age: 30, salary: 3000}
	fmt.Printf("Emp: %+v\n", empP)

	fmt.Println("\nUsing new keyword")
	empP2 := new(employee)
	fmt.Printf("Emp Pointer Address: %p\n", empP2)
	fmt.Printf("Emp Pointer: %+v\n", empP2)
	fmt.Printf("Emp Value: %+v\n", *empP2)
}
