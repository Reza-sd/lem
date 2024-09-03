package room

var oneLeave_Cases = TestCasesFunc{
	FuncName: "Room.oneLeave",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Do nothing if usedSeats=0",
			Got: func() any {
				r, _ := SampleRoom.Middle_Name_3().update.oneLeave()
				return r
			},

			Exp: newRoom().set.
				name(3).
				allSeats(1).
				usedSeats(0).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: func() any {
				r, _ := SampleRoom.Start_Name_0().update.oneLeave()
				return r
			},

			Exp: newRoom().set.
				name(0).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd - 1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
	}}
