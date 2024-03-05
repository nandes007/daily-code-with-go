package main

import "fmt"

const name = "test"

func main() {
	const a = 8
	fmt.Println(a)

}

func testGlobal() {
	fmt.Println(name)
}
