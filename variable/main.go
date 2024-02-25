package main

import "fmt"

var kkk int = 10 //This is global variable

func main() {
	//Single variable declartion without initial value
	var aaa int
	fmt.Println(aaa)

	//Single variable declarion with initial value
	var bbb int = 8
	fmt.Println(bbb)

	//Multiple declartion variable without initial value
	var ccc, ddd int
	fmt.Println(ccc)
	fmt.Println(ddd)

	var eee, fff int = 8, 9
	fmt.Println(eee)
	fmt.Println(fff)

	//Declare variable of different types
	var (
		ggg int
		hhh int    = 8
		iii string = "a"
	)

	fmt.Println(ggg)
	fmt.Println(hhh)
	fmt.Println(iii)

	//Variable declaration with no type of type interface
	var t = 123                  //type inferred will be int
	var u = "circle"             //type inferred will be string
	var v = 5.6                  //type inffered will be float64
	var w = true                 //type inffered will be bool
	var x = 'a'                  //type inffered will be rune
	var y = 3 + 5i               //type inffered will be complex128
	var z = Sample{name: "test"} //type inffered will be main.Sample

	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(kkk)
}

type Sample struct {
	name string
}
