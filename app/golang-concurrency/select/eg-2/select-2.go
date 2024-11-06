package main

import (
	"errors"
	"time"
)

//============================
func operation() error {
	println("Do operation...")
	time.Sleep(time.Second * 2)
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

	doSomething(time.Second * 3)
}
