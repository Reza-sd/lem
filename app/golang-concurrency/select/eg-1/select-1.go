package main

func m() {
	c := make(chan int)
	//d := 2

	go func() {

		c <- 0
	}()

	select {
	case res := <-c:
		println("Received result:", res)

	default:
		println("Timed out")

	}

}

func main() {
	m()

}
