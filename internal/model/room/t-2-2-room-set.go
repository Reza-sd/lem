package room

// ===================================================
var oneCome_Cases = TestCasesForFunc{
	FuncName: "Room.set.oneCome",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: SampleRoom.Middle_Name_3().UpdateOneCome(),

			Exp: &room{
				name:            3,
				allSeats:        1,
				usedSeats:       1,
				connectionSlice: []m{},
				Errdb:           map[e]e{Null: Null},
			},
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: SampleRoom.End_Name_1().UpdateOneCome(),

			Exp: &room{
				name:            1,
				allSeats:        MaxSeatsStartEnd,
				usedSeats:       UsedSeatsStartEnd + 1,
				connectionSlice: []m{},
				Errdb:           map[e]e{Null: Null},
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
			Got: SampleRoom.Middle_Name_3().UpdateOneLeave(),

			Exp: &room{
				name:            3,
				allSeats:        1,
				usedSeats:       0,
				connectionSlice: []m{},
				Errdb: map[e]e{
					Null:           Null,
					UpdateOneLeave: UpdateOneLeave10,
				},
			},
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: SampleRoom.Start_Name_0().UpdateOneLeave(),

			Exp: &room{
				name:            0,
				allSeats:        MaxSeatsStartEnd,
				usedSeats:       UsedSeatsStartEnd - 1,
				connectionSlice: []m{},
				Errdb:           map[e]e{Null: Null},
			},
		},
		//---------------------------------------
	}}

//==============================================================
