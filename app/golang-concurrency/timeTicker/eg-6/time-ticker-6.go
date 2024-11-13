package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

//================================
func safeNewTicker(d time.Duration) (*time.Ticker, error) {
	// Create a new ticker with the given duration.
	if d <= 0 {
		return nil, fmt.Errorf("non-positive interval for NewTicker")
	}
	return time.NewTicker(d), nil
}

//====================================
func doSomething() error {
	// Implement your task logic here
	// ...
	println("Do Something...")
	return nil // Return nil if successful, error otherwise
}

//=====================================
func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	//------------------------
	ticker, err := safeNewTicker(1 * time.Second)
	if err != nil {
		log.Fatal("Failed to create ticker:", err)
	}
	defer ticker.Stop()
	//------------------------
	for {
		//---------------------
		select {
		case <-ticker.C:
			if err := doSomething(); err != nil {
				log.Printf("Error in task: %v", err)
			}
		case <-ctx.Done():
			log.Println("Context canceled")
			return
		}
		//------------------------
	}
	//------------------------
}

//=====================================
