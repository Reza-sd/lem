package main

import (
	"time"
)

//==============
func checkSomething() {
	println("checking...")
}

//call a function at regular times
//===============
func main() {
	intervalDuration := time.Duration(1) * time.Second

	ticker := time.NewTicker(intervalDuration)
	//print(ticker.)
	i := 0
	for range ticker.C {
		i++
		checkSomething()

		if i > 5 {
			ticker.Stop()
			break
		}
	}
	println("END . iteration =", i)
}

//===============
