package passchildtypetoparent

import "fmt"

type iBase interface {
	say()
}

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

func check(b iBase) {
	b.say()
}
