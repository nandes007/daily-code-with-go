package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	//Declare an int
	// var a int

	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	b := 2
	fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))
}
