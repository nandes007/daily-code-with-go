package main

import "fmt"

func main() {
	//Default value will be false it not initialized
	var a bool
	fmt.Printf("a's value is %t\n", a)

	//And operation on one true and other false
	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Output of AND operation on one true and other false %t\n", andOperation)

	//OR operation on one true and other false
	orOperator := 1 < 2 || 1 > 3
	fmt.Printf("Output of OR operation on one true and other false %t\n", orOperator)

	//Negation Operation on a false value
	negotionOperation := !(1 > 2)
	fmt.Printf("Output of NEGOTION operation on false value %t\n", negotionOperation)
}