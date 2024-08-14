package graphpk

import (
	"reflect"
	"testing"
)

func TestInitGraph(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		tunnelMap     TunnelMap
		endRoomName   Mtg
		expectedGraph Graph
	}{
		{
			name: "Basic graph initialization",
			tunnelMap: TunnelMap{
				1: {2, 3},
				2: {1, 4},
				3: {1, 4},
				4: {2, 3},
			},
			endRoomName: 4,
			expectedGraph: Graph{
				EndRoomName: 4,
				RoomAvailableDb: map[Mtg]bool{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				TunnelsDb: &TunnelMap{
					1: {2, 3},
					2: {1, 4},
					3: {1, 4},
					4: {2, 3},
				},
			},
		},
		{
			name:        "Empty graph initialization",
			tunnelMap:   TunnelMap{},
			endRoomName: 0,
			expectedGraph: Graph{
				EndRoomName:     0,
				RoomAvailableDb: map[Mtg]bool{},
				TunnelsDb:       &TunnelMap{},
			},
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			myGraph := &Graph{}
			err := myGraph.InitGraph(tc.tunnelMap, tc.endRoomName)

			myGraph.Print()
			// Check for errors
			if err != nil {
				t.Fatalf("InitGraph returned an error: %v", err)
			}

			// Check EndRoomName
			if myGraph.EndRoomName != tc.expectedGraph.EndRoomName {
				t.Errorf("EndRoomName mismatch. Got %d, want %d", myGraph.EndRoomName, tc.expectedGraph.EndRoomName)
			}

			// Check RoomAvailableDb
			if !reflect.DeepEqual(myGraph.RoomAvailableDb, tc.expectedGraph.RoomAvailableDb) {
				t.Errorf("RoomAvailableDb mismatch. Got %v, want %v", myGraph.RoomAvailableDb, tc.expectedGraph.RoomAvailableDb)
			}

			// Check TunnelsDb
			if !reflect.DeepEqual(myGraph.TunnelsDb, tc.expectedGraph.TunnelsDb) {
				t.Errorf("TunnelsDb mismatch. Got %v, want %v", myGraph.TunnelsDb, tc.expectedGraph.TunnelsDb)
			}
		})
	}
}
