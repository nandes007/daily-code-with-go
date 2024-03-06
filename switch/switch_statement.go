package main

import "fmt"

func main() {
	//Example switch statement
	// switch statement; expression {}
	// statement and expression is optional.

	//Switch with statement and expression
	fmt.Println("Switch with statement and expression")
	switch ch := "b"; ch {
	case "a":
		fmt.Println("a")
	case "b", "c":
		fmt.Println("b or c")
	default:
		fmt.Println("No matching character")
	}

	//Switch with absent statement and expression
	fmt.Println("\nSwitch without statement and expression")
	age := 45
	switch {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	//Only switch with statement
	fmt.Println("\nOnly switch statement")
	switch age := 29; {
	case age < 18:
		fmt.Println("Kid")
	case age >= 18 && age < 40:
		fmt.Println("Young")
	default:
		fmt.Println("Old")
	}

	//Only switch expression
	fmt.Println("\nOnly switch expression")
	char := "a"
	switch char {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("No matching character")
	}

	//Fallthrough keyword
	//In below even though the second case matched it went through the thrid case because of fallthrough keyword
	fmt.Println("\nFallthrough Keyword")
	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	}

	//Break statement
	//break statement will terminate the execution of the switch.
	fmt.Println("\nBreak statement")
	switch char := "b"; char {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
		break
		fmt.Println("after b") // will not execute.
	default:
		fmt.Println("No matching character")
	}

	//Type swtich
	//switch statement can also be used to know the type of an interface at run time.
	fmt.Println("\nType Switch")
	printType("test_string")
	printType(2)
}

func printType(t interface{}) {
	switch v := t.(type) {
	case string:
		fmt.Println("Type: string")
	case int:
		fmt.Println("Type: int")
	default:
		fmt.Printf("Unknown Type %T", v)
	}
}
