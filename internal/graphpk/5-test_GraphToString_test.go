package graphpk

import (
	"fmt"

	"testing"
)

func TestGraphToString(t *testing.T) {
	// Create a sample graph for testing
	graph := &Graph{
		StartRoomName:        "start",
		EndRoomName:          "end",
		NumberOfAllRoom:      4,
		CurrentAntsInEndRoom: 0,
		Rooms: map[string]Room{
			"start": {Name: "start", Tunnels: []string{"A"}, MaxSeats: 100000, EmptySeats: 99995},
			"A":     {Name: "A", Tunnels: []string{"start", "B"}, MaxSeats: 1, EmptySeats: 0},
			"B":     {Name: "B", Tunnels: []string{"A", "end"}, MaxSeats: 1, EmptySeats: 1},
			"end":   {Name: "end", Tunnels: []string{"B"}, MaxSeats: 100000, EmptySeats: 99995},
		},
	}

	// Call the ToString method
	result, err := graph.ToString()

	// Check for errors
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	fmt.Println(result)
}
