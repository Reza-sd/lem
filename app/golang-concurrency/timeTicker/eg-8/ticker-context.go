package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context that will cancel after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Make sure to release resources

	// Create a new ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // Ensure that the ticker is stopped to prevent leaks

	// Run a loop that checks both the ticker and the context
	for {
		select {
		case <-ctx.Done():
			// Context was canceled, either due to timeout or explicit cancel
			fmt.Println("Context canceled, stopping ticker.")
			return
		case t := <-ticker.C:
			// Ticker ticked, perform the periodic action
			fmt.Println("Tick at", t)
		}
	}
}
