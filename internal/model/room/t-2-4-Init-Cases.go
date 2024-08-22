package room

var Init_Test = TestCasesFunc{
	FuncName: "Room.Init",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "if name=0 means start home",
			Got: func() any {
				myRoom := Room{}
				myRoom.Init(0, 5, []Mtr{1, 2, 3})
				return &myRoom
			},
			Exp: &Room{
				Name:            0,
				AllSeats:        5000,
				UsedSeats:       1000,
				ConnectionSlice: []Mtr{1, 2, 3},
			},
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "non start or end room",
			Got: func() any {
				myRoom := Room{}
				myRoom.Init(1, 5, []Mtr{1, 2, 3})
				return &myRoom
			},
			Exp: &Room{
				Name:            1,
				AllSeats:        1,
				UsedSeats:       0,
				ConnectionSlice: []Mtr{1, 2, 3},
			},
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "end room",
			Got: func() any {
				myRoom := Room{}
				myRoom.Init(5, 5, []Mtr{1, 2, 3})
				return &myRoom
			},
			Exp: &Room{
				Name:            5,
				AllSeats:        5000,
				UsedSeats:       1000,
				ConnectionSlice: []Mtr{1, 2, 3},
			},
		},
		//---------------------------------------
	}}
