package main

import "fmt"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("\nSkipped with blank identifier")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println("\nIOTA Expression")
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Println("\nEnum")
	fmt.Println(small)
	fmt.Println(medium)
	fmt.Println(large)
	fmt.Println(extraLarge)

	fmt.Println("\nTo String")
	var m Size = 1
	m.toString()
}
