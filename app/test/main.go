package main

import "fmt"

// Define an interface
type Writer interface {
	Write(data string)
}

// Define a struct implementing the interface
type ConsoleWriter struct{}

// Implement the Write method
func (cw ConsoleWriter) Write(data string) {
	fmt.Println(data)
}

// Define a struct embedding the interface
type Logger struct {
	Writer // Embedding the Writer interface
}

func main() {
	logger := Logger{ConsoleWriter{}}
	logger.Write("Hello, Go!") // Output: Hello, Go!

}
