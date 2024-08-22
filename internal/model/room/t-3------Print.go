package room

var Print_Cases = TestCasesFunc{
	FuncName: "Room.Print",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Print",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.ConnectionSlice = []Mtr{1, 2, 3000}
				myRoom.Print()

				return true
			},
			Exp: true,
		},
	}}
