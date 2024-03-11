package main

import (
	"errors"
	"fmt"
	"os"
)

//Error
//Any type which defines the Error method is set to be implementing the error interface.

//Interface of error
//type error interface { Error() string }
//Any type whitch defines the Error method is set to implementing the error interface.

// Example
func example1() {
	file, err := os.Open("non-existing.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "opened successfully")
	}
}

// Different ways of creating an error
func example2() {
	fmt.Println("\nDifferent way of creating error")
	sampleErr := errors.New("error occured")
	fmt.Println(sampleErr)
}

// Using fmt.Errorf("error is %s", "some_error_message").
func usingFmtErrorf() {
	fmt.Println("\nUsing fmt.Errorf")
	sampelErr := fmt.Errorf("err is: %s", "database connection issue")
	fmt.Println(sampelErr)
}

// Custom error
// The below example illustrates the use of custom error. In below example:
// - inputError is a struct that has the Error() method hence it is of type error interface
// - You can also add additional information to the custom erro by extending its fields or by adding new methods.
// - inputError has an additional field named missingFields and function getMissingFields function.
// - We can use type assertion to convert from error to inputError
// See that function in helper file.
func customError() {
	fmt.Println("\nCustom Error")
	err := validate("", "")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err)
			fmt.Printf("Missing Field is %s\n", err.getMissingField())
		}
	}
}

// Ignoring errors
// Underscore('_') operator can be used to ignore the error returned from a function call.
// Before we see a program it's important to note that error should never be ignored.
// It is not a recomended way. Let's see a program
func ignoringError() {
	fmt.Println("\nIgnoring Error")
	file, _ := os.Open("non-existing.txt")
	fmt.Println(file)
}
