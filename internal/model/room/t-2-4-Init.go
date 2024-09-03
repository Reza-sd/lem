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
				r := &room{
					name:            0,
					allSeats:        MaxSeatsStartEnd,
					usedSeats:       UsedSeatsStartEnd,
					connectionSlice: []Mtr{1, 2, 3},
				}
				r.get = getter{room: r}
				r.set = setter{room: r}
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
				r := &room{
					name:            1,
					allSeats:        1,
					usedSeats:       0,
					connectionSlice: []Mtr{1, 2, 3},
				}
				r.get = getter{room: r}
				r.set = setter{room: r}
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
				r := &room{
					name:            5,
					allSeats:        MaxSeatsStartEnd,
					usedSeats:       UsedSeatsStartEnd,
					connectionSlice: []Mtr{1, 2, 3},
				}
				r.get = getter{room: r}
				r.set = setter{room: r}
				return r
			},
		},
		//---------------------------------------
	}}
