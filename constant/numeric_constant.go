package main

import "fmt"

func main() {
	//Type int constant
	const aa int = 123
	var uu = aa
	fmt.Println("Typed named integer constant")
	fmt.Printf("uu: Type: %T Value: %v\n\n", uu, uu)

	//Below line will raise a compilator error
	//var v int32 = aa

	//Untyped named int constant
	const bb = 123
	var ww = bb
	var xx int32 = bb
	var yy float64 = bb
	var zz complex128 = bb
	fmt.Println("Untyped named integet constant")
	fmt.Printf("ww: Type: %T Value: %v\n", ww, ww)
	fmt.Printf("xx: Type: %T Value: %v\n", xx, xx)
	fmt.Printf("yy: Type: %T Value: %v\n", yy, yy)
	fmt.Printf("zz: Type: %T Value: %v\n", zz, zz)

	//Untyped unnamed int constant
	var ll = 123
	var mm int32 = 123
	var nn float64 = 123
	var oo complex128 = 123
	fmt.Println("Untyped unnamed integet constant")
	fmt.Printf("ll: Type: %T Value: %v\n", ll, ll)
	fmt.Printf("mm: Type: %T Value: %v\n", mm, mm)
	fmt.Printf("nn: Type: %T Value: %v\n", nn, nn)
	fmt.Printf("oo: Type: %T Value: %v\n", oo, oo)

	//Numeric expression
	var p = 5.2 / 3
	fmt.Printf("p: Type: %T Value: %v\n", p, p)
}
