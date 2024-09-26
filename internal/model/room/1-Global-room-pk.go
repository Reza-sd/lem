package room

import "main/pkg/logstack"

// =========================================
type rT = uint16
type errT = uint8 //error type
// -----------------------------------------
const (
	pkgName_           string = "room"
	MaxSeatsStartEnd   rT     = 5000
	UsedSeatsStartEnd  rT     = 1000
	AllSeatsNormalRoom rT     = 1
	startRoomName      rT     = 0 //always 1 (use 0 as null)

	//Rules
	maxName               = 100
	maxLenConnectionSlice = 5
)

// ========================================
const ( //func or method error code
	Null errT = iota
	PkgName

	NewRuledRoom_0
	NewRuledRoom_10
	NewRuledRoom_20
	NewRuledRoom_30
	NewRuledRoom_40
	NewRuledRoom_50
	NewRuledRoom_60

	//set
	set_name
	set_name_10
	//Get
	Get_OneRandomNextRoom    //func name
	Get_OneRandomNextRoom_10 //EmptyConnectionSlice

	//Act
	Act_UpdateOneCome
	Act_UpdateOneCome_10 //OverCap

	Act_UpdateOneLeave
	Act_UpdateOneLeave_10 //OverCap
)

var ErrCodeDes = map[errT]string{ //for log purpose
	Null:    "Null",
	PkgName: pkgName_,
	//room builder
	NewRuledRoom_0:  _NewRuledRoom_0,
	NewRuledRoom_10: _NewRuledRoom_10,
	NewRuledRoom_20: _NewRuledRoom_20,
	NewRuledRoom_30: _NewRuledRoom_30,
	NewRuledRoom_40: _NewRuledRoom_40,
	NewRuledRoom_50: _NewRuledRoom_50,
	NewRuledRoom_60: _NewRuledRoom_60,
	//Set
	set_name:    _set_name,
	set_name_10: _set_name_10,
	//Get
	Get_OneRandomNextRoom:    _Get_OneRandomNextRoom,
	Get_OneRandomNextRoom_10: _Get_OneRandomNextRoom_10,

	//Act
	Act_UpdateOneCome:    _Act_UpdateOneCome,
	Act_UpdateOneCome_10: _Act_UpdateOneCome10,

	Act_UpdateOneLeave:    _Act_UpdateOneLeave,
	Act_UpdateOneLeave_10: _Act_UpdateOneLeave10,
}

//------------------------------------
var logger = logstack.BuildNewLogger(pkgName_, ErrCodeDes, true, true)
//---------------------------------