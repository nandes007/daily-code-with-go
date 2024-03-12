package callchildfunction

import "fmt"

type iBase interface {
	say()
}

type base struct {
	color string
	clear func()
}

func (b *base) say() {
	b.clear()
}

type child struct {
	base  //embedding
	style string
}

func (c *child) clear() {
	fmt.Println("Clear from child's function")
}

func check(b iBase) {
	b.say()
}
