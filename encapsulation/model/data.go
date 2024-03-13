package model

import "fmt"

var (
	CompanyName     = "test"
	companyLocation = "somecity"
)

// Person struct
type Person struct {
	Name string
	age  int
}

// GetAge of person
func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) getName() string {
	return p.Name
}

type company struct {
}

// GetPerson get the person object
func GetPerson() *Person {
	p := &Person{
		Name: "test",
		age:  21,
	}
	fmt.Println("Model Package:")
	fmt.Println(p.Name)
	fmt.Println(p.age)
	return p
}

func getCompanyName() string {
	return CompanyName
}
