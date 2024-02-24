package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a uint64
	var a uint64 = 2

	//Size of uint64 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
