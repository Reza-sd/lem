package room

// ===================================================
var oneCome_Cases = TestCasesFunc{
	FuncName: "Room.update.oneCome",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: SampleRoom.Middle_Name_3().update.oneCome().ans,

			Exp: newPlainRoom().set.
				name(3).
				allSeats(1).
				usedSeats(1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: SampleRoom.End_Name_1().update.oneCome().ans,

			Exp: newPlainRoom().set.
				name(1).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd + 1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
	}}

// ==============================================================
var oneLeave_Cases = TestCasesFunc{
	FuncName: "Room.update.oneLeave",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Do nothing if usedSeats=0",
			Got: SampleRoom.Middle_Name_3().update.oneLeave().ans,

			Exp: newPlainRoom().set.
				name(3).
				allSeats(1).
				usedSeats(0).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: SampleRoom.Start_Name_0().update.oneLeave().ans,

			Exp: newPlainRoom().set.
				name(0).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd - 1).
				connectionSlice([]Mtr{}).
				room,
		},
		//---------------------------------------
	}}

//==============================================================
