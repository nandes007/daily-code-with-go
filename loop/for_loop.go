package main

import "fmt"

func main() {
	//Simple for loop
	fmt.Println("simple loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for loop with only declaration
	fmt.Println("loop with only declaration")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//break statement
	fmt.Println("break statement")
	j := 0
	for {
		fmt.Println(j)
		j++
		if j >= 5 {
			break
		}
	}

	//continue statement
	fmt.Println("continue statement")
	for i = 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//nested for loop
	fmt.Println("Nested for loop")
	for i = 0; i < 3; i++ {
		fmt.Printf("Outer loop iteration %d\n", i)
		for j = 0; j < 2; j++ {
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}

	//function call and assignment in init part
	i = 1
	for test(); i < 3; i++ {
		fmt.Println(i)
	}

	for i = 2; i < 3; i++ {
		fmt.Println(i)
	}

	//Implementing while loop using for loop
	fmt.Println("implement while loop using for loop")
	i = 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
}

func test() {
	fmt.Println("In test function")
}
