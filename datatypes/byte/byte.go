package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
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