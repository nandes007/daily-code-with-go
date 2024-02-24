package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a unit8
	var a uint8 = 2

	//Size of unit8 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
