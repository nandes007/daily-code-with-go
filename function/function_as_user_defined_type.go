package main

import "fmt"

func main() {
	areaF := getAreaFunc()
	print(3, 4, areaF)
}

type area func(int, int) int

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x, y))
}

func getAreaFunc() area {
	return func(x, y int) int {
		return x * y
	}
}