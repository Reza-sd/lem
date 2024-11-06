package main
import "time"

//==========================
func bidirectional(ch chan int) {

	println("set")
	ch <- 999

}
//==========================
func reciever(ch chan int){
	println(<-ch)
}
//==================================
func main() {

	ch := make(chan int)
	go bidirectional(ch)
	go reciever(ch)

	time.Sleep(time.Second * 1)
}

//==================================
