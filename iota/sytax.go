package main

import "fmt"

// IOTA
// Iota is an identifier which is used with constant and which can simplify constant definitions that use auto increment numbers.
// The IOTA keyword represent integer constant starting from zero.
// So essentially it can be used to create effective constant in Go.
// They can also be used to create enum in Go as we will see later in this tutorial.
const (
	a = iota
	b
	c
)

// skipped using blank identifier
const (
	d = iota
	_
	e
	f
)

// IOTA expression
const (
	g = iota
	h = iota + 4
	i = iota * 4
)

// Enum
// IOTA provides an automated way to create a enum in Golang
type Size int

const (
	small = Size(iota)
	medium
	large
	extraLarge
)

func (s Size) toString() {
	switch s {
	case small:
		fmt.Println("Small")
	case medium:
		fmt.Println("Medium")
	case large:
		fmt.Println("Large")
	case extraLarge:
		fmt.Println("Extra Large")
	default:
		fmt.Println("Invalid size entry")
	}
}
