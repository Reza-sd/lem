package main

import (
	"fmt"
)

//=============================
/*

Goroutines are started when you call them with the go keyword. They do not have a mechanism to be stopped from outside—they run until:

The function finishes.
The main program ends (if the main program exits, all goroutines exit).

- Synchronization with Goroutines:

When you have multiple goroutines, you often need them to communicate and synchronize. Go provides channels and sync primitives for this purpose.

Channels
Channels allow you to send and receive values between goroutines, making synchronization straightforward.
*/
//============================
func printMessage(message string, ch chan string) {
	ch <- message // Send message to channel
}

//============================
func main() {
	ch := make(chan string)

	go printMessage("Hello, Goroutines!", ch)

	// Receive message from channel
	msg := <-ch
	fmt.Println(msg)
}

/*
In this example:

printMessage sends a string to the ch channel.
The main function receives the message from ch and prints it.
*/
//============================
