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
	return newPlainRoom().set.
		name(1).
		allSeats(MaxSeatsStartEnd).
		usedSeats(UsedSeatsStartEnd).
		connectionSlice([]mtr{}).
		room

}

// --------------------------------------
func (s *SampleRoomSt) Start_Name_0() *room {
	return newPlainRoom().set.
		name(0).
		allSeats(MaxSeatsStartEnd).
		usedSeats(UsedSeatsStartEnd).
		connectionSlice([]mtr{}).
		room
}

// --------------------------------------
func (s *SampleRoomSt) Middle_Name_3() *room {
	return newPlainRoom().set.
		name(3).
		allSeats(1).
		usedSeats(0).
		connectionSlice([]mtr{}).
		room
}

// --------------------------------------
// ====================================================================
var test_roomBuilder_Cases = TestCasesForFunc{
	FuncName: "Room.roomBuilder",
	//Skip: true,
	TestCases: []TestCase{
		//---------------------------------------
		{
			Des: "if name=startRoomName means start home",
			Got: newRuledRoom(rmBuildArg{name: startRoomName, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().set.
				name(startRoomName).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd).
				connectionSlice([]mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "non start or end room",
			Got: newRuledRoom(rmBuildArg{name: 3, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().set.
				name(3).
				allSeats(1).
				usedSeats(0).
				connectionSlice([]mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
		{
			//Skip: true,
			Des: "end room",
			Got: newRuledRoom(rmBuildArg{name: 5, endRoomName: 5, connectionSlice: []mtr{1, 2, 3}}),
			Exp: newPlainRoom().set.
				name(5).
				allSeats(MaxSeatsStartEnd).
				usedSeats(UsedSeatsStartEnd).
				connectionSlice([]mtr{1, 2, 3}).
				room,
		},
		//---------------------------------------
	}}

//====================================================================
