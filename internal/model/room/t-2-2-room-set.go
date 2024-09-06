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

			Exp: newPlainRoom().
				SetName(3).
				SetAllSeats(1).
				SetUsedSeats(1).
				SetConnectionSlice([]mtr{}),
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: SampleRoom.End_Name_1().UpdateOneCome(),

			Exp: newPlainRoom().
				SetName(1).
				SetAllSeats(MaxSeatsStartEnd).
				SetUsedSeats(UsedSeatsStartEnd + 1).
				SetConnectionSlice([]mtr{}),
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

			Exp: newPlainRoom().
				SetName(3).
				SetAllSeats(1).
				SetUsedSeats(0).
				SetConnectionSlice([]mtr{}),
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: SampleRoom.Start_Name_0().UpdateOneLeave(),

			Exp: newPlainRoom().
				SetName(0).
				SetAllSeats(MaxSeatsStartEnd).
				SetUsedSeats(UsedSeatsStartEnd - 1).
				SetConnectionSlice([]mtr{}),
		},
		//---------------------------------------
	}}

//==============================================================
