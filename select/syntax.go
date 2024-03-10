package main

import (
	"fmt"
	"time"
)

//Select statement
//format:
// select {
// 	case_channel_send_or_receive:
//	// Dosomething
// 	case_channel_send_or_receive:
//	// Dosomething
// 	default:
//	// Dosomething
// }

func implementSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

func selectWithForLoop() {
	fmt.Println("\nWaiting statement with for loop")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func selectSendOperator() {
	fmt.Println("\nSelect with send operator")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwoSend(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case ch2 <- "To goTwo goroutine":
	}
}

func defaultOnSelect() {
	fmt.Println("\nSelect with default case")
	ch1 := make(chan string)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("Default statement executed")
	}
}

// Select with blocking timeout
// Blocking timeout in select can be achieved by using After() function of time package.
// Signature: func After(d Duration) <-chan Time
func selectWithTimeout() {
	fmt.Println("\nSelect with blocking timeout")
	ch1 := make(chan string)
	go goOne(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

// Empty select
func emptySelect() {
	select {}
}

// Select keyword with an infinite for loop outside
// One of the use case of having infinite for loop outside select statement
// could be that you are waiting for multiple operations to receive on the same channel for a certain time.
func selectInfiniteLoop() {
	fmt.Println("\nSelect statement with infinite for loop outside")
	news := make(chan string)
	go newsFeed(news)

	printAllNews(news)
}

// Select statement with a nil channel
func selectStatementWithNilChannel() {
	fmt.Println("\nSelect statement with nil channel")
	news := make(chan string)
	go newsFeed(news)
	printAllNewsWithNillSelectStatement(news)
}

// Break keyword in select
func breakKeyword() {
	fmt.Println("\nBreak Keyword")
	ch := make(chan string, 1)
	ch <- "Before break"

	select {
	case msg := <-ch:
		fmt.Println(msg)
		break
		fmt.Println("After break")
	default:
		fmt.Println("Default case")
	}
}
