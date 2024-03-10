package main

import (
	"fmt"
	"runtime"
)

func main() {
	executeFunctionWithGoroutine()

	createMultipleGoroutines()

	//Get runtime Num cpu
	fmt.Println("\nGet Runtime num CPU")
	fmt.Println(runtime.NumCPU())

	anonymousGoroutines()
}
