package main

// Exported and non exported struct, field struct and method struct
// struct, field and method with capital is exported otherwise not.
type Person struct {
	Name string
	age  int
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) getName() string {
	return p.Name
}

type company struct {
}
