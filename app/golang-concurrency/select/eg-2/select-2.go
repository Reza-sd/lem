package main

import (
	"errors"
	"time"
)

//============================
func operation() error {
	duration := 5 //sec
	println("Doing operation...it takes:", duration, "second")
	time.Sleep(time.Second * time.Duration(duration))
	println("END")
	return nil
}

//=========================
func doSomething(timeout time.Duration) error {
	ch := make(chan error)
	go func() {
		// Perform the operation
		err := operation()
		ch <- err
	}()

	select {
	case err := <-ch:
		return err
	case <-time.After(timeout):
		return errors.New("timeout")
	}
}

//=================================================
func main() {

	err := doSomething(time.Second * 3)
	if err != nil {

		println(err.Error())
	}
}

//=================================================
