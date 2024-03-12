package main

import (
	callchildfunction "golangbyexample.com/inheritance_in_golang/call_child_function"
	multipleinheritance "golangbyexample.com/inheritance_in_golang/multiple_inheritance"
	passchildtypetoparent "golangbyexample.com/inheritance_in_golang/pass_child_type_to_parent"
	simpleinheritance "golangbyexample.com/inheritance_in_golang/simple_inheritance"
	typehierarchy "golangbyexample.com/inheritance_in_golang/type_hierarchy"
)

func main() {
	simpleinheritance.CallParentMethoAndAttribute()

	passchildtypetoparent.PassChildTypeToParent()

	callchildfunction.CallChildFunction()

	multipleinheritance.MultipleInheritance()

	typehierarchy.TypeHierarchy()
}
