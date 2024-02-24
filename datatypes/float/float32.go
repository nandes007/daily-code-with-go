package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a float32
	var a float32 = 2

	//Size of float32 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
