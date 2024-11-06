package main

func m(a int, b int) int {
	c := make(chan int)
	d := 2

	go func() {

		c <- d
	}()

	select {
	case res := <-c:
		println("Received result:", res)
		return d
	default:
		println("Timed out")
		return 0
	}

}

func main() {
	println(m(4, 2))

}
