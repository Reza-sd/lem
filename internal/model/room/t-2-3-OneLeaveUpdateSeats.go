package room

var OneLeaveUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneLeaveUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Do nothing if usedSeats=0",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.OneLeaveUpdateSeats()
				return myRoom
			},
			Exp: Room{
				Name:            3,
				AllSeats:        1,
				UsedSeats:       0,
				ConnectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: func() any {
				myRoom := SampleRoom.Start_Name_0()
				myRoom.OneLeaveUpdateSeats()
				return myRoom
			},
			Exp: Room{
				Name:            0,
				AllSeats:        MaxSeatsStartEnd,
				UsedSeats:       UsedSeatsStartEnd - 1,
				ConnectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
	}}
