package room

// ===================================================
var oneCome_Cases = TestCasesForFunc{
	FuncName: "Room.set.oneCome",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: func() any {
				r, _ := SampleRoom.Middle_Name_3().UpdateOneCome()
				return r
			},

			Exp: &room{
				name:            3,
				allSeats:        1,
				usedSeats:       1,
				connectionSlice: []m{},
			},
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: func() any {
				r, _ := SampleRoom.End_Name_1().UpdateOneCome()
				return r
			},
			Exp: &room{
				name:            1,
				allSeats:        MaxSeatsStartEnd,
				usedSeats:       UsedSeatsStartEnd + 1,
				connectionSlice: []m{},
			},
		},
		//---------------------------------------
	}}

// ==============================================================
var oneLeave_Cases = TestCasesForFunc{
	FuncName: "Room.set.oneLeave",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Do nothing if usedSeats=0",
			Got: func() any {
				r, _ := SampleRoom.Middle_Name_3().UpdateOneLeave()
				return r
			},
			Exp: &room{
				name:            3,
				allSeats:        1,
				usedSeats:       0,
				connectionSlice: []m{},
			},
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: func() any {
				r, _ := SampleRoom.Start_Name_0().UpdateOneLeave()
				return r
			},
			Exp: &room{
				name:            0,
				allSeats:        MaxSeatsStartEnd,
				usedSeats:       UsedSeatsStartEnd - 1,
				connectionSlice: []m{},
			},
		},
		//---------------------------------------
	}}

//==============================================================
