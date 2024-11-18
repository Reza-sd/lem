package main

import (
	"fmt"
	"math/rand"
	"time"
)

//==========================
func boring(msg string, c chan string) {

	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		//rand.Intn(1e3) is a common expression used to generate a random integer between 0 (inclusive) and 1000 (exclusive)
		//1e3 = This is a shorthand notation for 10 raised to the power of 3, which equals 1000.

	}
}

//===============================
func main() {
	c := make(chan string)

	go boring("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.

	}

	fmt.Println("You are boring; I'm leaving")

}

//===============================
