package room

// ======================================================
var Print_Cases = TestCasesFunc{
	FuncName: "Room.Print",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Print",
			Got: func() any {
				SampleRoom.Middle_Name_3().set.connectionSlice([]Mtr{1, 2, 3000}).room.Print()

				return true
			},
			Exp: true,
		},
		//---------------------------------------
	}}

//======================================================
