package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int32
	var a int32 = 2

	//Size of int32 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
