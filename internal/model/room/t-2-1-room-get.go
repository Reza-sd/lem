package room

// ========================hasOneFreeSeat======================================
var hasOneFreeSeat_Cases = TestCasesForFunc{
	FuncName: "Room.HasOneFreeSeat",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "return true",
			Got: SampleRoom.Middle_Name_3().hasOneFreeSeat(),
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return false",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.SetUsedSeats(myRoom.GetAllSeats())
				return myRoom.hasOneFreeSeat()
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
			Got: newPlainRoom().GetOneRandomNextRoom(),
			Exp: m(0),
		},
		//---------------------------------------
		{
			Des: "return err if connection slice is empty",
			Got: func() any {
				r := newPlainRoom()
				r.GetOneRandomNextRoom()
				return r.Errdb[GetOneRandomNextRoom] != Null
			},
			Exp: true,
		},
		//---------------------------------------
		{
			Des: "return nil if connection slice is not empty",
			Got: func() any {
				r := newPlainRoom()
				r.SetConnectionSlice([]m{1, 2, 3}).GetOneRandomNextRoom()
				return r.Errdb[GetOneRandomNextRoom] == Null
			},
			Exp: true,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "return random if connection slice is not empty",
			Got: newPlainRoom().SetConnectionSlice([]m{5}).GetOneRandomNextRoom(),
			Exp: m(5),
		},
		//---------------------------------------
	},
}

// ======================================================
var Print_Cases = TestCasesForFunc{
	FuncName: "Room.Print",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Print",
			Got: func() any {
				SampleRoom.Middle_Name_3().SetConnectionSlice([]m{1, 2, 3000}).Print()

				return true
			},
			Exp: true,
		},
		//---------------------------------------
	}}

//======================================================
