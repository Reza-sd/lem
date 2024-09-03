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
func (s *SampleRoomSt) End_Name_1() *room {
	r := newRoom()
	r.name = 1
	r.allSeats = MaxSeatsStartEnd
	r.usedSeats = UsedSeatsStartEnd
	r.connectionSlice = []Mtr{}

	return r
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	r := newRoom()
	r.name = 0
	r.allSeats = MaxSeatsStartEnd
	r.usedSeats = UsedSeatsStartEnd
	r.connectionSlice = []Mtr{}

	return r
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	r := newRoom()
	r.name = 3
	r.allSeats = 1
	r.usedSeats = 0
	r.connectionSlice = []Mtr{}

	return r
}

// --------------------------------------
