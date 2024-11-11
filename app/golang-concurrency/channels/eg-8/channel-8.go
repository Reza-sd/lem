package main

import (
	"fmt"
	//"time"
)

//=====================================
/*
Why Use Channel Directions?
- Improved type safety: By specifying the direction of a channel, the compiler can enforce correct usage, preventing accidental sends or receives.

- Enhanced code clarity: Explicitly declaring channel directions makes code more readable and easier to understand.

- Better concurrency control: Unidirectional channels can be used to create more complex synchronization patterns and avoid race conditions.
*/
//=========================================
func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch)
}

//=========================================
func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
}

//=========================================
func main() {
	// Create a bidirectional channel
	ch := make(chan int)

	// Start two goroutines to send and receive data
	go sender(ch)
	go receiver(ch)

	// Wait for the goroutines to finish
	select {
	case <-ch:

	}
	//time.Sleep(5 * time.Second)
}

//=========================================
/*
```

This example demonstrates how to create a bidirectional channel using the `chan` keyword without any directional annotations. The `sender` function sends values to the channel, and the `receiver` function receives values from the channel.

Here's how the code works:

1. The `main` function creates a bidirectional channel `ch` using `make(chan int)`.
2. It then starts two goroutines: `sender` and `receiver`, both of which have access to the `ch` channel.
3. The `sender` function sends values to the channel, one by one, using the `<-` operator. It then closes the channel when it's done sending.
4. The `receiver` function receives values from the channel using the `for val := range ch` loop. It prints the received values.
5. The `main` function waits for 5 seconds to allow the goroutines to finish their work.

When you run this code, you should see the following output:

```
Sent: 0
Received: 0
Sent: 1
Received: 1
Sent: 2
Received: 2
Sent: 3
Received: 3
Sent: 4
Received: 4
```

The key points to note in this example are:

- The `chan` keyword is used to create a bidirectional channel.
- The `<-` operator is used to send and receive values from the channel.
- The `sender` function sends values to the channel, while the `receiver` function receives values from the channel.
- The `close(ch)` statement closes the channel when the sender is done.

Let me know if you have any other questions!
*/
