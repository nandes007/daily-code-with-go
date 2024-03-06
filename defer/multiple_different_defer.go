package main

import "fmt"

func main() {
	defer fmt.Println("Defer in main")
	fmt.Println("Start main")
	f1()
	fmt.Println("Finish main")

	multipleDeferFunction()
}

func f1() {
	defer fmt.Println("Defer in f1")
	fmt.Println("Start f1")
	f2()
	fmt.Println("Finish f1")
}

func f2() {
	defer fmt.Println("Defer in f2")
	fmt.Println("Start f2")
	fmt.Println("Finish f2")
}

// Multiple defer function in same function
func multipleDeferFunction() {
	fmt.Println("\nMultiple defer function in same function")
	i := 0
	i = 1
	defer fmt.Println(i)
	i = 2
	defer fmt.Println(i)
	i = 3
	defer fmt.Println(i)
}
