package main

import "fmt"

func fillPartOfElement() {
	sample := [4]int{5, 8}
	fmt.Printf("Sample: Len: %d, %v\n", len(sample), sample)
}

// Accessing array
func accessingArray() {
	fmt.Println("\nAccessing Array")
	sample := [2]string{"aa", "bb"}
	fmt.Println(sample[0])
	fmt.Println(sample[1])

	sample[0] = "xx"
	fmt.Println(sample)
	//sample[3] = "yy" //will raise error
}

func arrayValue() {
	fmt.Println("\nArray are value in go")
	//A copy of the array will be created when
	//- An array variable is assigned to another array variable.
	//- An array variable is passed as an argument to a function.
	sample1 := [2]string{"a", "b"}
	fmt.Printf("Sample 1 Before: %v\n", sample1)
	sample2 := sample1
	sample2[0] = "c"
	fmt.Printf("Sample 1 After assignment: %v\n", sample1)
	fmt.Printf("Sample 2: %v\n", sample2)
	test(sample1)
	fmt.Printf("Sample1 After Test Function Call: %v\n", sample1)
}

func test(sample [2]string) {
	sample[0] = "d"
	fmt.Printf("Sample in Test function: %v\n", sample)
}

// Iterating an array
func iteratingAnArray() {
	fmt.Println("\nItterating an array")
	letters := [3]string{"a", "b", "c"}
	//Using for loop
	fmt.Println("\nUsing for loop")
	len := len(letters)
	for i := 0; i < len; i++ {
		fmt.Println(letters[i])
	}
	//Using for Range operator
	fmt.Println("\nUsing for-range loop")
	for i, letter := range letters {
		fmt.Printf("%d %s\n", i, letter)
	}
}

// Multidimensional array
func multidimensionalArray() {
	fmt.Println("Array Multidimensional")
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("First Run")
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	//Accessing multidimensional array
	fmt.Println("\nAccessing multidimensional array")
	sample[0][0] = 6
	sample[1][2] = 1
	fmt.Println("\nSecond Run")
	for _, row := range sample {
		for _, value := range row {
			fmt.Println(value)
		}
	}
}
