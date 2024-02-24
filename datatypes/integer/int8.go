package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int 8
	var a int8 = 2

	//Size of int8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
