package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int64
	var a int = 2

	//Size of int64 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
