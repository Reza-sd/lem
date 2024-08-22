package room

var OneComeUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneComeUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: func() any {
				Sample_Room_Middle.OneComeUpdateSeats()
				expRoom := *Sample_Room_Middle
				expRoom.UsedSeats = 10
				expRoom.Print()
				Sample_Room_Middle.Print()
				print("######")
				return &expRoom
			},
			Exp: &Sample_Room_Middle,
		},
		//---------------------------------------

		//---------------------------------------

	}}
