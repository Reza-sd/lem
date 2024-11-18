package main

import (
	"fmt"
	"math/rand"
	"time"
)
//=========================
func boringChannel(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
//=========================
func main() {
	//c := fanIn(boringChannel("Joe"), boringChannel("Ann"))
	c := boringChannel("Joe")
	start := time.Now()
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case t := <-time.After(1 * time.Second): //return a channel that blocks for specified duration. After the interval.the channel delivers the current time,once.
			fmt.Println("You are too slow. time:", t)
			// Calculate elapsed time
			elapsed := time.Since(start)
			fmt.Printf("Elapsed time: %v\n", elapsed)
			return
		}
	}
}
//=========================
