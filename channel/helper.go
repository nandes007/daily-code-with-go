package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finised")
	ch <- 1
}

func receive(ch chan int) {
	val := <-ch
	fmt.Printf("Receiving value from channel finised. Value received: %d\n", val)
}

func sum(ch chan int, len int) {
	sum := 0
	for i := 0; i < len; i++ {
		sum += <-ch
	}
	fmt.Printf("Sum: %d\n", sum)
}

func sumTwo(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}
