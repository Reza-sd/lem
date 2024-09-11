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
func stat(statusCode statCodeT, previousStatusCodesSlice []statCodeT) statArrT {
	if statusCode == 0 && len(previousStatusCodesSlice) == 0 { //nill or empty
		return nil
	}

	if previousStatusCodesSlice == nil {
		previousStatusCodesSlice = statArrT{}
	}
	return append(previousStatusCodesSlice, statusCode)
}

//========================================
