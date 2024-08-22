package room

var HasOneFreeSeat_Cases = TestCasesFunc{
	FuncName: "Room.HasOneFreeSeat",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: Sample_Room_Middle.HasOneFreeSeat(),
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return false",
			Got: func() any {
				Sample_Room_Middle.UsedSeats = Sample_Room_Middle.AllSeats
				return Sample_Room_Middle.HasOneFreeSeat()
			},
			Exp: false,
		},
		//---------------------------------------

	}}
