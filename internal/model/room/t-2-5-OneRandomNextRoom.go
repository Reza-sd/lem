package room

var OneRandomNextRoom_Cases = TestCasesFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return 0 if connection slice is empty",
			Got: func() any {
				myRoom := Room{}
				//myRoom.Init(0, 5, []Mtr{1, 2, 3})
				next, _ := myRoom.OneRandomNextRoom()
				return Mtr(next)
			},
			Exp: Mtr(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: func() any {
				myRoom := Room{}
				_, err := myRoom.OneRandomNextRoom()
				return err != nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: func() any {
				myRoom := Room{
					ConnectionSlice: []Mtr{1, 2, 3},
				}
				_, err := myRoom.OneRandomNextRoom()
				return err == nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return random if connection slice is not empty",
			Got: func() any {
				myRoom := Room{
					ConnectionSlice: []Mtr{5},
				}
				next, _ := myRoom.OneRandomNextRoom()
				return next
			},
			Exp: Mtr(5),
		},
		//---------------------------------------
	},
}
