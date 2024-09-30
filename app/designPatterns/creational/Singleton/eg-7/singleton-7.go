package main

import (
	"fmt"
	"sync"
)

//========================================================
// Singleton struct
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

//-----------------------------------
// GetInstance returns the single instance of Singleton
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "Initial data"}
	})
	return instance
}

// SetData sets the data field of the Singleton
func (s *Singleton) SetData(data string) {
	s.data = data
}

// GetData returns the data field of the Singleton
func (s *Singleton) GetData() string {
	return s.data
}

//========================================================
func main() {
	// Get the Singleton instance
	instance1 := GetInstance()
	fmt.Printf("Instance 1 initial data: %s\n", instance1.GetData())

	// Modify the Singleton data
	instance1.SetData("Modified data")

	// Get the Singleton instance again
	instance2 := GetInstance()
	fmt.Printf("Instance 2 data: %s\n", instance2.GetData())

	// Check if both instances are the same
	fmt.Printf("Are instance1 and instance2 the same? %v\n", instance1 == instance2)

	// Demonstrate concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			instance := GetInstance()
			fmt.Printf("Goroutine %d: instance address: %p\n", id, instance)
		}(i)
	}
	wg.Wait()
}

//========================================================
