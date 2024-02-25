package math

import "fmt"

func init() {
	fmt.Println("In subtract init")
}

func Subtract(a, b int) int {
	return a - b
}
