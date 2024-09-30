package main

import (
	"fmt"
	"sync"
)

/*
The Singleton Pattern is a creational design pattern that ensures a class (or a struct, in Go) has only one instance and provides a global point of access to that instance. This is useful when you want to control access to shared resources, such as a database connection or a configuration manager, ensuring that only one instance exists throughout the program's lifecycle.

In Go, the Singleton pattern can be implemented using:

sync.Once to ensure thread-safe initialization.
Package-level variables to manage the single instance.

//--------------
Example: Logger Singleton
In this example, we'll implement a Logger struct that follows the Singleton pattern. This logger will have a single instance, and we will ensure that only one instance exists, even if multiple goroutines access it concurrently.

//--------------
*/
//==============================================
//1. Define the Singleton Struct (Logger):
// Logger is the singleton struct
type Logger struct {
	LogLevel string
}

var loggerInstance *Logger
var once sync.Once

// GetLoggerInstance returns the singleton instance of Logger
func GetLoggerInstance() *Logger {
	once.Do(func() {
		loggerInstance = &Logger{LogLevel: "INFO"}
		fmt.Println("Creating a new logger instance.")
	})
	return loggerInstance
}

// Log prints a message with the current log level
func (l *Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.LogLevel, message)
}

//==============================================
func main() {
	// First call to get the logger instance
	logger1 := GetLoggerInstance()
	logger1.Log("This is the first log message.")

	// Second call to get the logger instance (should return the same instance)
	logger2 := GetLoggerInstance()
	logger2.Log("This is the second log message.")

	// Check if both logger1 and logger2 point to the same instance
	if logger1 == logger2 {
		fmt.Println("logger1 and logger2 are the same instance.")
		println()
	}
	//===============================================
	/*
			Thread-Safe Singleton with Multiple Goroutines:
		If multiple goroutines call GetLoggerInstance() concurrently, sync.Once guarantees that the initialization block is executed only once, ensuring thread-safety.
	*/
	// Simulate multiple goroutines trying to access the singleton
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger := GetLoggerInstance()
			logger.Log(fmt.Sprintf("Log message from goroutine %d", i))
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	//=================================
}

//==============================================
