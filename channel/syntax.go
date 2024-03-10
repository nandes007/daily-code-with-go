package main

import (
	"fmt"
	"time"
)

// Channel
// Channel is a data type in Go which provides synchrounization and communication between goroutines.
// Signature: var variable_name chan type
// Note: Length and Capacity of unbeffered channel is zero
func delareChannel() {
	fmt.Println("Declare channel")
	var a chan int
	a = make(chan int) //define channel
	fmt.Println(a)
}

// Operations on channel
// There are two major operations which can be done on a channel
// Send and Receive
// Send is used to send data to channel: ch <- val
// Receive is used to read data from channel: val := <- ch
func sendAndReceive() {
	fmt.Println("\nSend and receive data from one goroutine")
	ch := make(chan int)

	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving from channel")
	go receive(ch)

	time.Sleep(time.Second * 1)
}

func timeout() {
	fmt.Println("\nIllustrate blocking on send and timeout on receive")
	ch := make(chan int)
	go send(ch)

	go receive(ch)
	time.Sleep(time.Second * 2)
}

// Make buffered channel
// a = make(chan, capacity)
func bufferedChannel() {
	fmt.Println("\nBuffered channel")
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("Sending value to channel complete")
	val := <-ch
	fmt.Printf("Receiving Value from channel finised. Value received: %d\n", val)
}

// Channel Direction
// - chan :bidirectional channel(Both read and write)
// - chan <- :only writing to channel
// - <-chan :only reading from channel(input channel)
func read(ch chan<- int) {
	ch <- 2
}

func write(ch <-chan int) {
	s := <-ch
	fmt.Println(s)
}

func onlyReadChannel() {
	fmt.Println("\nOnly read channel")
	ch := make(chan int, 3)
	read(ch)
	fmt.Println(<-ch)
}

func onlyWriteChannel() {
	fmt.Println("\nOnly write channel")
	ch := make(chan int, 3)
	ch <- 2
	write(ch)
}

// Capacity of channel using cap() function
// The capacity of a buffered channel is the number of elements which that channel can hold.
// Capacity refers to the size of the buffer of the channel.
func capacityChannel() {
	fmt.Println("\nCapacity channel")
	ch := make(chan int, 3)
	fmt.Printf("Capacity: %d\n", cap(ch))
}

// Lenth of channel using len() function
// Builtin len() function can be used to get the length of a channel.
// The length of a channel is the number of elements that are already there in the channel.
func lengthChannel() {
	fmt.Println("\nLength of channel")
	ch := make(chan int, 3)
	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))
	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
}

// Close operation on channel
// Close is an inbuilt function that can be used to close a channel.
// Closing of a channel means that no more data can we send to the channel.
func closeChannel() {
	fmt.Println("\nClose channel")
	ch := make(chan int)
	go sum(ch, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	time.Sleep(time.Second * 1)
}

// Check the channel has been close or not
// Signature: val, ok <- ch
// The value of ok will be true if channel is not close, false if the channel is closed.
func checkChannel() {
	fmt.Println("\nCheck channel")
	ch := make(chan int, 1)
	ch <- 2
	val, ok := <-ch
	fmt.Printf("Val: %d OK: %t\n", val, ok)

	close(ch)
	val, ok = <-ch
	fmt.Printf("Val: %d OK: %t\n", val, ok)
}

// For rage loop on a channel
// For range loop can be used to receive data from the channel until it closed.
func forRangeChannel() {
	fmt.Println("\nFor Range Channel")
	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	go sumTwo(ch)
	time.Sleep(time.Second * 1)
}

// Nil channel
// The zero value of the channel is nil.
// Hence only declaring a channel creates a nil channel as default zero value of the channel is nil.
func nilChannel() {
	fmt.Println("\nNil Channel")
	var a chan int
	fmt.Print("Default zero value of channel: ")
	fmt.Println(a)
}
