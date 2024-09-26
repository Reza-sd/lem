package main

type Logger interface {
	Log(message string)
}

type Service interface {
	DoSomething() error
}

type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) {
	println(message)
}

// -------
type MyService struct {
	logger Logger
}

func (s *MyService) DoSomething() error {
	s.logger.Log("Doing something")
	// Implementation
	return nil
}

// ------------
func NewMyService(logger Logger) Service {
	return &MyService{logger: logger}
}

// Usage
func main() {
	logger := &ConsoleLogger{}
	service := NewMyService(logger)
	service.DoSomething()
}
