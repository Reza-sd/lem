package main

import (
	"fmt"
)

//===============================
func main() {
	//------------------------------------------------------
	ch := make(chan int) // create a channel

	// Start a goroutine that sends a value into the channel
	go func() {
		ch <- 42 // send value to channel
		ch <- 78
	}()

	// Receive the value from the channel
	val := <-ch
	fmt.Println("Received 1:", val)  // Output: Received: 42
	fmt.Println("Received 2:", <-ch) // Output: Received: 78
	//------------------------------------------------------
	/*
		Buffered Channels
		By default, channels are unbuffered, meaning they have no capacity, and sending/receiving operations block until the other side is ready. However, Go also supports buffered channels, where you can specify a capacity during creation.

		With buffered channels:

		Sending only blocks if the buffer is full.
		Receiving only blocks if the buffer is empty.
	*/
	//------------------------------------------------------
	ch_buffer := make(chan int, 2) // buffered channel with capacity 2
	ch_buffer <- 1
	ch_buffer <- 2

	// You can receive the values without blocking
	fmt.Println(<-ch_buffer) // Output: 1
	fmt.Println(<-ch_buffer) // Output: 2

	//------------------------------------------------------
	//Closing Channels: You can close a channel to indicate that no more values will be sent. Closing a channel is useful for signaling completion.

	close(ch_buffer)
	//ch_buffer <- 3 //(panic: send on closed channel)
	//------------------------------------------------------

	//Receiving from a closed channel will yield zero values (depending on the type) after all buffered values are read.

	fmt.Println("after close channel", <-ch_buffer) // Output: 2

	//------------------------------------------------------

	//------------------------------------------------------

	//------------------------------------------------------

	//------------------------------------------------------

	//------------------------------------------------------

	//------------------------------------------------------

	//------------------------------------------------------

}

//===============================
