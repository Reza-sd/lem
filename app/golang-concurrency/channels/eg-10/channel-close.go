package main

import (
	"fmt"
)

//===================================
/*
n this example, the for loop continues until the channel is closed. The ok value in the receive operation indicates whether the channel is still open. If ok is false, the channel is closed.
*/

//===================================
func main() {
	c := make(chan int)
	//------------------
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//------------------
	for {
		v, ok := <-c
		//If ok is false, the channel is closed.
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println(v)
	}
	//------------------
}

//=========================================
