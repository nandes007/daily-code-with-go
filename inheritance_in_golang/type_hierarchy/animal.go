package typehierarchy

import "fmt"

type iAnimal interface {
	breathe()
}

type animal struct {
}

func (a *animal) breathe() {
	fmt.Println("Animal brethe")
}

type iAquatic interface {
	iAnimal
	swim()
}

type aquatic struct {
	animal
}

func (a *aquatic) swim() {
	fmt.Println("Aquatic swim")
}

type iNonAquatic interface {
	iAnimal
	walk()
}

type nonAquatic struct {
	animal
}

func (a *nonAquatic) walk() {
	fmt.Println("Non-Aquatic walk")
}

type shark struct {
	aquatic
}

type lion struct {
	nonAquatic
}

func checkAquatic(a iAquatic)       {}
func checkNonAquatic(a iNonAquatic) {}
func checkAnimal(a iAnimal)         {}
