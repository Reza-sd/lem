package main

import (
	"fmt"
	"time"
)

//===============================================
func fetchWithTimeout(url string, timeout time.Duration) error {
	ch := make(chan error)
	println("Start fetching url...", url)
	startTime := time.Now()

	go func() {
		// Do some work
		_ = url
		time.Sleep(2 * time.Second)
		ch <- nil
	}()

	select {
	case err := <-ch:
		fmt.Println("Actual duration to do =", time.Since(startTime))
		return err
	case <-time.After(timeout * time.Second):
		return fmt.Errorf("operation timed out")
	}
}

//===============================================
func main() {
	var timeout time.Duration = 3

	err := fetchWithTimeout("www.google.com", timeout)

	if err != nil {
		println("URL not Reachablable")
	} else {
		println("fetching successful")
	}
}

//===============================================
