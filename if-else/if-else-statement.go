package main

import "fmt"

func main() {
	a := 6
	if a > 5 {
		fmt.Println("a is grather than 5")
	}

	b := 4
	if b > 3 && b < 6 {
		fmt.Println("b is within range")
	}

	var (
		c int = 1
		d int = 2
	)

	if c > d {
		fmt.Println("c is grather than d")
	} else {
		fmt.Println("d is grather than c")
	}

	//If else ladder
	fmt.Println("\nIf Else Ladder")
	age := 29
	if age < 18 {
		fmt.Println("Kid")
	} else if age >= 18 && age < 40 {
		fmt.Println("Young")
	} else {
		fmt.Println("Old")
	}

	// Nested if else statement
	fmt.Println("\nNested if else statement")
	e, f, g := 1, 2, 3
	if e > f {
		if e > g {
			fmt.Println("Biggest is e")
		} else if f > g {
			fmt.Println("Biggest is f")
		}
	} else if f > g {
		fmt.Println("Biggest is f")
	} else {
		fmt.Println("Biggest is g")
	}

	//If with short statment
	fmt.Println("\nIf with short statement")
	if h := 6; h > 5 {
		fmt.Println("h is grather than 5")
	} else {
		fmt.Println("h is less than 5")
	}
}
