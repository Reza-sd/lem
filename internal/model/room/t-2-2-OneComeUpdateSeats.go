package room

var OneComeUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneComeUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase 1 UsedSeats when Room in middle type",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.OneComeUpdateSeats()
				return myRoom
			},
			Exp: Room{
				Name:            3,
				AllSeats:        1,
				UsedSeats:       1,
				ConnectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
		{
			Des: "increase 1 UsedSeats when Room in End type",
			Got: func() any {
				myRoom := SampleRoom.End_Name_1()
				myRoom.OneComeUpdateSeats()
				return myRoom
			},
			Exp: Room{
				Name:            1,
				AllSeats:        5000,
				UsedSeats:       1001,
				ConnectionSlice: []Mtr{},
			},
		},
		//---------------------------------------
		//---------------------------------------

	}}
