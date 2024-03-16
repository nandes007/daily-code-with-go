package main

import (
	"fmt"
	"unsafe"
)

func example() {
	s := &sample{a: 1, b: "test"}

	//Getting the address of field b in struct s
	p := unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Offsetof(s.b))

	//Typecasting it to a string pointer and printing the value of it
	fmt.Println(*(*string)(p))
}

func convertToUnsafe() {
	fmt.Println("\nConvert to unsafe")
	s := &sample{
		a: 1,
		b: "test",
	}

	//Get the address as a uintptr
	startAddress := uintptr(unsafe.Pointer(s))
	fmt.Printf("Start Address of s: %d\n", startAddress)
}
