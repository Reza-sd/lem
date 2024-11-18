package main

import (
	"fmt"
	"math/rand"
)

//=========================
func generateBoringChannel(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 500; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				//do nothing
			case <-quit:
				//cleanup()
				quit <- "See you!"
				return
			}
		}
	}()
	return c
}

//=========================
func main() {
	quit := make(chan string) //as a quit signal

	c := generateBoringChannel("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

//=========================
