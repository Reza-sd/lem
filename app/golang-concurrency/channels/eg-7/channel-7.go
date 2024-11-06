package main

import "time"


//==========================
func bidirectional(ch chan int) {
	println("start")

	ch <- 999
	value := <-ch
	println("value=",value)
}

//==================================
func main() {

	ch := make(chan int)
	go bidirectional(ch)

	time.Sleep(time.Second * 1)
}

//==================================
