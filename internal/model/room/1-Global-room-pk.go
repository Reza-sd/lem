package room

import "main/pkg/logstack"

// =========================================
type RT = uint16
type errT = uint8 //error type
type errArrT = []errT

// -----------------------------------------
const (
	_pkgName           string = "room"
	MaxSeatsStartEnd   RT     = 5000
	UsedSeatsStartEnd  RT     = 1000
	AllSeatsNormalRoom RT     = 1
	startRoomName      RT     = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// =====================================
// var stat=logstack.StatWapper
//var Log = logstack

// ========================================
const ( //func or method error code
	Null errT = iota
	PkgName
	//Get
	OneRandomNextRoom   //func name
	OneRandomNextRoom10 //EmptyConnectionSlice

	//Act
	UpdateOneCome
	UpdateOneCome10 //OverCap

	UpdateOneLeave
	UpdateOneLeave10 //OverCap
)

var ErrCodeDes = map[errT]string{ //for log purpose
	PkgName: _pkgName,
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
var l = logstack.BuildNewLogger(_pkgName, ErrCodeDes, true, true)
