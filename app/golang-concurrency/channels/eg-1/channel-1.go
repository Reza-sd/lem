package main

import (
	"time"
	//"fmt"
)

//======================================
func paint(item string, signal chan bool) {
	println(item)
	time.Sleep(200 * time.Millisecond)
	signal <- true
}

//======================================
func main() {
	//--------------------
	start := time.Now()
	defer func() {
		println("Duration(Milliseconds)=", time.Since(start).Milliseconds())
	}()
	//--------------------
	exitSignalChannel := make(chan bool)
	//--------------------
	mySlice := []string{"A", "B", "C", "D", "E"}
	//--------------------
	for _, ch := range mySlice {
		go paint(ch, exitSignalChannel)
	}
	//--------------------
	//time.Sleep(time.Second * 2)
	println(<-exitSignalChannel)
}

//======================================
