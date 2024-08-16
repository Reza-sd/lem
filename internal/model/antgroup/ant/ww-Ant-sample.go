package ant

type SampleAnt struct{}

// -------Ant type example-----------
func (s *SampleAnt) Ant1() Ant {

	return Ant{
		Name:            1,
		CurrentRoomName: 0,
		VisitedRoomsArr: []Mta{0},
		StepNumber:      0,
	}

}

// ----------------
func (s *SampleAnt) Ant2() Ant {
	return Ant{
		Name:            2,
		CurrentRoomName: 0,
		VisitedRoomsArr: []Mta{0},
		StepNumber:      0,
	}
}

//------------------------
