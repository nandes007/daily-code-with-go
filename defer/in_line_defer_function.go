package main

import "fmt"

func main() {
	defer func() { fmt.Println("In inline defer") }()
	fmt.Println("Executed")
}
