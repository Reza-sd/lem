package main

import (
        "context"
        "fmt"
        "time"
)

func longRunningTask(ctx context.Context) {
        for {
                select {
                case <-ctx.Done():
                        fmt.Println("Task canceled")
                        return
                case <-time.After(time.Second):
                        fmt.Println("Doing some work...")
                }
        }
}

func main() {
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        go longRunningTask(ctx)

        time.Sleep(10 * time.Second)
        fmt.Println("Main function exiting")
}