package antpk

// ==================Samples=============================
var (
	//-------Ant type example-----------

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
		AntsDb: map[string]Ant{
			"L1": ant1,
			"L2": ant2,
		},
		SequenceNumber: 2,
	}
	//------------------
	travelPlan1 = TravelPlan{
		FinalSequence: 2,
		TheBestPlan:   antGroup3,
	}
	//------------------
	sample_AntGroup_2ants_initmode_room_0 = AntGroup{
		NumberOfAnts: 2,
		AntsDb: map[string]Ant{
			"L1": {Name: "L1",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
			"L2": {Name: "L2",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
		},
		SequenceNumber: 0,
		UsedTunnel:     map[int]map[string]string{},
	}
	//------------------
	Sample_AntGroup_1ant_initmode_room_0 = AntGroup{
		NumberOfAnts: 1,
		AntsDb: map[string]Ant{
			"L1": {Name: "L1",
				CurrentRoomName: "room_0",
				VisitedRoomsArr: []string{"room_0"},
				StepNumber:      0,
			},
		},
		SequenceNumber: 0,
		UsedTunnel:     make(map[int]map[string]string),
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
