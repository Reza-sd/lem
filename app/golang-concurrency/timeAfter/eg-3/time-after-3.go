package main

import (
	"fmt"
	"time"
)

func main() {
	// Example 1: Basic timeout using time.After
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout occurred after 2 seconds")
	}

	// Example 2: Choose between a channel and a timeout
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Data received"
	}()

	select {
	case result := <-ch:
		fmt.Println("Received:", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out")
	}

	// Example 3: Implementing a timeout for an API call
	done := make(chan bool)
	go func() {
		// Simulated API call
		time.Sleep(1 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("API call completed successfully")
	case <-time.After(2 * time.Second):
		fmt.Println("API call timed out")
	}

	// Example 4: Periodic timeout with multiple cases
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Intermediate timeout")
		}
	}
}

// Example 5: Function with timeout pattern
func withTimeout(operation func() error, timeout time.Duration) error {
	done := make(chan error, 1)

	go func() {
		done <- operation()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(timeout):
		return fmt.Errorf("operation timed out after %v", timeout)
	}
}
