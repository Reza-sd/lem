package room

type SampleRoom struct{}

func (s *SampleRoom) Sample_Room_End() *Room {

	return &Room{
		Name:            1,
		AllSeats:        5000,
		UsedSeats:       1000,
		ConnectionSlice: []Mtr{2, 3, 4},
	}
}

// ----------------
func (s *SampleRoom) Sample_Room_Start_name_0() *Room {

	return &Room{
		Name:            0,
		AllSeats:        5000,
		UsedSeats:       1000,
		ConnectionSlice: []Mtr{1, 2, 3},
	}
}
// ----------------
func (s *SampleRoom) Sample_Room_Middle() *Room {

	return &Room{
		Name:            3,
		AllSeats:        1,
		UsedSeats:       0,
		ConnectionSlice: []Mtr{2, 4},
	}
}
// ----------------