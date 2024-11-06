package main

import (
	"fmt"
)

func main() {
	// Create an unbuffered channel to send integers
	ch := make(chan int)

	// Start a goroutine to send values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel to signal the end
	}()

	// Iterate over the values received from the channel
	// The loop continues until the channel is closed, indicating that no more values are available.
	for v := range ch {
		fmt.Println(v)
	}
}
