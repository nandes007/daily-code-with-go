package simpleinheritance

import "fmt"

//all declaration in base.go file
//This section is base use case if inheritance is child type should be able access common data an method of parent type.
//This done in GO via embedding.
//Note::One of limitation of this implementation is that we cannot pass the child type to a function that expects the base type as Go
//does allow type inheritance. See pass_child_type_to_parent for fix this.

func CallParentMethoAndAttribute() {
	fmt.Println("Child type call method in parent via embedding")
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
}
