package main

import "fmt"

func main() {
	defer test()
	fmt.Println("Executed in main")
}

func test() {
	fmt.Println("In Defer")
}
