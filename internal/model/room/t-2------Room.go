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
func (s *SampleRoomSt) End_Name_1() *Room {

	r := &Room{
		name:            1,
		allSeats:        MaxSeatsStartEnd,
		usedSeats:       UsedSeatsStartEnd,
		connectionSlice: []Mtr{},
	}
	r.get = getter{room: r}
	r.set = setter{room: r}
	return r
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *Room {

	r := &Room{
		name:            0,
		allSeats:        MaxSeatsStartEnd,
		usedSeats:       UsedSeatsStartEnd,
		connectionSlice: []Mtr{},
	}
	r.get = getter{room: r}
	r.set = setter{room: r}
	return r
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *Room {

	r := &Room{
		name:            3,
		allSeats:        1,
		usedSeats:       0,
		connectionSlice: []Mtr{},
	}
	r.get = getter{room: r}
	r.set = setter{room: r}
	return r
}

// --------------------------------------
