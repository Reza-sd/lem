package main

import (
	"fmt"
	"time"
)

//==============================================
func producer(ch chan<- int, count int) {
	//----------------
	for i := 1; i <= count; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(time.Millisecond * 500) // Simulate work
	}
	//----------------
	close(ch)
}

//===================
func consumer(ch <-chan int, id int) {
	//----------------
	for val := range ch {
		fmt.Printf("Consumer %d consumed %d\n", id, val)
		time.Sleep(time.Millisecond * 1000) // Simulate work
	}
	//----------------
}

//===================
func main() {
	//----------------
	bufferSize := 3
	ch := make(chan int, bufferSize)
	//----------------
	go producer(ch, 5)
	//----------------
	// Start multiple consumers
	for i := 1; i <= 2; i++ {
		// consumer(ch, i)
		go consumer(ch, i)
	}
	//----------------
	// Wait for consumers to finish
	time.Sleep(time.Second * 5)
	//----------------
}

//==============================================
