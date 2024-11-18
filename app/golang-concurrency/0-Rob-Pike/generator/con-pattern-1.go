package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	c := boringChannel("Boring!")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)

	}

	fmt.Println("You are boring; I'm leaving")

}

//===============================
