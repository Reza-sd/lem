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

func sender(ch chan<- string) {
	for i := 1; i <= 4; i++ {
		ch <- fmt.Sprintf("Msg %v ,send at time: %v", i, time.Now().UnixMilli())
		time.Sleep(time.Millisecond * 150)
	}
	close(ch)
}

//==================================
//Receive-only channels : These channels can only be used to receive values. They are declared as <-chan T.

func receiver(ch <-chan string) {
	for v := range ch {
		fmt.Println(v, ",receive at time", time.Now().UnixMilli())
		time.Sleep(time.Millisecond * 500)
	}
}

//==========================
// func bidirectional(ch chan int) {
// 	println("start")

// 	ch <- 999
// 	value := <-ch
// 	fmt.Println(value)
// 	println("END: bidirectional")
// }

//==================================
func main() {
	//ch1 := make(chan int)
	// go bidirectional(ch1)

	ch := make(chan string)
	go sender(ch)

	receiver(ch)

	//time.Sleep(time.Second*1)

	//ch1 := make(chan int)
	//go bidirectional(ch1)

	time.Sleep(time.Second * 1)
}

//==================================
