package room


type SampleRoom struct{}

func (s *SampleRoom) Sample_Room_1 () *Room {

	return &Room{
		Name: Mtr(1),
		AllSeats: Mtr(100),
		UsedSeats: Mtr(0),
		ConnectionSlice: []Mtr{2,3,4},
	}
}