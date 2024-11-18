package main

import (
	"fmt"
	"math/rand"
)

//=========================
func generateBoringChannel(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 500; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				//do nothing
			case <-quit:
				return
			}
		}
	}()
	return c
}

//=========================
func main() {
	quit := make(chan bool)
	c := generateBoringChannel("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
}

//=========================
