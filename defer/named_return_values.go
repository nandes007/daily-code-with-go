package main

import "fmt"

func main() {
	s := test2()
	fmt.Println(s)
}

func test2() (size int) {
	defer func() { size = 20 }()
	size = 30
	return
}
