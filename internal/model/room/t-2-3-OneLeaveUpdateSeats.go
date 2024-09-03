package room

var OneLeaveUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneLeaveUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "Do nothing if usedSeats=0",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.OneLeaveUpdateSeats()
				return myRoom
			},
			Exp: func() any {
				r := &Room{
					name:            3,
					allSeats:        1,
					usedSeats:       0,
					connectionSlice: []Mtr{},
				}
				r.get = getter{room: r}
				r.set = setter{room: r}
				return r

			},
			// Room{
			// 	name:            3,
			// 	allSeats:        1,
			// 	usedSeats:       0,
			// 	connectionSlice: []Mtr{},
			// },
		},
		//---------------------------------------
		{
			Des: "decrease one UsedSeats when Room in Start type",
			Got: func() any {
				myRoom := SampleRoom.Start_Name_0()
				myRoom.OneLeaveUpdateSeats()
				return myRoom
			},
			Exp: func() any {
				r := &Room{
					name:            0,
					allSeats:        MaxSeatsStartEnd,
					usedSeats:       UsedSeatsStartEnd - 1,
					connectionSlice: []Mtr{},
				}
				r.get = getter{room: r}
				r.set = setter{room: r}
				return r

			},
			// Room{
			// 	name:            0,
			// 	allSeats:        MaxSeatsStartEnd,
			// 	usedSeats:       UsedSeatsStartEnd - 1,
			// 	connectionSlice: []Mtr{},
			// },
		},
		//---------------------------------------
	}}
