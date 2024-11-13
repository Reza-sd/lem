package main

import (
	"fmt"
	"time"
)

/*
In Go, time.After and select are powerful tools for managing asynchronous operations and timeouts.

*/
func main() {
	//time.After= This function returns a channel that will receive a time.Time value after a specified duration. It's commonly used to set a timeout for operations or to schedule events in the future.

	timeout := time.After(2 * time.Second) //creates a channel that will receive a value after 2 seconds.

	done := make(chan bool) //a channel used to signal when the long-running operation is complete.

	// Simulate a long-running operation
	go func() {
		time.Sleep(3 * time.Second)
		done <- true
	}()

	//The select statement allows you to wait on multiple channels. It will select the first channel that is ready to receive or send a value. This is often used to implement non-blocking I/O operations and to handle multiple concurrent tasks.

	select {
	//The select statement waits for either the timeout channel or the done channel to receive a value.
	case <-timeout:
		fmt.Println("Timed out")
	case <-done:
		fmt.Println("Done")
	}

	//Context Cancellation: For more complex scenarios, the context package provides a mechanism for canceling operations and propagating cancellation signals.

}
