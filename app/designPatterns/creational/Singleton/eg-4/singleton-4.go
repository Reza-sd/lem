package main

import (
	"fmt"
	"sync"
)

// =================================
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

// -------------------------
func getInstance(i int) *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	print(i, "-")
	return singleInstance
}

//=================================

func main() {

	for i := 0; i < 30; i++ {
		//print(i)
		go getInstance(i)
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

//=================================
