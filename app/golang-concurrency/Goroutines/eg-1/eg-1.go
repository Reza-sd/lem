package main

/*
Select statements in Go are inherently thread-safe.



*/
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 42
	}()

	go func() {
		ch2 <- "hello"
	}()

	select {
	case v := <-ch1:
		println("Received from ch1:", v)
	case v := <-ch2:
		println("Received from ch2:", v)
	}
}
