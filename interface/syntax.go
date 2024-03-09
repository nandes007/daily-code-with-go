package main

import "fmt"

// Interface a collection of method signatures.
// Signature
// type name_of_interface interface{}
func craeteVariableOfAnimal() {
	var a animal
	fmt.Println(a)
}

func implementInterfaceOfAnimal() {
	var a animal
	a = lion{age: 10}
	a.breathe()
	a.walk()

	a = dog{age: 2}
	a.breathe()
	a.walk()
}

// Interface types as argument to a function
func callBreathe(a animal) {
	a.breathe()
}

func callWalk(a animal) {
	a.walk()
}

func interfaceTypesAsArgument() {
	fmt.Println("\nInterface Types As Argument to a function")
	l := lion{age: 10}
	callBreathe(l)
	callWalk(l)

	d := dog{age: 5}
	callBreathe(d)
	callWalk(d)
}

func runtimePolymorphism() {
	fmt.Println("\nRuntime Polymorphism")
	indianTax := &indianTax{
		income:        1000,
		taxPercentage: 30,
	}
	singaporeTax := &singaporeTax{
		income:        2000,
		taxPercentage: 10,
	}

	taxSystems := []taxSystem{indianTax, singaporeTax}
	totalTax := calculateTotalTax(taxSystems)

	fmt.Printf("Total Tax is %d\n", totalTax)
}

// Pointer receiver while implementing an interface
//Notes:
//-If a type implements all methods of an interface using value receiver,
//then both variable of that type as well pointer to the variable of that type can be used while assigning to that interface
//or while passing to a function which accepts an argument as that interface.

//-If a type implements all methods of an interface using pointer receiver,
//then the only pointer to the variable of that type can be used while assigning to that interface
//or while passing to a function that accepts an argument as that interface.

func pointerReceiverWhileImplementingInterface() {
	fmt.Println("\nPointer Receiver While Implementing Interface")
	var a animal

	a = lion{age: 10}
	a.breathe()
	a.walk()

	a = &lion{age: 5}
	a.breathe()
	a.walk()
}

func implementingPointerReceiver() {
	fmt.Println("\nImplementing pointer receiving")
	var a animalPointer

	// a = lion{age: 10} //This will raise error
	// a.breathePointer()
	// a.walkPointer()

	a = &lion{age: 5}
	a.breathePointer()
	a.walkPointer()
}

func customTypeImplementingInterface() {
	fmt.Println("\nCustom type implementing interface")
	var a animal

	a = cat("smoky")
	a.breathe()
	a.walk()
}

// Type implementing multiple interface
func implementingMultipleInterface() {
	fmt.Println("\nImplementing multiple interface")
	var a animal
	l := lion{}
	a = l
	a.breathe()
	a.walk()

	var m mamal
	m = l
	m.feed()
}

// Zero value of interface
func zeroValueOfInterface() {
	fmt.Println("\nZero value of interface")
	var a animal
	fmt.Println(a)
}

// Inner working of interface
func innerWorkingInterface() {
	fmt.Println("\nInner working interface")
	var a animal
	a = lion{age: 10}
	fmt.Printf("Underlying Type: %T\n", a)
	fmt.Printf("Underlying Value: %v\n", a)
}

// Embedding interface
func embedInterface() {
	fmt.Println("\nEmedding interface")
	var h human
	h = employee{name: "Johm"}
	h.breathe()
	h.walk()
	h.speak()
}

// Embedding interface in a struct
func embeddingInterfaceInStruct() {
	fmt.Println("\nEmbedding Interface In Struct")
	d := dog{age: 5}
	p1 := pet1{name: "Milo", a: d}

	fmt.Println(p1.name)
	p1.a.breathe()
	p1.a.walk()

	p2 := pet2{name: "Oscar", animal: d}
	fmt.Println(p1.name)
	p2.breathe()
	p2.walk()
}

// Type assertion
// Type assertion provides a way to access the underlying variable inside the interface value of the interface
// by asserting the correct type of underlying value. Below is the syntax for that where i is an interface.
// Signature
// val := i.({type})
func typeAssertion() {
	fmt.Println("\nType Assertion")
	a := lion{age: 10}
	print(a)
	printTwo(a)
}

// Type switch
// Type switch enables us to do above type assertion in series. See below code example for the same
func typeSwitch() {
	fmt.Println("\nType Switch")
	a := lion{age: 10}
	printSwitch(a)
}

// Empty Interface
// An empty interface has no methods, hence by default all concrete types implement the empty interface.
// If you write a function that accepts an empty interface then you can pass any type to that function.
// See working code below:
func test(a interface{}) {
	fmt.Printf("(%v, %T)\n", a, a)
}

func emptyInterface() {
	fmt.Println("\nEmpty Interface")
	test("thisisstring")
	test("10")
	test(true)
}
