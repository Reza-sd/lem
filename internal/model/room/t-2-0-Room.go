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
func (s *SampleRoomSt) End_Name_1() *room {
	return newRoom().set.
		name(1).
		allSeats(MaxSeatsStartEnd).
		usedSeats(UsedSeatsStartEnd).
		connectionSlice([]Mtr{}).
		room

}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	return newRoom().set.
		name(0).
		allSeats(MaxSeatsStartEnd).
		usedSeats(UsedSeatsStartEnd).
		connectionSlice([]Mtr{}).
		room
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	return newRoom().set.
		name(3).
		allSeats(1).
		usedSeats(0).
		connectionSlice([]Mtr{}).
		room
}

// --------------------------------------
