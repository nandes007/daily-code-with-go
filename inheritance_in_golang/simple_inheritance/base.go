package simpleinheritance

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("Hi from say function")
}

type child struct {
	base  //embedding
	style string
}
