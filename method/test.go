package main

import "fmt"

func Test() {
	fmt.Println("\nExported and Unexported Struct, Field and Method")
	//STRUCTUR IDENTIFIER
	p := &Person{
		Name: "test",
		age:  21,
	}

	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//STRUCTURE'S FIELD
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//STRUCTURE'S METHOD
	fmt.Println(p.GetAge())
	fmt.Println(p.getName())
}
