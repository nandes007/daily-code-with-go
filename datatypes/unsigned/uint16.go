package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Declare a uint16
	var a uint16 = 2

	//Size of uint16 in bytes
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}
