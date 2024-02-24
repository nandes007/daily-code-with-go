package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a uint
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
