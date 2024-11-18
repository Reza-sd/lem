package main

import (
	"fmt"
	"math/rand"
	"time"
)

//==============================================
func boringChannel(msg string) <-chan string {
	c := make(chan string)

	go func() {

		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

//============================================
//a func which recieve 2 channel and return 1 channel. (generator pattern)
/*
fanIn that merges two input channels into a single output channel. It's a common pattern in concurrent programming, often used to combine results from multiple goroutines into a single stream.


*/

// returns recive-only channel of strings.
//Mergers and Acquisition: The Power of Aggregator Strategy

func fanIn(inputl, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			//Reads a value from input1 and sends it to the output channel c.
			//This process continues indefinitely.
			c <- <-inputl
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c

}

//=========================
func main() {
	c := fanIn(boringChannel("Joe"), boringChannel("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)

	}
	fmt.Println("You're both boring; I'm leaving.")

}

//=========================
