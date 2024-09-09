package room

// ====================================

type SampleRoomSt struct{}

var SampleRoom = SampleRoomSt{}

// ====================================
func (s *SampleRoomSt) End_Name_1() *room {
	return newPlainRoom().
		SetName(1).
		SetAllSeats(MaxSeatsStartEnd).
		SetUsedSeats(UsedSeatsStartEnd).
		SetConnectionSlice([]RT{})
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	return newPlainRoom().
		SetName(0).
		SetAllSeats(MaxSeatsStartEnd).
		SetUsedSeats(UsedSeatsStartEnd).
		SetConnectionSlice([]RT{})
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	return newPlainRoom().
		SetName(3).
		SetAllSeats(1).
		SetUsedSeats(0).
		SetConnectionSlice([]RT{})
}

// ====================================================================
