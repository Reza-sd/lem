package main

import (
	"fmt"
	"time"
)

//====================================
// goroutine to send data to the channel
func set_number_to_channel(number chan int) {
	number <- 15
}

//======================================
// goroutine to send data to the channel
func set_msg_to_channel(message chan string) {

	// sleeps the process by 2 seconds
	time.Sleep(2 * time.Second)

	message <- "Learning Go Select"
}

//======================================
func main() {

	// create channels
	number_channel := make(chan int)
	message_channel := make(chan string)

	// function call with goroutine
	go set_number_to_channel(number_channel)
	go set_msg_to_channel(message_channel)

	// selects and executes a channel
	select {
	case firstChannel := <-number_channel:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message_channel:
		fmt.Println("Channel Data:", secondChannel)
	}

}

//======================================
