package main

import "fmt"

func start() {
	go start2()
	fmt.Println("In Goroutine")
}

func start2() {
	fmt.Println("In Goroutine 2")
}

func execute(id int) {
	fmt.Printf("id: %d\n", id)
}
