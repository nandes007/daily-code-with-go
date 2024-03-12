package multipleinheritance

import "fmt"

// Multiple inheritance - child struct should be able to access multiple properties
// and methods from two base struct and also sub-typing should be possible.
func MultipleInheritance() {
	fmt.Println("\nMultiple inheritance")
	base1 := base1{color: "Red"}
	base2 := base2{}
	child := &child{
		base1: base1,
		base2: base2,
		style: "somestyle",
	}
	child.say()
	child.walk()
	check1(child)
	check2(child)
}
