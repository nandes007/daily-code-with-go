package main

import (
	"fmt"
	"time"
)

//Goroutines
//Goroutines can be thought of as a lightweight thread that has a separate independent execution and
//which can execute concurrently with other goroutines.
//It is a function or method that is executing concurrently with other goroutines.
//It is entirely managed by the GO runtime. Golang is a concurrent language. Each goroutine is an independent execution.
//It is goroutine that helps achieve concurrency in golang

func executeFunctionWithGoroutine() {
	fmt.Println("Execute function with goroutine")
	go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func createMultipleGoroutines() {
	fmt.Println("\nCreating Multiple Goroutines")
	fmt.Println("Start")
	for i := 0; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Finished")
}

// Anonymous goroutines
// Format
// go func(){ //body }(args..)
func anonymousGoroutines() {
	fmt.Println("\nAnonymous Goroutines")
	go func() {
		fmt.Println("In Goroutine")
	}()

	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}
