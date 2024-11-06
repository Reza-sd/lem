package main

import (
	"errors"
	"time"
)

//==========================
func fetchFromRemoteServer() (string, error) {

	duration := time.Duration(10) //sec
	print("fetching data From Remote Server...wait...")
	//println(time.Duration(duration))
	time.Sleep(time.Second * duration)

	println("END")

	return "name:Reza", nil
}

//=================================
func fetchData() (result string, err error) {
	c := make(chan string)
	errc := make(chan error)

	go func() {
		// Fetch data asynchronously
		data, err := fetchFromRemoteServer()
		if err != nil {
			errc <- err
		} else {
			c <- data
		}
	}()

	select {
	case <-time.After(3 * time.Second):
		return "", errors.New("timeout")
	case err := <-errc:
		return "", err
	case result := <-c:
		return result, nil
	}
}

//=====================================
func main() {

	data, err := fetchData()
	if err != nil {

		println(err.Error())
	} else {
		println("successfully fetched data")
		println(data)
	}
}

//=============================
