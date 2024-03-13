package model

import "fmt"

// This function to call data for exported and non-exported data within same package.
func Call() {
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println(p)
	c := &company{}
	fmt.Println(c)

	//Structure's Method
	fmt.Println(p.Name)
	fmt.Println(p.age)

	//Function
	person2 := GetPerson()
	fmt.Println(person2)
	companyName := getCompanyName()
	fmt.Println(companyName)

	//Variables
	fmt.Println(companyLocation)
	fmt.Println(CompanyName)
}
