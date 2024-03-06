package main

import "fmt"

func main() {
	defer fmt.Println("Defer in main")
	panic("panic with defer")
	fmt.Println("after panic in f2")
}
