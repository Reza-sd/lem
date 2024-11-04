package main

import (
	"time"
)

//===========================================
func say(s string, t time.Duration, isGoroutine bool) {
	for i := 0; i < 5; i++ {
		//fmt.Println(i,s)
		time.Sleep(t * time.Millisecond)
		if isGoroutine {
			print(i, "<-", s, "=", t, " ")
		} else {
			print(i, "-", s, "=", t, " \n")
		}

	}
}

//===========================================
func main() {
	//defer println("\n <<END from main>>")
	// defer func() {
	//     println("\n <<END from main>>")
	// }()
	say("mio", 100, false)

	go say("world", 100, true) //will run in a separate goroutine.

	say("hello", 250, false)

	go say("woof", 100, true)

	time.Sleep(2 * time.Second) // Wait for goroutine to finish
	println("\n\n <<END from main>>")
}

//===========================================
