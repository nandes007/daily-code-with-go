package main

import (
	"fmt"
	"unsafe"
)

type sample struct {
	a int
	b string
}

func main() {
	s := &sample{
		a: 1,
		b: "test",
	}

	//Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	//Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))
}
