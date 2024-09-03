package room

var oneCome_Cases = TestCasesFunc{
	FuncName: "Room.oneCome",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: func() any {
				r, _ := SampleRoom.Middle_Name_3().update.oneCome()
				return r

			},

			Exp: newRoom().set.
				name(3).
				allSeats(1).
				usedSeats(1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: func() any {
				r, _ := SampleRoom.End_Name_1().update.oneCome()
				return r
			},

			Exp: newRoom().set.
				name(1).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd + 1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
	}}
