package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main() {
	r := 'a'

	//Print size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Code Point
	fmt.Printf("Unicode CodePoint: %U\n", r)

	//Print Character
	fmt.Printf("Character: %c\n", r)
	s := "0b£"

	//This will print the unicode Points
	fmt.Printf("%U\n", []rune(s))

	//This will the decimal value of unicode code point
	fmt.Println([]rune(s))
}