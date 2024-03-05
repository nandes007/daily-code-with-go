package main

import (
	"fmt"
	"math"
)

func main() {
	//Type constant
	const a int32 = 8
	fmt.Println(a)

	fmt.Println("")

	//Unanamed untype constant
	fmt.Printf("Type: %T Value: %v\n", 123, 123)
	fmt.Printf("Type: %T Value: %v\n", "circle", "circle")
	fmt.Printf("Type: %T Value: %v\n", 5.6, 5.6)
	fmt.Printf("Type: %T Value: %v\n", true, true)
	fmt.Printf("Type: %T Value: %v\n", 'a', 'a')
	fmt.Printf("Type: %T Value: %v\n", 3+5i, 3+5i)

	//Named untype constant
	const b = 123     //Default hidden type is int
	const c = "cicle" //Default hidden type is string
	const d = 5.6     //Default hidden type is float64
	const e = true    //Default hidden type is bool
	const f = 'a'     //Default hidden type is rune
	const g = 3 + 5i  //Default hidden type is complex128

	fmt.Println("")
	fmt.Printf("Type: %T Value: %v\n", b, b)
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Printf("Type: %T Value: %v\n", d, d)
	fmt.Printf("Type: %T Value: %v\n", e, e)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", g, g)

	fmt.Println("")
	var f1 float32
	var f2 float64
	f1 = math.Pi
	f2 = math.Pi

	fmt.Printf("Type: %T Value: %v\n", math.Pi, math.Pi)
	fmt.Printf("Type: %T Value: %v\n", f1, f1)
	fmt.Printf("Type: %T Value: %v\n", f2, f2)

}
