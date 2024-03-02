package main

import "fmt"

var max = func(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	res := max(2, 3)
	fmt.Println(res)
}