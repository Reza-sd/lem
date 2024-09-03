package room

var OneRandomNextRoom_Cases = TestCasesFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return 0 if connection slice is empty",
			//Got: newRoom().get.OneRandomNextRoom().ans.(Mtr),
			Got: newRoom().get.OneRandomNextRoom().ans,
			Exp: Mtr(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: newRoom().get.OneRandomNextRoom().err != nil,
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: newRoom().set.connectionSlice([]Mtr{1, 2, 3}).room.get.OneRandomNextRoom().err == nil,
			Exp: true,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "return random if connection slice is not empty",
			Got: newRoom().set.connectionSlice([]Mtr{5}).room.get.OneRandomNextRoom().ans,
			Exp: Mtr(5),
		},
		//---------------------------------------
	},
}
