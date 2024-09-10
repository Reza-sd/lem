package room

// =========================================
type RT = uint16
type statCodeT = uint8 //error type
type statArrT = []statCodeT

// -----------------------------------------
const (
	PkgName            string = "room"
	MaxSeatsStartEnd   RT     = 5000
	UsedSeatsStartEnd  RT     = 1000
	AllSeatsNormalRoom RT     = 1
	startRoomName      RT     = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// -----------------------------------------
const ( //func or method status code
	Null statCodeT = iota
	Get_OneRandomNextRoom
	Get_OneRandomNextRoom_EmptyConnectionSlice

	Act_UpdateOneCome
	Act_UpdateOneCome_OverCap

	Act_UpdateOneLeave
	Act_UpdateOneLeave_OverCap
)

var StatusCodeDescription = map[statCodeT]string{ //for log purpose
	Get_OneRandomNextRoom:                      "Get_OneRandomNextRoom",
	Get_OneRandomNextRoom_EmptyConnectionSlice: "EmptyConnectionSlice",

	Act_UpdateOneCome:         "Act_UpdateOneCome",
	Act_UpdateOneCome_OverCap: "no more free seat",

	Act_UpdateOneLeave:         "Act_UpdateOneLeave",
	Act_UpdateOneLeave_OverCap: "no one is here!",
}

// =========================================
func stat(statusCode statCodeT, previousStatusCodesSlice statArrT) statArrT {
	if statusCode == 0 && len(previousStatusCodesSlice) == 0 { //nill or empty
		return nil
	}

	if previousStatusCodesSlice == nil {
		previousStatusCodesSlice = statArrT{}
	}
	return append(previousStatusCodesSlice, statusCode)
}

//========================================
