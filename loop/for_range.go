package main

import "fmt"

func main() {
	letters := []string{"a", "b", "c"}

	//With index and value
	fmt.Println("Both index and value")
	for i, letter := range letters {
		fmt.Printf("Index: %d Value:%s\n", i, letter)
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	//Only index
	fmt.Println("\nOnly index")
	for i := range letters {
		fmt.Printf("Index: %d\n", i)
	}

	//Without index and value, Just print array values
	fmt.Println("\nWithout index and value")
	i := 0
	for range letters {
		fmt.Printf("Index: %d Value: %s\n", i, letters[i])
		i++
	}

	fmt.Println("============================")
	//Sample loop string
	sample := "a£b"
	//With index and value
	fmt.Println("\nWith index and value")
	for i, letter := range sample {
		fmt.Printf("Start index: %d Value:%s\n", i, string(letter))
	}

	//Only value
	fmt.Println("\nOnly value")
	for _, letter := range sample {
		fmt.Printf("Value:%s\n", string(letter))
	}

	//Only index
	fmt.Println("\nOnly Index")
	for i := range sample {
		fmt.Printf("Start Index: %d\n", i)
	}

	//For range for map
	fmt.Println("\nFor range for map")
	sample2 := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	//Iterating over all keys and values
	fmt.Println("Both key and value")
	for k, v := range sample2 {
		fmt.Printf("key :%s value :%s\n", k, v)
	}

	//Iterating over only keys
	fmt.Println("\nOnly keys")
	for k := range sample2 {
		fmt.Printf("key :%s\n", k)
	}

	//Iterating over only values
	fmt.Println("\nOnly values")
	for _, v := range sample2 {
		fmt.Printf("value :%s\n", v)
	}

	//For range with a channel
	fmt.Println("\nFor range with a channel")
	ch := make(chan string)
	go pushToChannel(ch)
	for val := range ch {
		fmt.Println(val)
	}
}

func pushToChannel(ch chan<- string) {
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)
}
