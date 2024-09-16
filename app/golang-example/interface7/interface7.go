package main
/*
This example demonstrates several principles of loosely coupled code:

1. Interface-based design: We define interfaces for `Logger`, `DataStore`, and `Service`.

2. Dependency Injection: The `MyService` struct takes `Logger` and `DataStore` as dependencies.

3. Composition: The `App` struct composes the `Service` interface.

4. Use of context: We use `context.Context` for cancellation and timeouts.

5. Error handling: Errors are returned and wrapped with additional context.

6. Functional options: The `NewApp` function uses functional options for configuration.

7. Small, focused interfaces: Each interface defines only the methods it needs.

8. Package organization: While this is all in one file for brevity, in a real application, these components would likely be in separate packages.

This design allows for easy testing and modification. For example, you could easily swap out the `InMemoryStore` for a database implementation, or replace the `SimpleLogger` with a more sophisticated logging solution, without changing the `MyService` or `App` code.
//------------------------
1. The interface definitions (Logger, DataStore, Service)
2. The concrete implementations (SimpleLogger, InMemoryStore, MyService)
3. The use of context for cancellation and timeouts
4. The error handling approach
5. The functional options pattern used in the App struct
6. The overall structure and how it promotes loose coupling
//------------------------------------------------------------------------
*/
import (
    //4. Use of context: We use `context.Context` for cancellation and timeouts.
    "context"
    "errors"
    "fmt"
    "time"
)
//--------------------------------------
// Interfaces
//1. Interface-based design: We define interfaces for `Logger`, `DataStore`, and `Service`.
type Logger interface {
    Log(message string)
}

type DataStore interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key, value string) error
}

type Service interface {
    Process(ctx context.Context, input string) (string, error)
}
//---------------------------------------------
// Logger implementation
type SimpleLogger struct{}

func (l *SimpleLogger) Log(message string) {
    fmt.Println(message)
}

// DataStore implementation
type InMemoryStore struct {
    data map[string]string
}

func NewInMemoryStore() *InMemoryStore {
    return &InMemoryStore{data: make(map[string]string)}
}

func (s *InMemoryStore) Get(ctx context.Context, key string) (string, error) {
    if ctx.Err() != nil {
        return "", ctx.Err()
    }
    if value, ok := s.data[key]; ok {
        return value, nil
    }
    return "", errors.New("key not found")
}

func (s *InMemoryStore) Set(ctx context.Context, key, value string) error {
    if ctx.Err() != nil {
        return ctx.Err()
    }
    s.data[key] = value
    return nil
}

// Service implementation
//2. Dependency Injection: The `MyService` struct takes `Logger` and `DataStore` as dependencies.

type MyService struct {
    logger Logger
    store  DataStore
}

func NewMyService(logger Logger, store DataStore) Service {
    return &MyService{
        logger: logger,
        store:  store,
    }
}

func (s *MyService) Process(ctx context.Context, input string) (string, error) {
    s.logger.Log("Processing input: " + input)
    
    value, err := s.store.Get(ctx, input)
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return "", fmt.Errorf("operation timed out: %w", err)
        }
        //5. Error handling: Errors are returned and wrapped with additional context.
        return "", fmt.Errorf("failed to get value: %w", err)
    }
    
    result := "Processed: " + value
    return result, nil
}

// Functional options for configuring the application
type AppOption func(*App)

func WithTimeout(timeout time.Duration) AppOption {
    return func(a *App) {
        a.timeout = timeout
    }
}

// App struct to tie everything together#
//3. Composition: The `App` struct composes the `Service` interface.
type App struct {
    service Service
    timeout time.Duration
}

//6. Functional options: The `NewApp` function uses functional options for configuration.
func NewApp(service Service, options ...AppOption) *App {
    app := &App{
        service: service,
        timeout: 5 * time.Second, // default timeout
    }
    
    for _, option := range options {
        option(app)
    }
    
    return app
}

func (a *App) Run(input string) {
    ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
    defer cancel()
    
    result, err := a.service.Process(ctx, input)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Result: %s\n", result)
}

func main() {
    logger := &SimpleLogger{}
    store := NewInMemoryStore()
    service := NewMyService(logger, store)
    
    app := NewApp(service, WithTimeout(3*time.Second))
    
    // Simulate some data
    store.Set(context.Background(), "hello", "world")
    
    app.Run("hello")
    app.Run("nonexistent")
}