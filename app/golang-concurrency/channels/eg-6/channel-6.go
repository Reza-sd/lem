package main

import (
	"fmt"
	"time"
)

//==================================
/*
Channel Directions in Golang:
-channels : are a powerful mechanism for communication and synchronization between goroutines.

-direction : They can be unidirectional, allowing data flow in only one direction, or bidirectional, allowing data flow in both directions.


*/
//==================================
//Send-only channels : These channels can only be used to send values. They are declared as chan<- T.

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

//==================================
//Receive-only channels : These channels can only be used to receive values. They are declared as <-chan T.

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

//==========================
func bidirectional(ch chan int) {
	println("start")

	ch <- 999
	value := <-ch
	fmt.Println(value)
	println("END: bidirectional")
}

//==================================
func main() {
	// ch1 := make(chan int)
	// go bidirectional(ch1)

	// ch := make(chan int)
	// go producer(ch)
	// go consumer(ch)

	//time.Sleep(time.Second*1)

	ch1 := make(chan int)
	go bidirectional(ch1)

	time.Sleep(time.Second * 1)
}

//==================================
