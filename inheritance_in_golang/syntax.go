package main

import "fmt"

// Program 1
func programOne() {
	fmt.Println("Program 1")
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
	// fmt.Println("The color is " + child.color)
}

// Type Hierarchy
func typeHierarchy() {
	fmt.Println("\nType Hierarchy")
	shark := &shark{}
	checkAquatic(shark)
	checkAnimal(shark)
	lion := &lion{}
	checkNonAquatic(lion)
	checkAnimal(lion)
}
