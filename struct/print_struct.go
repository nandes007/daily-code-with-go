package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	name   string
	age    int
	salary int
}

type employee2 struct {
	Name   string
	Age    int
	salary int
}

func main() {
	//There are 2 ways to print all struct variables including all its key and values:
	//- Using fmt package
	//- Print struct in JSON form using json/encoding package

	emp := employee{name: "Sam", age: 31, salary: 2000}

	fmt.Println("Using fmt packge")
	fmt.Printf("Emp: %v\n", emp)  //will print value only
	fmt.Printf("Emp: %+v\n", emp) //wil print both field and value.
	fmt.Printf("Emp: %#v\n", emp)
	fmt.Println(emp)

	fmt.Println("\nPrint struct in JSON form")
	//To be noted that both Marshal and Marshalindent function can only
	//access the exported fields of a struct,
	//which means that only capitalized fields can be accessed and encoded in JSON form.
	emp2 := employee2{Name: "Sam", Age: 30, salary: 2000}
	//Marshal
	empJSON, err := json.Marshal(emp2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Marshal function output %s\n", string(empJSON))

	//MarshalIndent
	empJSON, err = json.MarshalIndent(emp2, "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent function output %s\n", string(empJSON))
}
