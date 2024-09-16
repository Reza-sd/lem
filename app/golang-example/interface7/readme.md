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
//---------------------------------------------------
1. The interface definitions (Logger, DataStore, Service)
2. The concrete implementations (SimpleLogger, InMemoryStore, MyService)
3. The use of context for cancellation and timeouts
4. The error handling approach
5. The functional options pattern used in the App struct
6. The overall structure and how it promotes loose coupling
//--------------------------------------------------------