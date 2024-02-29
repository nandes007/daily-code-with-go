package main

import "fmt"

func main() {
	//String in double quotes
	x := "this\nthat"
	fmt.Printf("x is: %s\n", x)

	//String in back quotes
	y := `this\nthat`
	fmt.Printf("y is %s\n", y)
	s := "ab£"

	//This will print the byte sequence.
	//Since character a and b occupies 1 byte each and £ character occupies 2 bytes.
	//The final output will 4 bytes
	fmt.Println([]byte(s))

	//The output will be 4 for same reason as above
	fmt.Println(len(s))

	//range loops over sequences of byte which form each character
	for _, c := range s {
		fmt.Println(c)
	}

	//Concatenation
	fmt.Println("c" + "d")
}