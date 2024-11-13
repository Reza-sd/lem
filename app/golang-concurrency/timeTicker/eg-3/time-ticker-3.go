package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Stop the ticker when done to free resources
	//Always call ticker.Stop() when you’re done with a ticker. This stops the ticker from running in the background and consuming resources.

	done := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	//loop: This loop listens on multiple channels. When the done channel is signaled, the loop exits. Otherwise, on every tick, it prints the current time.
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}
