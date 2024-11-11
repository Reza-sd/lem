package main

import (
	"fmt"
)

//====================================
/*
The sender can send two values (1 and 2) without blocking because the buffer capacity is 2.
When attempting to send the third value (3), the sender blocks until a receiver retrieves a value from the buffer.
The receiver processes values in the order they were sent.
*/
//====================================
func main() {
	//----------------------------
	ch := make(chan int, 2) // Buffered channel with capacity 2
	//----------------------------
	// Sender goroutine
	go func() {
		fmt.Println("Sending value 1")
		ch <- 1
		fmt.Println("Sent value 1")

		fmt.Println("Sending value 2")
		ch <- 2
		fmt.Println("Sent value 2")

		fmt.Println("Sending value 3")
		ch <- 3 // Blocks since buffer is full
		fmt.Println("Sent value 3")
	}()
	//----------------------------
	// Receiver
	fmt.Println("Receiving first value")
	val1 := <-ch
	fmt.Println("Received:", val1)

	fmt.Println("Receiving second value")
	val2 := <-ch
	fmt.Println("Received:", val2)

	fmt.Println("Receiving third value")
	val3 := <-ch
	fmt.Println("Received:", val3)
	//----------------------------
}

//====================================
