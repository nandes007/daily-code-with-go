package main

import (
	"fmt"
	"runtime/debug"
)

func print(a []string, index int) {
	fmt.Println(a[index])
}

func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func f1() {
	defer fmt.Println("Defer in f1")
	f2()
	fmt.Println("After panic in f1")
}

func f2() {
	defer fmt.Println("Defer in f2")
	panic("Panic Demo")
	fmt.Println("After panic in f2")
}

// Recover
func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func checkAndPrintWithRecoverSubsequently(a []string, index int) {
	defer handleOutOfBounds()
	checkAndPrint(a, index)
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("recovering from panic:", r)
	}
}

// Recover not within defer function
func checkAndPrintWithoutDefer(a []string, index int) {
	handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

// Panic/Recover and Goroutine
func checkAndPrintGoroutine(a []string, index int) {
	defer handleOutOfBounds()
	go checkAndPrint(a, index)
}

// Stack Trace
func checkAndPrintWithRecoverStackTrace(a []string, index int) {
	defer handleOutOfBoundsStackTrace()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBoundsStackTrace() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		fmt.Println("Stack Trace:")
		debug.PrintStack()
	}
}

// Return value of the function when panic is recovered
func checkAndGet(a []int, index int) (int, error) {
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	return a[index], nil
}

// named return value
func checkAndGetNamedReturnValue(a []int, index int) (value int, err error) {
	value = 10
	defer handleOutOfBounds()
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	value = a[index]
	return value, nil
}
