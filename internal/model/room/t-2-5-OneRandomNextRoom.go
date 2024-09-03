package room

var OneRandomNextRoom_Cases = TestCasesFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return 0 if connection slice is empty",
			Got: func() any {
				next, _ := newRoom().get.OneRandomNextRoom()
				return Mtr(next)
			},
			Exp: Mtr(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: func() any {
				_, err := newRoom().get.OneRandomNextRoom()
				return err != nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: func() any {
				_, err := newRoom().set.connectionSlice([]Mtr{1, 2, 3}).room.get.OneRandomNextRoom()
				return err == nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "return random if connection slice is not empty",
			Got: func() any {
				next, _ := newRoom().set.connectionSlice([]Mtr{5}).room.get.OneRandomNextRoom()
				return next
			},
			Exp: Mtr(5),
		},
		//---------------------------------------
	},
}
