package room

var HasOneFreeSeat_Cases = TestCasesFunc{
	FuncName: "Room.HasOneFreeSeat",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				return myRoom.HasOneFreeSeat()

			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return false",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.UsedSeats = myRoom.AllSeats

				return myRoom.HasOneFreeSeat()
			},
			Exp: false,
		},
		//---------------------------------------

	}}
