package main

import (
	"fmt"
	"time"
)

func periodicTask(done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	//duration:=time.Duration()
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
	// for range ticker.C {
	// 	// This code runs every second
	// 	fmt.Println("Tick at", time.Now())
	// }

	// for {
	//     select {
	//     case <-ticker.C:
	//         fmt.Println("Doing work...at time=",time.Now())
	//     case <-time.After(5 * time.Second):
	//         fmt.Println("No activity for 5 seconds, exiting...")
	//         return
	//     }
	// }
}

func main() {
	//ticker := time.NewTicker(2 * time.Second) // Ticks every 2 seconds
	done := make(chan bool)

	go periodicTask(done)

	time.Sleep(5 * time.Second) // Run for 10 seconds
	//ticker.Stop()                // Stop the ticker
	done <- true // Signal the goroutine to stop
	fmt.Println("Ticker stopped")

}
