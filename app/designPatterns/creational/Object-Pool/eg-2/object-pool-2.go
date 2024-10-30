package main

import (
    "fmt"
    "sync"
)

// Connection is a placeholder struct representing a database connection
type Connection struct {
    ID int
}

// Connect simulates an action that opens a database connection
func (conn *Connection) Connect() {
    fmt.Printf("Connection %d opened.\n", conn.ID)
}

// Close simulates closing the connection
func (conn *Connection) Close() {
    fmt.Printf("Connection %d closed.\n", conn.ID)
}

// ConnectionPool struct manages a pool of reusable Connection objects
type ConnectionPool struct {
    pool *sync.Pool
}

// NewConnectionPool initializes the connection pool with a specified number of connections
func NewConnectionPool(size int) *ConnectionPool {
    var counter int
    return &ConnectionPool{
        pool: &sync.Pool{
            New: func() interface{} {
                counter++
                conn := &Connection{ID: counter}
                conn.Connect()
                return conn
            },
        },
    }
}

// Get retrieves a connection from the pool
func (cp *ConnectionPool) Get() *Connection {
    return cp.pool.Get().(*Connection)
}

// Put returns a connection back to the pool
func (cp *ConnectionPool) Put(conn *Connection) {
    conn.Close()
    cp.pool.Put(conn)
}

func main() {
    // Initialize a connection pool
    connPool := NewConnectionPool(2)

    // Get a connection from the pool
    conn1 := connPool.Get()
    fmt.Printf("Using connection %d\n", conn1.ID)

    // Return the connection back to the pool
    connPool.Put(conn1)

    // Get another connection from the pool
    conn2 := connPool.Get()
    fmt.Printf("Using connection %d\n", conn2.ID)

    // Put the connection back to the pool
    connPool.Put(conn2)
}
