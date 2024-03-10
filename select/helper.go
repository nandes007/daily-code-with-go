package main

import (
	"fmt"
	"time"
)

func goOne(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "From goTwo gorotine"
}

func goTwoSend(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func printAllNews(news chan string) {
	for {
		select {
		case n := <-news:
			fmt.Println(n)
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout: News feed finised")
			return
		}
	}
}

func printAllNewsWithNillSelectStatement(news chan string) {
	for {
		select {
		case n := <-news:
			fmt.Println(n)
			news = nil // select statement with nil channel.
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout: News feed finised")
			return
		}
	}
}

func newsFeed(ch chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Microsecond * 400)
		ch <- fmt.Sprintf("News: %d", i+1)
	}
}
