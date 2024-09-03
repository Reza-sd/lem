package room

var OneComeUpdateSeats_Cases = TestCasesFunc{
	FuncName: "Room.OneComeUpdateSeats",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in middle type",
			Got: func() any {
				myRoom := SampleRoom.Middle_Name_3()
				myRoom.update.oneCome()
				return myRoom
			},
			Exp: func() any {
				r := newRoom()
				r.name = 3
				r.allSeats = 1
				r.usedSeats = 1
				r.connectionSlice = []Mtr{}
				return r
			},
		},
		//---------------------------------------
		{
			Des: "increase one UsedSeats when Room in End type",
			Got: func() any {
				myRoom := SampleRoom.End_Name_1()
				myRoom.update.oneCome()
				return myRoom
			},
			Exp: func() any {
				r := newRoom()
				r.name = 1
				r.allSeats = MaxSeatsStartEnd
				r.usedSeats = UsedSeatsStartEnd + 1
				r.connectionSlice = []Mtr{}

				return r

			},
		},
		//---------------------------------------
	}}
