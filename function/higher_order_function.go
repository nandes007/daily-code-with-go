package main

import "fmt"

func main() {
	areaF := getAreaFunc()
	print(3, 4, areaF)
}

func print(x, y int, area func(int, int) int) {
	fmt.Printf("Area is: %d\n", area(x, y))
}

func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}