package main

import "fmt"

//Goto statement allows unconditional jump to a labeled statement with in the same function.
//Below is the format of the goto statement
//goto label
// ...
// ...
//label: statement

func gotoStatement() {
	fmt.Println("goto statement")
	fmt.Println("a")
	goto FINISH
	fmt.Println("b")
FINISH:
	fmt.Println("c")
}

func labelBeforeGotoStatement() {
	fmt.Println("\nLabel Before goto statement")
	i := 0
	fmt.Println("here")
START:
	for i < 10 {
		if i%2 == 0 {
			i++
			goto START
		}
		fmt.Println(i)
		i++
	}
}
