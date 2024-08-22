package room

type SampleRoom struct{}

func (s *SampleRoom) Sample_Room_1_5000_1000() *Room {

	return &Room{
		Name:            1,
		AllSeats:        5000,
		UsedSeats:       1000,
		ConnectionSlice: []Mtr{2, 3, 4},
	}
}
