package main

//In compile-time polymorphism, the call is resolved during compile time by the compiler.
//Some if the forms for compile-time polimorphism are:
//- Method/Function Overloading: more than one method/function exists with the same but with different signature of possibly different return types
//- Operator overloading: the same operator is used for operating on different data types

// Go doesn't support method overloading. For example below program:
type maths struct{}

// func (m *maths) add(a, b int) int {
// 	return a + b
// }

// func (m *maths) add(a, b, c int) int {
// 	return a + b + c
// }

//The program above will raise error:
//(*maths).add redeclared in this block previous declaration at ...
//Go also doesn't support operator overloading

// Variadic function in go can use as alternative to do method overloading in GO.
func (m *maths) add(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}
