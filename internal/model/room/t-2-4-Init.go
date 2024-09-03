package room

var Init_Cases = TestCasesFunc{
	FuncName: "Room.Init",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "if name=0 means start home",
			Got: func() any {
				myRoom := newRoom()
				myRoom.initiator(0, 5, []Mtr{1, 2, 3})
				return myRoom
			},
			Exp: func() any {
				r := newRoom()
				r.name = 0
				r.allSeats = MaxSeatsStartEnd
				r.usedSeats = UsedSeatsStartEnd
				r.connectionSlice = []Mtr{1, 2, 3}

				return r

			},
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "non start or end room",
			Got: func() any {
				myRoom := newRoom()
				myRoom.initiator(1, 5, []Mtr{1, 2, 3})
				return myRoom
			},
			Exp: func() any {
				r := newRoom()
				r.name = 1
				r.allSeats = 1
				r.usedSeats = 0
				r.connectionSlice = []Mtr{1, 2, 3}

				return r

			},
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "end room",
			Got: func() any {
				myRoom := newRoom()
				myRoom.initiator(5, 5, []Mtr{1, 2, 3})
				return myRoom
			},
			Exp: func() any {
				r := newRoom()
				r.name = 5
				r.allSeats = MaxSeatsStartEnd
				r.usedSeats = UsedSeatsStartEnd
				r.connectionSlice = []Mtr{1, 2, 3}

				return r
			},
		},
		//---------------------------------------
	}}
