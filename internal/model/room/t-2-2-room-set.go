package room

// ====================================================================
var builder_Cases = TestCasesFunc{
	FuncName: "Room.set.builder",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "if name=0 means start home",
			Got: roomBuilder(0, 5, []Mtr{1, 2, 3}),
			Exp: newRoom().set.
				name(0).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd).
				connectionSlice([]Mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "non start or end room",
			Got: roomBuilder(1, 5, []Mtr{1, 2, 3}),

			Exp: newRoom().set.
				name(1).
				allSeats(1).
				usedSeats(0).
				connectionSlice([]Mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "end room",
			Got: roomBuilder(5, 5, []Mtr{1, 2, 3}),

			Exp: newRoom().set.
				name(5).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd).
				connectionSlice([]Mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
	}}

//====================================================================
