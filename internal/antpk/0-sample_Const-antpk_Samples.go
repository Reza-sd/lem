package antpk

import (
//"main/internal/graphpk"
//"strings"
)

// ==================Samples=============================
var (
	//-------Ant type example-----------

	// //-------Ant type example-----------
	ant1 = Ant{
		Name:            1,
		CurrentRoomName: 0,
		VisitedRoomsArr: []t{0},
		StepNumber:      0,
	}
	// //-------Ant type example-----------
	ant2 = Ant{
		Name:            2,
		CurrentRoomName: 0,
		VisitedRoomsArr: []t{0},
		StepNumber:      0,
	}
	// //-------AntGroup type example------

	// antGroup3 = AntGroup{
	// 	NumberOfAnts: 2,
	// 	AntsDb: map[string]Ant{
	// 		"L1": ant1,
	// 		"L2": ant2,
	// 	},
	// 	CurrentSequenceNumber: 2,
	// 	NotArrivedAntsName: map[string]struct{}{
	// 		"L1": struct{}{},
	// 		"L2": struct{}{},
	// 	},
	// }
	// //------------------
	// // travelPlan1 = TravelPlan{
	// // 	FinalSequence: 2,
	// // 	TheBestPlan:   antGroup3,
	// // }
	// //------------------
	// sample_AntGroup_2ants_initmode_room_0 = AntGroup{
	// 	NumberOfAnts: 2,
	// 	AntsDb: map[string]Ant{
	// 		"L1": {Name: "L1",
	// 			CurrentRoomName: "room_0",
	// 			VisitedRoomsArr: []string{"room_0"},
	// 			StepNumber:      0,
	// 		},
	// 		"L2": {Name: "L2",
	// 			CurrentRoomName: "room_0",
	// 			VisitedRoomsArr: []string{"room_0"},
	// 			StepNumber:      0,
	// 		},
	// 	},
	// 	CurrentSequenceNumber: 0,
	// 	UsedTunnel:            make(map[string][]string),
	// }
	// //------------------
	Sample_AntGroup_1ant_initmode_room_0 = AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[t]Ant{
			1: ant1,
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[t][]t),
		NotArrivedAntsName:    make(map[t]struct{}),
	}
	// //------------------
	// sample_room_0_startRoom = graphpk.Room{
	// 	Name:        "room_0",
	// 	MaxSeats:    1000,
	// 	EmptySeats:  1000,
	// 	Connections: []string{"room_1"},
	// }
	// //------------------------
	// sample_room_1_EndRoom = graphpk.Room{
	// 	Name:        "room_1",
	// 	MaxSeats:    1000,
	// 	EmptySeats:  1000,
	// 	Connections: []string{"room_0"},
	// }
)

// ===============================================

// -----------------------------------------------------
// func assertIfStruct(t testing.TB, s interface{}, expectation bool) {
// 	t.Helper()

// 	v := reflect.ValueOf(s)
// 	if v.Kind() != reflect.Struct && expectation {
// 		t.Errorf("not struct")
// 	}

// }
