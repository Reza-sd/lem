package room

var HasOneFreeSeat_Cases = TestCasesFunc{
	FuncName: "Room.HasOneFreeSeat",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: SampleRoom.Middle_Name_3().HasOneFreeSeat(),
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return false",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.set.usedSeats(myRoom.get.allSeats())
				return myRoom.HasOneFreeSeat()
			},
			Exp: false,
		},
		//---------------------------------------

	}}
