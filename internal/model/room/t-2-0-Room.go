package room

import (
	ts "main/pkg/teststack"
)

// ====================================

type TestCasesForFunc = ts.TestCasesforFunc
type TestCase = ts.TestCase

type SampleRoomSt struct{}

var SampleRoom = SampleRoomSt{}

// ====================================
func (s *SampleRoomSt) End_Name_1() *room {
	return newPlainRoom().
		SetName(1).
		SetAllSeats(MaxSeatsStartEnd).
		SetUsedSeats(UsedSeatsStartEnd).
		SetConnectionSlice([]mtr{})
}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	return newPlainRoom().
		SetName(0).
		SetAllSeats(MaxSeatsStartEnd).
		SetUsedSeats(UsedSeatsStartEnd).
		SetConnectionSlice([]mtr{})
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	return newPlainRoom().
		SetName(3).
		SetAllSeats(1).
		SetUsedSeats(0).
		SetConnectionSlice([]mtr{})
}

// ====================================================================
var test_roomBuilder_Cases = TestCasesForFunc{
	FuncName: "Room.roomBuilder",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "if name=startRoomName means start home",
			Got: newRuledRoom(rmBuildArg{name: startRoomName, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().
				SetName(startRoomName).
				SetAllSeats(MaxSeatsStartEnd).
				SetUsedSeats(UsedSeatsStartEnd).
				SetConnectionSlice([]mtr{1, 2, 3}),
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "non start or end room",
			Got: newRuledRoom(rmBuildArg{name: 3, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().
				SetName(3).
				SetAllSeats(1).
				SetUsedSeats(0).
				SetConnectionSlice([]mtr{1, 2, 3}),
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "end room",
			Got: newRuledRoom(rmBuildArg{name: 5, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().
				SetName(5).
				SetAllSeats(MaxSeatsStartEnd).
				SetUsedSeats(UsedSeatsStartEnd).
				SetConnectionSlice([]mtr{1, 2, 3}),
		},
		//---------------------------------------
	}}

//====================================================================
