package room

var Init_Cases = TestCasesFunc{
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
				AllSeats:        MaxSeatsStartEnd,
				UsedSeats:       UsedSeatsStartEnd,
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
				AllSeats:        MaxSeatsStartEnd,
				UsedSeats:       UsedSeatsStartEnd,
				ConnectionSlice: []Mtr{1, 2, 3},
			},
		},
		//---------------------------------------
	}}
