package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2) // Ticks every 2 seconds

	go func() {
		for range ticker.C {
			fmt.Println("Tick at:", time.Now())
		}
	}()

	time.Sleep(9 * time.Second) // Let it run for a minute
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
