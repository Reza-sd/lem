package room

import "main/pkg/logstack"

// =========================================
type rT = uint16
type errT = uint8 //error type
// -----------------------------------------
const (
	_PKG_NAME              string = "room"
	_MAX_SEATS_START_END   rT     = 5000
	_USED_SEATS_START_END  rT     = 1000
	_ALL_SEATS_NORMAL_ROOM rT     = 1
	_START_ROOM_NAME       rT     = 0 //always 1 (use 0 as null)

	_EMPTY_USED_SEATS       = 0
	_EMPTY_CONNECTION_SLICE = 0
	//Rules
	_MAX_NAME                 = 100
	_MAX_LEN_CONNECTION_SLICE = 5
)

// ------------------------------------
var logger = logstack.Help_BuildNewLogger(_PKG_NAME, ErrCodeDb, true, true)

//var kir = logstack

//---------------------------------
