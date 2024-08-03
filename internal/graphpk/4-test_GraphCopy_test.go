package graphpk

import (
	"reflect"
	"testing"
)

func Test_GraphCopyFresh(t *testing.T) {
	// Create a base graph for testing
	baseGraph := Graph{
		StartRoomName:        "start",
		EndRoomName:          "end",
		CurrentAntsInEndRoom: 0,
		NumberOfAllRoom:      4,
		Rooms: map[string]Room{
			"start": {Name: "start", Tunnels: []string{"A"}, MaxSeats: 100000, EmptySeats: 100000},
			"A":     {Name: "A", Tunnels: []string{"start", "B"}, MaxSeats: 1, EmptySeats: 1},
			"B":     {Name: "B", Tunnels: []string{"A", "end"}, MaxSeats: 1, EmptySeats: 1},
			"end":   {Name: "end", Tunnels: []string{"B"}, MaxSeats: 100000, EmptySeats: 100000},
		},
	}

	// Create an empty second graph
	secondGraph := Graph{}

	// Call the function to copy
	GraphCopyFresh(baseGraph, &secondGraph)

	// Test cases
	t.Run("Copy basic properties", func(t *testing.T) {
		if secondGraph.StartRoomName != baseGraph.StartRoomName {
			t.Errorf("StartRoomName mismatch. Got %s, want %s", secondGraph.StartRoomName, baseGraph.StartRoomName)
		}
		if secondGraph.EndRoomName != baseGraph.EndRoomName {
			t.Errorf("EndRoomName mismatch. Got %s, want %s", secondGraph.EndRoomName, baseGraph.EndRoomName)
		}
		if secondGraph.CurrentAntsInEndRoom != 0 {
			t.Errorf("CurrentAntsInEndRoom should be 0, got %d", secondGraph.CurrentAntsInEndRoom)
		}
		if secondGraph.NumberOfAllRoom != baseGraph.NumberOfAllRoom {
			t.Errorf("NumberOfAllRoom mismatch. Got %d, want %d", secondGraph.NumberOfAllRoom, baseGraph.NumberOfAllRoom)
		}
	})

	t.Run("Copy rooms", func(t *testing.T) {
		//t.Skip()
		if len(secondGraph.Rooms) != len(baseGraph.Rooms) {
			t.Errorf("Number of rooms mismatch. Got %d, want %d", len(secondGraph.Rooms), len(baseGraph.Rooms))
		}

		for roomName, value := range baseGraph.Rooms {
			copiedRoom, exists := secondGraph.Rooms[roomName]
			if !exists {
				t.Errorf("Room %s not found in copied graph", roomName)
				continue
			}

			if copiedRoom.Name != value.Name {
				t.Errorf("Room %s: Name mismatch. Got %s, want %s", roomName, copiedRoom.Name, value.Name)
			}
			if copiedRoom.MaxSeats != value.MaxSeats {
				t.Errorf("Room %s: MaxSeats mismatch. Got %d, want %d", roomName, copiedRoom.MaxSeats, value.MaxSeats)
			}
			if copiedRoom.EmptySeats != value.EmptySeats {
				t.Errorf("Room %s: EmptySeats mismatch. Got %d, want %d", roomName, copiedRoom.EmptySeats, value.EmptySeats)
			}

			// Check if Tunnels are copied correctly
			if !reflect.DeepEqual(copiedRoom.Tunnels, value.Tunnels) {
				t.Errorf("Room %s: Tunnels mismatch. Got %v, want %v", roomName, copiedRoom.Tunnels, value.Tunnels)
			}

			// Ensure that the Tunnels slice is a separate copy
			if len(value.Tunnels) > 0 && len(copiedRoom.Tunnels) > 0 {
				if &value.Tunnels[0] == &copiedRoom.Tunnels[0] {
					t.Errorf("Room %s: Tunnels slice was not copied, it's the same slice", roomName)
				}
			}
		}
	})
}
