package main

import "fmt"

func main() {
	result := sum(2,3)
	fmt.Println(result)
}

func sum(a,b int) (result int) {
	return a + b
}