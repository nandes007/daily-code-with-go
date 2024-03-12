package multipleinheritance

import "fmt"

type iBase1 interface {
	say()
}

type iBase2 interface {
	walk()
}

type base1 struct {
	color string
}

func (b *base1) say() {
	fmt.Println("Hi from say function")
}

type base2 struct{}

func (b *base1) walk() {
	fmt.Println("Hi from walk function")
}

type child struct {
	base1 //embedding
	base2 //embedding
	style string
}

func (c *child) clear() {
	fmt.Println("Clear from child's function")
}

func check1(b iBase1) {
	b.say()
}

func check2(b iBase2) {
	b.walk()
}
