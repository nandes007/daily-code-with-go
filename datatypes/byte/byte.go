package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	/*
		*Byte
		in Go is an alias for uint8 meaning it is an integer value. This integer value is of 8 bits and it represents one byte i.e number between 0-255). A single byte therefore can represent ASCII characters. Golang does not have any data type of ‘char’. Therefore
		var r byte = 'a'
	*/
	var r byte = 'a'

	//Print size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print character
	fmt.Printf("Character: %c\n", r)
	s := "abc"

	//This will the decimal value of byte
	fmt.Println([]byte(s))
}