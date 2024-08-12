package antpk

//"main/internal/graphpk"
//"strings"

// ==================Samples=============================
var (
	//-------Ant type example-----------

	// //-------Ant type example-----------
	ant1 = Ant{
		Name:            1,
		CurrentRoomName: 0,
		VisitedRoomsArr: []mta{0},
		StepNumber:      0,
	}
	// //-------Ant type example-----------
	ant2 = Ant{
		Name:            2,
		CurrentRoomName: 0,
		VisitedRoomsArr: []mta{0},
		StepNumber:      0,
	}
	// //-------AntGroup type example------

	// //------------------
	// // travelPlan1 = TravelPlan{
	// // 	FinalSequence: 2,
	// // 	TheBestPlan:   antGroup3,
	// // }
	// //------------------

	//------------------
	Sample_AntGroup_1ant_initmode_room_0 = AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[mta]Ant{
			1: ant1,
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[mta][]mta),
		NotArrivedAntsName:    map[mta]struct{}{1: {}},
	}
	//----------------------
	Sample_AntGroup_2ant_initmode_room_0 = AntGroup{
		NumberOfAnts: 2,
		AntsDb: map[mta]Ant{
			1: ant1,
			2: ant2,
		},
		CurrentSequenceNumber: 0,
		UsedTunnel:            make(map[mta][]mta),
		NotArrivedAntsName:    map[mta]struct{}{1: {}, 2: {}},
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
