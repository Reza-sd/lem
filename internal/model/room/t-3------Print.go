package room

var Print_Cases = TestCasesFunc{
	FuncName: "Room.Print",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Print",
			Got: func() any {
				r := SampleRoom.Middle_Name_3()
				r.connectionSlice = []Mtr{1, 2, 3000}
				r.Print()
				return true
			},
			Exp: true,
		},
	}}
