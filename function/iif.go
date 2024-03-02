package main

import "fmt"

func main() {
	//IIF (Immediately Invoked Function)
	//That is function which can be definied and executed at the same time.
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)
}