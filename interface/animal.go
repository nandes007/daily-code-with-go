package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type animalPointer interface {
	breathePointer()
	walkPointer()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

// interface are implemented implicitly
type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

func (l *lion) breathePointer() {
	fmt.Println("Lion brathes")
}

func (l *lion) walkPointer() {
	fmt.Println("Lion Walk")
}

// Non-struct custom type implementing an interface
type cat string

func (c cat) breathe() {
	fmt.Println("Cat Breathes")
}

func (c cat) walk() {
	fmt.Println("Cat Walk")
}

// implement multiple interface
type mamal interface {
	feed()
}

func (l lion) feed() {
	fmt.Println("Lion feeds young")
}

// Embedding interface in a struct
type pet1 struct {
	a    animal
	name string
}

type pet2 struct {
	animal
	name string
}

// Type Assertion
func print(a animal) {
	l := a.(lion)
	fmt.Printf("Age: %d\n", l)
}

// Second way to get underlying value using type assertion
func printTwo(a animal) {
	l, ok := a.(lion)
	if ok {
		fmt.Println(l)
	} else {
		fmt.Println("a is not type lion")
	}

	d, ok := a.(dog)
	if ok {
		fmt.Println(d)
	} else {
		fmt.Println("d is not type dog")
	}
}

func printSwitch(a animal) {
	switch v := a.(type) {
	case lion:
		fmt.Println("Type: lion")
	case dog:
		fmt.Println("Type: dog")
	default:
		fmt.Printf("Unknown type %T", v)
	}
}
