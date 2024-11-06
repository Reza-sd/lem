package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)

	// Send values to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// Close the channel to signal the end
		close(ch)
	}()

	// Range over the channel to receive and print values
	for num := range ch {
		fmt.Println(num)
	}
}
