package graphpk

import (
	"reflect"
	"testing"
)

func TestInitGraph(t *testing.T) {
	tests := []struct {
		name          string
		tunnelArr     []string
		start         string
		end           string
		expectedErr   bool
		expectedGraph *Graph
	}{
		{
			name:        "Valid initialization",
			tunnelArr:   []string{"start-A", "A-B", "B-end"},
			start:       "start",
			end:         "end",
			expectedErr: false,
			expectedGraph: &Graph{
				StartRoomName:        "start",
				EndRoomName:          "end",
				CurrentAntsInEndRoom: 0,
				NumberOfAllRoom:      4,
				Rooms: map[string]Room{
					"start": {Name: "start", Connections: []string{"A"}, MaxSeats: 100000, EmptySeats: 100000},
					"A":     {Name: "A", Connections: []string{"start", "B"}, MaxSeats: 1, EmptySeats: 1},
					"B":     {Name: "B", Connections: []string{"A", "end"}, MaxSeats: 1, EmptySeats: 1},
					"end":   {Name: "end", Connections: []string{"B"}, MaxSeats: 100000, EmptySeats: 100000},
				},
			},
		},
		{
			name:        "Invalid tunnel format",
			tunnelArr:   []string{"invalid_format"},
			start:       "start",
			end:         "end",
			expectedErr: true,
			expectedGraph: &Graph{
				StartRoomName:        "start",
				EndRoomName:          "end",
				CurrentAntsInEndRoom: 0,
				Rooms:                map[string]Room{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myGraph := &Graph{}
			err := myGraph.InitGraph(tt.tunnelArr, tt.start, tt.end)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Check if the graph properties are set correctly
				if myGraph.StartRoomName != tt.expectedGraph.StartRoomName {
					t.Errorf("StartRoomName mismatch. Got %s, want %s", myGraph.StartRoomName, tt.expectedGraph.StartRoomName)
				}
				if myGraph.EndRoomName != tt.expectedGraph.EndRoomName {
					t.Errorf("EndRoomName mismatch. Got %s, want %s", myGraph.EndRoomName, tt.expectedGraph.EndRoomName)
				}
				if myGraph.CurrentAntsInEndRoom != tt.expectedGraph.CurrentAntsInEndRoom {
					t.Errorf("CurrentAntsInEndRoom mismatch. Got %d, want %d", myGraph.CurrentAntsInEndRoom, tt.expectedGraph.CurrentAntsInEndRoom)
				}
				if myGraph.NumberOfAllRoom != tt.expectedGraph.NumberOfAllRoom {
					t.Errorf("NumberOfAllRoom mismatch. Got %d, want %d", myGraph.NumberOfAllRoom, tt.expectedGraph.NumberOfAllRoom)
				}

				// Check if the rooms are set correctly
				if !reflect.DeepEqual(myGraph.Rooms, tt.expectedGraph.Rooms) {
					t.Errorf("Rooms mismatch. Got %+v, want %+v", myGraph.Rooms, tt.expectedGraph.Rooms)
				}
			}
		})
	}
}
