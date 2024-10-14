package room

// ====================================

type SampleRoomSt struct{}

var SampleRoom = SampleRoomSt{}

// ====================================
func (s *SampleRoomSt) End_Name_1() *roomT {
	r := &roomT{}
	// 	data: roomTdataT{
	// 		name:            1,
	// 		allSeats:        MaxSeatsStartEnd,
	// 		usedSeats:       UsedSeatsStartEnd,
	// 		connectionSlice: []rT{},
	// 	},
	// }
	r.data.name = 1
	r.data.allSeats = _MAX_SEATS_START_END
	r.data.usedSeats = _USED_SEATS_START_END
	r.data.connectionSlice = []rT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *roomT {
	r := &roomT{}
	// 	data: roomTdataT{
	// 		name:            0,
	// 		allSeats:        MaxSeatsStartEnd,
	// 		usedSeats:       UsedSeatsStartEnd,
	// 		connectionSlice: []rT{},
	// 	},
	// }
	r.data.name = 0
	r.data.allSeats = _MAX_SEATS_START_END
	r.data.usedSeats = _USED_SEATS_START_END
	r.data.connectionSlice = []rT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r

}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *roomT {
	r := &roomT{}
	// 	data: roomTdataT{
	// 		name:            3,
	// 		allSeats:        1,
	// 		usedSeats:       0,
	// 		connectionSlice: []rT{},
	// 	},
	// }
	r.data.name = 3
	r.data.allSeats = 1
	r.data.usedSeats = 0
	r.data.connectionSlice = []rT{}

	r.Get.room = r
	r.set.room = r
	r.Act.room = r
	return r

}

//=================================================
