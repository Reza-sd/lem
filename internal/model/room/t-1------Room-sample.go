package room


type SampleRoom struct{}

func (s *SampleRoom) Sample_Room_1 () *Room {

	return &Room{
		Name: 1,
		AllSeats: 100,
		UsedSeats: 50,
		ConnectionSlice: []Mtr{2,3,4},
	}
}