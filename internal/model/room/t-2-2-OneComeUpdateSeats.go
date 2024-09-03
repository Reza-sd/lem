package room

var OneComeUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneComeUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.OneComeUpdateSeats()
				return myRoom
			},
			Exp: Room{
				name:            3,
				allSeats:        1,
				usedSeats:       1,
				connectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: func() any {
				myRoom := SampleRoom.End_Name_1()
				myRoom.OneComeUpdateSeats()
				return myRoom
			},
			Exp: Room{
				name:            1,
				allSeats:        MaxSeatsStartEnd,
				usedSeats:       UsedSeatsStartEnd + 1,
				connectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
	}}
