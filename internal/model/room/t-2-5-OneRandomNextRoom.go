package room

var OneRandomNextRoom_Cases = TestCasesFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return 0 if connection slice is empty",
			Got: func() any {
				myRoom := newRoom()
				next, _ := myRoom.OneRandomNextRoom()
				return Mtr(next)
			},
			Exp: Mtr(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: func() any {
				myRoom := newRoom()
				_, err := myRoom.OneRandomNextRoom()
				return err != nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: func() any {
				myRoom := newRoom().set.connectionSlice([]Mtr{1, 2, 3}).room
				//myRoom.set.connectionSlice([]Mtr{1, 2, 3})
				_, err := myRoom.OneRandomNextRoom()
				return err == nil
			},
			Exp: true,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "return random if connection slice is not empty",
			Got: func() any {
				myRoom := newRoom().set.connectionSlice([]Mtr{5}).room
				//myRoom.set.connectionSlice([]Mtr{5})
				next, _ := myRoom.OneRandomNextRoom()
				return next
			},
			Exp: Mtr(5),
		},
		//---------------------------------------
	},
}
