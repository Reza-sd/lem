package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic timeout example
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: no message received after 2 seconds")
	}

	// Multiple operations with timeout
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Hello!"
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println("Received:", msg)
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Still waiting...")
		}
	}
}
