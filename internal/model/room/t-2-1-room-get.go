package room

// ========================hasOneFreeSeat======================================
var hasOneFreeSeat_Cases = TestCasesForFunc{
	FuncName: "Room.HasOneFreeSeat",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: SampleRoom.Middle_Name_3().get.hasOneFreeSeat(),
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return false",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.set.usedSeats(myRoom.get.allSeats())
				return myRoom.get.hasOneFreeSeat()
			},
			Exp: false,
		},
		//---------------------------------------

	}}

// ========================OneRandomNextRoom====================
var test_oneRandomNextRoom_Cases = TestCasesForFunc{
	FuncName: "Room.OneRandomNextRoom",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return 0 if connection slice is empty",
			Got: newPlainRoom().get.OneRandomNextRoom().ans,
			Exp: Mtr(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: newPlainRoom().get.OneRandomNextRoom().sCode != null,
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: newPlainRoom().set.connectionSlice([]Mtr{1, 2, 3}).room.get.OneRandomNextRoom().sCode == null,
			Exp: true,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "return random if connection slice is not empty",
			Got: newPlainRoom().set.connectionSlice([]Mtr{5}).room.get.OneRandomNextRoom().ans,
			Exp: Mtr(5),
		},
		//---------------------------------------
	},
}

//==============================================================
