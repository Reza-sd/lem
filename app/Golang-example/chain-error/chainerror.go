package main

import (
	"errors"
	"fmt"
)

type chainError struct {
	msg  string
	prev error
}

func (e chainError) Error() string {
	if e.prev != nil {
		return e.msg + ": " + e.prev.Error()
	}
	return e.msg
}

func ChainError(msg string, err error) error {
	return chainError{msg: msg, prev: err}
}

func main() {
	err1 := errors.New("first error")
	err2 := ChainError("second error", err1)
	err3 := ChainError("third error", err2)

	fmt.Println(err3)
}