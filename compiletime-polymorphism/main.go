package main

import "fmt"

func main() {
	m := &maths{}

	fmt.Printf("Result : %d\n", m.add(2, 3))
	fmt.Printf("Result : %d\n", m.add(2, 3, 4))
}
