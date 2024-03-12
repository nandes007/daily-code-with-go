package passchildtypetoparent

import "fmt"

// all declaration in base.go file in this package.
// This is where GO interfaces come to fix sub-typing
func PassChildTypeToParent() {
	fmt.Println("\nPass child type to parent for fix sub-typing")
	base := base{color: "Red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}
	child.say()
	fmt.Println("The color is " + child.color)
	check(child)
}

//Since base struct implements the "say" method in turn,
//child struct embeds base.
//So the child method indirect implements "say" method and becomes a type if "iBase" and that is why we can pass
//the child to check function now.

//Another limitation:
//Let’s say child and base both have one more function “clear”.
//Now “say” method makes a call to “clear” method.
//Then when “say” method is called using child struct, in turn,
//“say” method will call “clear” method of base and not “clear” method of the child.
//See callchildfunctionpackage for fix above case
