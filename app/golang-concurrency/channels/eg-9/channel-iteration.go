package main

import "fmt"

//=====================================
/*
1. Iterating Over Channels

Go provides a concise way to iterate over channels using a for-range loop. This loop automatically receives values from the channel until it's closed. Here's a basic example:

2. Closing Channels

Closing a channel signals that no more values will be sent on it. This is essential for two reasons:

Iteration Termination: When a channel is closed, the for-range loop terminates gracefully.
Error Handling: A closed channel can be used to propagate errors or signals between goroutines.

Closing a channel in Go is important when signaling to the receiver(s) that no more values will be sent. Only the sender should close a channel—receivers should never close it.
*/

//===================================
func main() {
	//----------------
	c := make(chan int)
	// Start a goroutine to send data
	//--------------------
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // Send values into the channel
		}
		close(c) // Close the channel once done
	}()
	//-------------------
	// Range over the channel to receive values

	for v := range c {
		fmt.Println(v) // Print values received from the channel
	}
	//--------------------------
}
