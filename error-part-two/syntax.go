package main

import (
	"errors"
	"fmt"
	"os"
)

//Wrapping of error
//What does wrapping of error mean?
//It means to create a hierarchy of errors in which a particular instance of error wraps
//another error and that particular instance itself can be wrapped inside another error.
//syntax for wrapping erro: e := fmt.Errorf("... %w ...", ..., err, ...)
//%w directive is used for wrapping the error. The fmt.Errof should be called with only one %w directive.

func wrappedError() {
	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(e2)
	fmt.Println(e3)
}

// Use case of wrapping the errors.
func useCaseWrappingErr() {
	fmt.Println("\nUse case of wrapping the errors")
	num := 3
	err := checkPositiveAndEven(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Givennumber is positive and even")
	}
}

// Unwrap an error
// signature: func Unwrap(err error) error
func unWrappedError() {
	fmt.Println("\nUnwrapped Error")
	e1 := errorOne{}
	e2 := fmt.Errorf("E2: %w", e1)
	e3 := fmt.Errorf("E3: %w", e2)
	fmt.Println(errors.Unwrap(e3))
	fmt.Println(errors.Unwrap(e2))
	fmt.Println(errors.Unwrap(e1))
}

// Check if two error are equal
// Using == (equality operator)
// Using Is function
// signature: func Is(err, target error) bool
func checkError() {
	fmt.Println("\nCheck if two error are equal")
	var err1 errorOne
	err2 := do()
	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}
	if errors.Is(err1, err2) {
		fmt.Println("Is function: Both errors are equal")
	}
}

// Is function preferable to using the equality operator
// because it checks for equality by unwrapping the first error sequentially
// and matches it with the target error at each step of unwrap
func example() {
	fmt.Println("\nExample to use Is function is preferable")
	err1 := errorOne{}
	err2 := do()

	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	} else {
		fmt.Println("Equality Operator: Both errors are not equal")
	}

	if errors.Is(err2, err1) {
		fmt.Println("Is function: Both errors are equal")
	}
}

// Get the underlying error from an error represented by the error interface
// There are two ways of getting the underlying type
// -Using the .({type}) assert
// -Using th As function of errors package, signature: func As(err error, target interface{}) bool
func getUnderlyingError() {
	fmt.Println("\nGet the underlying error from an error represeted by the error interface")
	err := openFile("non-existing.txt")

	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("Using Assert: error e is of type path error. Path: %v\n", e.Path)
	} else {
		fmt.Println("Using Assert: error not of type path error")
	}

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Using As function: Error e is of type path error. Path: %v\n", err)
	}
}

// As function is preferable to using the .({type}) assert
// because it checks for a match by unwrapping the first error sequentially
// and matches it with the target error at each step of unwrap.
func preferAsFunctionTypeAssert() {
	fmt.Println("\nAs function is preferable to using the .({type})")
	var pathError *os.PathError
	err := openFileWrappedError("non-existing.txt")

	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("Using Assert: Error e is of type path error. Error: %v\n", e)
	} else {
		fmt.Printf("Using Assert: Error not of type path error\n")
	}

	if errors.As(err, &pathError) {
		fmt.Printf("Using As function: Error e is of type path error. Error: %v\n", pathError)
	}
}
