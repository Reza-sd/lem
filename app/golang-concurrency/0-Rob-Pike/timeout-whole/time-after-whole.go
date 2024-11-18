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
	c := boringChannel("Joe")
	start := time.Now()

	//Timeout for whole conversation using select
	timeout := time.After(1 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case t := <-timeout:
			 //Create the timer once, outside the loop, to time out the entire conversation. (In the previous program, we had a timeout for each message.)
			
			fmt.Println("You are too slow. time:", t)
			// Calculate elapsed time
			elapsed := time.Since(start)
			fmt.Printf("Elapsed time: %v\n", elapsed)
			return
		}
	}
}
//=========================
