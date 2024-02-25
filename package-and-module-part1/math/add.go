package math

import "fmt"

func init() {
	fmt.Println("In add init")
}

func Add(a, b int) int {
	return a + b
}
