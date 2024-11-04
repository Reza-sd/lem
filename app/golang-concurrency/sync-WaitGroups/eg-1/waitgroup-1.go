package main

import (
	"fmt"
	"sync"
)

//===============================
/*
WaitGroups
Another way to synchronize is by using a sync.WaitGroup, which allows the main goroutine to wait until other goroutines have completed.
*/
//==================================
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Signal WaitGroup that this goroutine is done
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

//==================================
func main() {
	var wg sync.WaitGroup
	wg.Add(1) // Add a goroutine to the WaitGroup

	go printNumbers(&wg)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished.")
}

/*
. Tips and Gotchas
Don't Forget to Wait: The main function exits as soon as all non-daemon code is finished. If you forget to synchronize, goroutines may not finish.
Avoid Race Conditions: When sharing data, use channels or mutexes to avoid race conditions, as goroutines run concurrently and may access shared data simultaneously.
*/
//==================================
