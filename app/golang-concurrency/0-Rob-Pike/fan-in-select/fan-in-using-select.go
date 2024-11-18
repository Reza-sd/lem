package main

import (
	"fmt"
	"math/rand"
	"time"
)

//=========================
type Message struct {
	str  string
	wait chan bool
}

//==============================================
func boringChannel(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)

	go func() {

		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-waitForIt

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

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	// go func() {
	// 	for {
	// 		//Reads a value from input1 and sends it to the output channel c.
	// 		//This process continues indefinitely.
	// 		c <- <-inputl
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c <- <-input2
	// 	}
	// }()

	//Fan-in using select
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c

}

//=========================
func main() {
	c := fanIn(boringChannel("Joe"), boringChannel("Ann"))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		msg1.wait <- true
		msg2.wait <- true

	}

	fmt.Println("You're both boring; I'm leaving.")

}

//=========================
