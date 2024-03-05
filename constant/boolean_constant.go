package main

import "fmt"

func main() {
	type myBool bool

	//Type boolean constant
	const aa bool = true
	var uu = aa
	fmt.Println("Type named boolean constant")
	fmt.Printf("uu: Type: %T Value: %v\n\n", uu, uu)

	//Below line will raise a compilation error
	//var vv myBool = aa

	//Untyped named boolean constant
	const bb = true

	var ww myBool = bb
	var xx = bb
	fmt.Println("Untyped named boolean constant")
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n\n", xx, xx)

	//Untyped unnamed boolean constant
	var yy myBool = true
	var zz = true
	fmt.Println("Untyped unnamed boolean constant")
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)
}
