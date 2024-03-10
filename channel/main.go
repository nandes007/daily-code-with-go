package main

func main() {
	// delareChannel()

	// sendAndReceive()

	// timeout()

	//Send on a channel is blocked when the channel is full
	//This will raise: fatal error: all goroutines are asleep - deadlock!
	// ch := make(chan int, 1)
	// ch <- 1
	// ch <- 1
	// fmt.Println("Sending value to channel complete")
	// val := <-ch
	// fmt.Printf("Receiving value from channel finised. Value received: %d\n", val)

	// onlyReadChannel()

	onlyWriteChannel()

	capacityChannel()

	lengthChannel()

	closeChannel()

	checkChannel()

	forRangeChannel()

	nilChannel()
}
