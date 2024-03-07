package main

import "fmt"

func main() {
	//Both number of elements and actual elements
	sample1 := [2]int{1, 2}
	fmt.Printf("Sample1: Len: %d, %v\n", len(sample1), sample1)

	//Only actual elements
	sample2 := [...]int{2, 3}
	fmt.Printf("Sample2: Len: %d, %v\n", len(sample2), sample2)

	//Only number of elements
	sample3 := [2]int{}
	fmt.Printf("Sample3: Len: %d, %v\n", len(sample3), sample3)

	//Without both number of elements and actual elements
	sample4 := [...]int{}
	fmt.Printf("Sample4: Len: %d, %v\n", len(sample4), sample4)

	//Spesified part of elements
	fillPartOfElement()

	//Accessing Element of Array
	accessingArray()

	//Array are value in go
	arrayValue()

	//Iterating an array
	iteratingAnArray()

	//Multidimensional array
	multidimensionalArray()
}
