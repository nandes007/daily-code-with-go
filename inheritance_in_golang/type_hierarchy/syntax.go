package typehierarchy

import "fmt"

// Type Hierarchy
func TypeHierarchy() {
	fmt.Println("\nType Hierarchy")
	shark := &shark{}
	checkAquatic(shark)
	checkAnimal(shark)
	lion := &lion{}
	checkNonAquatic(lion)
	checkAnimal(lion)
}
