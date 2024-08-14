package antpk

//"main/internal/graphpk"
//"strings"

// ==================Samples=============================

// //-------Ant type example-----------
func Sample_ant1() Ant {

	return Ant{
		CurrentRoomName: 0,
		VisitedRoomsArr: []mta{0},
		StepNumber:      0,
	}

}

// //-------Ant type example-----------
func Sample_ant2() Ant {
	return Ant{
		//Name:            2,
		CurrentRoomName: 0,
		VisitedRoomsArr: []mta{0},
		StepNumber:      0,
	}
}

// //-------AntGroup type example------

// //------------------
// // travelPlan1 = TravelPlan{
// // 	FinalSequence: 2,
// // 	TheBestPlan:   antGroup3,
// // }
// //------------------

// ------------------
func Sample_AntGroup_1ant_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[mta]Ant{
			1: Sample_ant1(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[mta][]mta),
		NotArrivedAntsName:    map[mta]struct{}{1: {}},
	}
}

// ----------------------
func Sample_AntGroup_2ant_initmode_room_0() AntGroup {
	return AntGroup{
		NumberOfAnts: 2,
		AntsDb: map[mta]Ant{
			1: Sample_ant1(),
			2: Sample_ant2(),
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[mta][]mta),
		NotArrivedAntsName:    map[mta]struct{}{1: {}, 2: {}},
	}
}

//--------------------------------------
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

// ===============================================

// -----------------------------------------------------
// func assertIfStruct(t testing.TB, s interface{}, expectation bool) {
// 	t.Helper()

// 	v := reflect.ValueOf(s)
// 	if v.Kind() != reflect.Struct && expectation {
// 		t.Errorf("not struct")
// 	}

// }
