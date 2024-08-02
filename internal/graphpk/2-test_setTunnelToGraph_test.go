package graphpk

import (
	"reflect"
	"testing"
)

func TestSetTunnelAndSeatsToGraphRooms(t *testing.T) {
	tests := []struct {
		name          string
		graph         *Graph
		tunnelArr     []string
		expectedErr   bool
		expectedRooms map[string]Room
	}{
		{
			name: "Valid tunnels",
			graph: &Graph{
				Rooms:         make(map[string]Room),
				StartRoomName: "start",
				EndRoomName:   "end",
			},
			tunnelArr:   []string{"start-A", "A-B", "B-end"},
			expectedErr: false,
			expectedRooms: map[string]Room{
				"start": {Name: "start", Tunnels: []string{"A"}, MaxSeats: 100000, EmptySeats: 100000},
				"A":     {Name: "A", Tunnels: []string{"start", "B"}, MaxSeats: 1, EmptySeats: 1},
				"B":     {Name: "B", Tunnels: []string{"A", "end"}, MaxSeats: 1, EmptySeats: 1},
				"end":   {Name: "end", Tunnels: []string{"B"}, MaxSeats: 100000, EmptySeats: 100000},
			},
		},
		{
			name: "Invalid tunnel format",
			graph: &Graph{
				Rooms: make(map[string]Room),
			},
			tunnelArr:     []string{"invalid_format"},
			expectedErr:   true,
			expectedRooms: map[string]Room{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.graph.setTunnelAndSeatsToGraphRooms(tt.tunnelArr)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Check if the rooms are set correctly
				for roomName, expectedRoom := range tt.expectedRooms {
					actualRoom, exists := tt.graph.Rooms[roomName]
					if !exists {
						t.Errorf("Expected room %s to exist, but it doesn't", roomName)
						continue
					}

					if !reflect.DeepEqual(actualRoom, expectedRoom) {
						t.Errorf("Room %s doesn't match expected. Got %+v, want %+v", roomName, actualRoom, expectedRoom)
					}
				}

				// Check if there are no extra rooms
				if len(tt.graph.Rooms) != len(tt.expectedRooms) {
					t.Errorf("Number of rooms doesn't match. Got %d, want %d", len(tt.graph.Rooms), len(tt.expectedRooms))
				}
			}
		})
	}
}
