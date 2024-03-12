package callchildfunction

import "fmt"

func CallChildFunction() {
	fmt.Println("\nCall child function")
	base := base{color: "Red", clear: func() {
		fmt.Println("Clear from child's function")
	}}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
}
