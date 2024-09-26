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

// ------------------------------------
var logger = logstack.BuildNewLogger(pkgName_, ErrCodeDes, true, true)

//---------------------------------
