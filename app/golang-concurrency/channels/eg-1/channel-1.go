package main

import (
	"time"
)
//======================================
func paint(item string) {
	println(item)
	time.Sleep(150 * time.Millisecond)
}

//======================================
func main() {
	start := time.Now()
	defer func() {
		println("Duration(Milliseconds)=", time.Since(start).Milliseconds())
	}()

	mySlice := []string{"A", "B", "C", "D", "E"}


	for _, ch := range mySlice {
		go paint(ch)
	}

	time.Sleep(time.Second*2)
}
//======================================