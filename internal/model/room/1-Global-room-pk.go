package room

import (
	logstack "main/pkg/logstack"
)

// =========================================
type RT = uint16
type statCodeT = uint8 //error type
type statArrT = []statCodeT

// -----------------------------------------
const (
	_roomPk            string = "room"
	MaxSeatsStartEnd   RT     = 5000
	UsedSeatsStartEnd  RT     = 1000
	AllSeatsNormalRoom RT     = 1
	startRoomName      RT     = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// =====================================
// var stat=logstack.StatWapper
var Log = logstack

// ========================================
const ( //func or method status code
	Null statCodeT = iota
	PkgName
	//Get
	OneRandomNextRoom
	OneRandomNextRoom10 //EmptyConnectionSlice

	//Act
	UpdateOneCome
	UpdateOneCome10 //OverCap

	UpdateOneLeave
	UpdateOneLeave10 //OverCap
)

var CodeDes = map[statCodeT]string{ //for log purpose
	PkgName: _roomPk,
	//Get
	OneRandomNextRoom:   _OneRandomNextRoom,
	OneRandomNextRoom10: _OneRandomNextRoom10,

	//Act
	UpdateOneCome:   _UpdateOneCome,
	UpdateOneCome10: _UpdateOneCome10,

	UpdateOneLeave:   _UpdateOneLeave,
	UpdateOneLeave10: _UpdateOneLeave10,
}

// =========================================
