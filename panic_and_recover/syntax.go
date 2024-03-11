package main

import (
	"fmt"
	"time"
)

//Panic
//Panic is meant to exit from a program in abnolmal conditions. Panic can occur in a program in two ways
//- Runtime error in the program
//- By calling the panic function explicitly.
//This can be called by the programmer when the program cannot continue and it has to exit
//signature: func panic(v interface{})

// Runtime error in the program can happen in below cases
// Out of bounds array access
// Calling a function on a nil pointer
// Sending on a closed channel
// Incorrect type assertion
func outOfBoundsArrayAccess() {
	fmt.Println("Out of bounds array access")
	a := []string{"a", "b"}
	print(a, 2)
}

// Calling the panic function explicitly
func callingPanicExplicitly() {
	a := []string{"a", "b"}
	checkAndPrint(a, 2)
}

// Panic with defer
func panicWithDefer() {
	defer fmt.Println("Defer in main")
	panic("Panic with Defer")
	fmt.Println("After panic in f2")
}

// Recover in golang
// Go provides a built-in function recover for recovering from a panic.
// signature: func recover() interface{}
// Notes::
// We already learn above that defer function is the only function that is called after the panic.
// So it makes sense to put the recover function in the defer function only.
// If the recover function is not within the defer function then it will not stop panic.
func recoverPanic() {
	fmt.Println("\nRecover panic")
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	fmt.Println("Exiting normally")
}

func recoverPanicSubsequently() {
	fmt.Println("\nRecover panic subsequently")
	a := []string{"a", "b"}
	checkAndPrintWithRecoverSubsequently(a, 2)
	fmt.Println("Exiting normally")
}

// Example if recover function is not within defer function
func recoverNotInDefer() {
	a := []string{"a", "b"}
	checkAndPrintWithoutDefer(a, 2)
	fmt.Println("Exiting normal")
}

// Panic/Recover and Goroutine
// An important point to note about the recover function is that it can only recover the
// panic happening in the same goroutine. If the panic is happening in different goroutine and
// recover is in a different goroutine the it won't stop panic.
func panicRecoverGoroutine() {
	fmt.Println("\nPanic Recover and Goroutine")
	a := []string{"a", "b"}
	checkAndPrintGoroutine(a, 2)
	time.Sleep(time.Second)
	fmt.Println("Exiting normally")
}

// Printing stack trace
// Debug package of golang also provide StackTrace function
// that can be used print the stack trace of the panic in the recover function
func debugStackTrace() {
	fmt.Println("\nDebug stack trace")
	a := []string{"a", "b"}
	checkAndPrintWithRecoverStackTrace(a, 2)
	fmt.Println("Exiting normally")
}

// Return value of the function when panic is recovered
// When the panic is recovered then the return value of panicking function will be the
// default value of the return types of the panicking function
func returnPanicValue() {
	fmt.Println("\nReturn value of the function when panic is recovered")
	a := []int{5, 6}
	val, err := checkAndGet(a, 2)
	fmt.Printf("Val: %d\n", val)
	fmt.Println("Error: ", err)
}

// If don't want to return default zero value of types then named return value can be used.
// See function in helper.go
func namedReturnValue() {
	fmt.Println("\nNamed return value")
	a := []int{5, 6}
	val, err := checkAndGetNamedReturnValue(a, 2)
	fmt.Printf("Val: %d\n", val)
	fmt.Println("Error: ", err)
}
