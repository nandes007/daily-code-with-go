package main

import "fmt"

func declareSlice() {
	sample := []int{}
	fmt.Println("Declaring array")
	fmt.Println(len(sample))
	fmt.Println(cap(sample))
	fmt.Println(sample)

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)
}

func creatingSliceFromArray() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\nCreating slice from array")

	//Both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length=%d\n", len(num1))
	fmt.Printf("capacity=%d\n", cap(num1))

	//Only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("length=%d\n", len(num2))
	fmt.Printf("capacity=%d\n", cap(num2))

	//Only End
	num3 := numbers[:3]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num3)
	fmt.Printf("length=%d\n", len(num3))
	fmt.Printf("capacity=%d\n", cap(num3))

	//None
	num4 := numbers[:]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num4)
	fmt.Printf("length=%d\n", len(num4))
	fmt.Printf("capacity=%d\n", cap(num4))

	//The newly created slice still refer the original array.
	fmt.Println("\nRefer to original array")
	numbers[3] = 8
	fmt.Printf("num2=%v\n", num2)
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("num4=%v\n", num4)
}

func sliceFromSlice() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("\nSlice From Slice")

	//Both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length=%d\n", len(num1))
	fmt.Printf("capacity=%d\n", cap(num1))

	//Only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("length=%d\n", len(num2))
	fmt.Printf("capacity=%d\n", cap(num2))

	//Only end
	num3 := numbers[:3]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num3)
	fmt.Printf("length=%d\n", len(num3))
	fmt.Printf("capacity=%d\n", cap(num3))

	//None
	num4 := numbers[:]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num4)
	fmt.Printf("length=%d\n", len(num4))
	fmt.Printf("capacity=%d\n", cap(num4))

	//Change element
	numbers[3] = 8
	fmt.Println("\nChange Element")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("num4=%v\n", num4)
}

func createSliceUsingMakeFunction() {
	//Create slice using make function:
	//signature:
	//func make([]{type}, length, capacity int) []{type}
	//Capacity is optional
	numbers := make([]int, 3, 5)
	fmt.Println("\nCreate slice with make function")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//With capacity ommited
	numbers = make([]int, 3)
	fmt.Println("\nCapacity Ommited")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

func createSliceUsingNewFunction() {
	//Create slice using new function
	numbers := new([]int)
	fmt.Println("\nCreate slice using new function")
	fmt.Printf("numbers=%v\n", *numbers)
	fmt.Printf("length=%d\n", len(*numbers))
	fmt.Printf("capacity=%d\n", cap(*numbers))
}

func lengthVsCapacity() {
	//Lenth vs capacity
	numbers := make([]int, 3, 5)
	fmt.Println("\nLength vs Capacity")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//This line will cause a runtime error index out of range [4] with length 3
	//numbers[4] = 5

	//Increasing the length from 3 to 5
	numbers = numbers[0:5]
	fmt.Println("\nIncreasing length from 3 to 5")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Decreasing the length from 3 to 2
	numbers = numbers[0:2]
	fmt.Println("\nDecreasing length from 3 to 2")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// Accessing and modifying slice elements
func AccessingAndModifyingSlice() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[:]

	//Modifying the slice
	fmt.Println("\nAccessing and Modifying Slice")
	fmt.Println("Modifying Slice")
	fmt.Printf("Array=%v\n", array)
	fmt.Printf("Slice=%v\n", slice)

	//Modifying the array. Would reflect back in slice too
	array[1] = 2
	fmt.Println("\nModifying Underlying Array")
	fmt.Printf("Array=%v\n", array)
	fmt.Printf("Slice=%v\n", slice)
}

func iteratingSlice() {
	//Slice can be iterating using:
	//- for loop
	//- for-range loop
	letters := []string{"a", "b", "c"}
	fmt.Println("\nIterating Slice")
	//Using for loop
	fmt.Println("\nUsing for loop")
	len := len(letters)
	for i := 0; i < len; i++ {
		fmt.Println(letters[i])
	}

	//Using for-range operator
	fmt.Println("\nUsing for-range loop")
	for i, letter := range letters {
		fmt.Printf("%d %s\n", i, letter)
	}
}

// Appending to a slice
func appendingToSlice() {
	//Signature:
	//func append(slice []Type, elems ...Type)[]Type
	numbers := make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("\nAppending Slice")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 4
	numbers = append(numbers, 4)
	fmt.Println("\nAppend number 4")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 5
	numbers = append(numbers, 5)
	fmt.Println("\nAppend number 5")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// Appending slice more than capacity
func appendMoreThanCapacity() {
	numbers := make([]int, 3, 3)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("\nAppend more than capacity")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 4
	numbers = append(numbers, 4)
	fmt.Println("\nAppend number 4")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// Append one slice to another
func appendOneSliceToAnother() {
	numbers1 := []int{1, 2}
	numbers2 := []int{3, 4}
	numbers := append(numbers1, numbers2...)
	fmt.Println("\nAppend one slice to another")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// Copy a slice
// go builtin package provides copy function that can be used to copy a slice.
// signature:
// func copy(dst, src []Type) int
func copySlice() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 5)

	numberOfElementsCopied := copy(dst, src)
	fmt.Println("\nCopied slice")
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//After changing numbers2
	dst[0] = 10
	fmt.Println("\nAfter changing dst")
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)
}

// Nil slice
// The default zero value of slice is nil.
func nilSlice() {
	var numbers []int
	fmt.Println("\nNil Slice")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
	numbers = append(numbers, 1)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// Multidimensional slice
func multidimensionalSlice() {
	twoDSlice1 := make([][]int, 3)
	for i := range twoDSlice1 {
		twoDSlice1[i] = make([]int, 3)
	}
	fmt.Println("\nMultidimensional Slice")
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice1))
	fmt.Printf("Number of column in arsliceray: %d\n", len(twoDSlice1[0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice1)*len(twoDSlice1[0]))
	for _, row := range twoDSlice1 {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	twoDSlice2 := make([][]int, 2)
	twoDSlice2[0] = []int{1, 2, 3}
	twoDSlice2[1] = []int{4, 5, 6}
	fmt.Println()
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice2))
	fmt.Printf("Number of columns is arsliceray: %d\n", len(twoDSlice2[0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(twoDSlice2)*len(twoDSlice2[0]))
	fmt.Println("Second Slice")
	for _, row := range twoDSlice2 {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}

func differentInnerLength() {
	twoDSlice := make([][]int, 2)
	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{4, 5}

	fmt.Println("\nDifferent inner length")
	fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice))
	fmt.Printf("Len of first row: %d\n", len(twoDSlice[0]))
	fmt.Printf("Len of second row: %d\n", len(twoDSlice[1]))
	fmt.Println("Traversing slice")
	for _, row := range twoDSlice {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}

// Three dimensional array
func threeDimensionalArray() {
	sample := make([][][]int, 2)
	for i := range sample {
		sample[i] = make([][]int, 2)
		for j := range sample[i] {
			sample[i][j] = make([]int, 3)
		}
	}

	fmt.Println("\nThree dimensional array")
	fmt.Printf("Length of first dimension: %d\n", len(sample))
	fmt.Printf("Length of second dimension: %d\n", len(sample[0]))
	fmt.Printf("Length of thrid dimension: %d\n", len(sample[0][0]))
	fmt.Printf("Overall Dimension of th slice: %d*%d*%d\n", len(sample), len(sample[0]), len(sample[0][0]))
	fmt.Printf("Total number of elements in slice: %d\n", len(sample)*len(sample[0])*len(sample[0][0]))
	for _, first := range sample {
		for _, second := range first {
			for _, value := range second {
				fmt.Println(value)
			}
		}
	}
}
