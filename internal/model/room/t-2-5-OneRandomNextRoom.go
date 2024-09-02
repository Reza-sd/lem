package room

var OneRandomNextRoom_Cases = TestCasesFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return error if connection slice is empty",
			Got: func() any {
				myRoom := Room{}
				//myRoom.Init(0, 5, []Mtr{1, 2, 3})
				next, _ := myRoom.OneRandomNextRoom()
				return Mtr(next)
			},
			Exp: Mtr(0),
		},
		//---------------------------------------
	},
}
