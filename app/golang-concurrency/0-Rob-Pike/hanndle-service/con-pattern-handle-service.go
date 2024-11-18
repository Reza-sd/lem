package main

import (
	"fmt"
	"math/rand"
	"time"
)

//============================
/*
Channels as a handle on a service |
Our boring function returns a channel that lets us communicate with the |
boring service it provides. |
We can have more instances of the service. |
*/
//==================================
//Channels are first-class values, just like strings or integers.
//==========================
//function returning a channel.
// returns recive-only channel of strings.

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

//===============================
func main() {

	//function returning a channel.
	joe := boringChannel("Joe!")
	ann := boringChannel("Ann!")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)

	}

	fmt.Println("You are boring; I'm leaving")

}

//===============================
