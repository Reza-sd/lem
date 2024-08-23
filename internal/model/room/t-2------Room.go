package room

import (
	teststack "main/pkg/teststack"
)

// ====================================

type TestCasesFunc = teststack.TestCasesFunc
type TestCase = teststack.TestCase
type AllTestCasesPkSt = teststack.AllTestCasesPkSt

type SampleRoomSt struct{}

var SampleRoom = SampleRoomSt{}

// ====================================
// --------------------------------------
func (s *SampleRoomSt) End_Name_1() Room {

	return Room{
		Name:            1,
		AllSeats:        MaxSeatsStartEnd,
		UsedSeats:       UsedSeatsStartEnd,
		ConnectionSlice: []Mtr{},
	}
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() Room {

	return Room{
		Name:            0,
		AllSeats:        MaxSeatsStartEnd,
		UsedSeats:       UsedSeatsStartEnd,
		ConnectionSlice: []Mtr{},
	}
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() Room {

	return Room{
		Name:            3,
		AllSeats:        1,
		UsedSeats:       0,
		ConnectionSlice: []Mtr{},
	}
}

// --------------------------------------
