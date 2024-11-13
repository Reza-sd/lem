package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   Immediate First Tick
	   To execute an action immediately and then at regular intervals:
	   This pattern allows for an immediate first execution followed by regular ticks.

	*/
	ticker := time.NewTicker(1 * time.Second)
	for t := time.Now(); ; t = <-ticker.C {
		fmt.Printf("Tick at: %v\n", t.UTC())
		// Perform action here
	}

}
