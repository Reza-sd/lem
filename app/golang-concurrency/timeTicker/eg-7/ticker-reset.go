package main

import (
	"fmt"
	"time"
)

//===============================
func doSomething(t *time.Ticker) {
	for range t.C {
		fmt.Println("Ticking...", time.Now())
	}
}

//===============================
func main() {
	fmt.Println("Start App: ", time.Now())
	ticker := time.NewTicker(time.Second * 2) // Initial 2-second interval

        //Stopping the Ticker: Always remember to stop the ticker when you're done with it to free up resources:
	defer ticker.Stop()

	go doSomething(ticker)

	time.Sleep(7 * time.Second) // Wait for 5 seconds
	fmt.Println("Wait for 5 seconds...")

	ticker.Reset(time.Second * 1) // Reset the interval to 1 seconds
	time.Sleep(5 * time.Second)
	fmt.Println("END.")

}

//===============================
