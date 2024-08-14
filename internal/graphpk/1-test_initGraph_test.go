package graphpk

import (
	"reflect"
	"testing"
)

func TestInitGraph(t *testing.T) {
	tests := []struct {
		name           string
		tunnelMap      map[mtg][]mtg
		endRoomName    mtg
		expectedGraph  Graph
		expectedError  error
	}{
		{
			name: "Basic initialization",
			tunnelMap: map[mtg][]mtg{
				1: {2, 3},
				2: {1, 4},
				3: {1, 4},
				4: {2, 3},
			},
			endRoomName: 4,
			expectedGraph: Graph{
				EndRoomName: 4,
				RoomHasEmptySeatDb: map[mtg]bool{
					1: true,
					2: true,
					3: true,
					4: true,
				},
				TunnelsDb: map[mtg][]mtg{
					1: {2, 3},
					2: {1, 4},
					3: {1, 4},
					4: {2, 3},
				},
			},
			expectedError: nil,
		},
		// {
		// 	name:           "Empty tunnel map",
		// 	tunnelMap:      map[mtg][]mtg{},
		// 	endRoomName:    1,
		// 	expectedGraph:  Graph{
		// 		EndRoomName:        1,
		// 		RoomHasEmptySeatDb: map[mtg]bool{},
		// 		TunnelsDb:          map[mtg][]mtg{},
		// 	},
		// 	expectedError: nil,
		// },
	}
//------------------------------------------
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myGraph := &Graph{}
			err := myGraph.InitGraph(tt.tunnelMap, tt.endRoomName)
			myGraph.Print()
			if err != tt.expectedError {
				t.Errorf("InitGraph() error = %v, expectedError %v", err, tt.expectedError)
				return
			}

			if myGraph.EndRoomName != tt.expectedGraph.EndRoomName {
				t.Errorf("EndRoomName = %v, want %v", myGraph.EndRoomName, tt.expectedGraph.EndRoomName)
			}

			if !reflect.DeepEqual(myGraph.RoomHasEmptySeatDb, tt.expectedGraph.RoomHasEmptySeatDb) {
				t.Errorf("RoomHasEmptySeatDb = %v, want %v", myGraph.RoomHasEmptySeatDb, tt.expectedGraph.RoomHasEmptySeatDb)
			}

			if !reflect.DeepEqual(myGraph.TunnelsDb, tt.expectedGraph.TunnelsDb) {
				t.Errorf("TunnelsDb = %v, want %v", myGraph.TunnelsDb, tt.expectedGraph.TunnelsDb)
			}
		})
	}
	//-------------------------------------
}