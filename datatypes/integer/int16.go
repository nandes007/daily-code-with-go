package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a int16
	var a int16 = 2

	//Size of int16 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
