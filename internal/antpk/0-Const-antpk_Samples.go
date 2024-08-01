package antpk

// ==================Samples=============================
var (
	//-------Ant type example-----------
	ant1 = Ant{
		Name:            "L1",
		CurrentRoomName: "room_1",
		VisitedRoomsArr: []string{"room_0", "room_1"},
		StepNumber:      1,
	}
	//-------Ant type example-----------
	ant2 = Ant{
		Name:            "L2",
		CurrentRoomName: "mio_5",
		VisitedRoomsArr: []string{"mio_7", "mio_3", "mio_5"},
		StepNumber:      2,
	}
	//-------AntGroup type example------

	antGroup3 = AntGroup{
		NumberOfAnts: 2,
		AntsMap: map[int]Ant{
			1: ant1,
			2: ant2,
		},
		NumberOfSequence: 2,
	}
	//------------------
	travelPlan1 = TravelPlan{
		FinalSequence: 2,
		TheBestPlan:   antGroup3,
	}
	//------------------
	sample_AntGroup_2ants_initmode_room_0 = AntGroup{
		NumberOfAnts: 2,
		AntsMap: map[int]Ant{
			1: {Name: "L1",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
			2: {Name: "L2",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
		},
		NumberOfSequence: 0,
	}
	//------------------
	sample_AntGroup_1ant_initmode_room_0 = AntGroup{
		NumberOfAnts: 1,
		AntsMap: map[int]Ant{
			1: {Name: "L1",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
		},
		NumberOfSequence: 0,
	}
	//------------------
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
