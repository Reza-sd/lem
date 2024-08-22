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
func(s *SampleRoomSt) End()Room{

	return Room{
		Name:            1,
		AllSeats:        5000,
		UsedSeats:       1000,
		ConnectionSlice: []Mtr{2, 3, 4},
	}
} 
// --------------------------------------
func(s *SampleRoomSt) Start()Room{

	return Room{
		Name:            0,
		AllSeats:        5000,
		UsedSeats:       1000,
		ConnectionSlice: []Mtr{1, 2, 3},
	}
} 
// --------------------------------------
func(s *SampleRoomSt) Middle()Room{

	return Room{
		Name:            3,
		AllSeats:        1,
		UsedSeats:       0,
		ConnectionSlice: []Mtr{2, 4},
	}
} 

// --------------------------------------
