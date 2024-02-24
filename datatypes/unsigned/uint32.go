package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a uint32
	var a uint32 = 2

	//Size of uint32 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
