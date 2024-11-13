package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new ticker that ticks every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Receive ticks from the ticker channel
	start := time.Now()

	for t := range ticker.C {
		TimeAfterStart := time.Since(start)
		fmt.Println("TimeAfterStart=", TimeAfterStart)
		fmt.Println("Tick at:", t)
		if TimeAfterStart > (9 * time.Second) {
			break
		}
	}
}
