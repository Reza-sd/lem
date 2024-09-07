package room

// =========================================
type RT = uint16
type ET = uint8 //error type

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
	Null ET = iota
	GetOneRandomNextRoom
	GetOneRandomNextRoom10
	GetOneRandomNextRoom20

	UpdateOneCome
	UpdateOneCome10

	UpdateOneLeave
	UpdateOneLeave10
)

var StatusCodeDescription = map[ET]string{ //for log purpose
	GetOneRandomNextRoom:   "GetOneRandomNextRoom",
	GetOneRandomNextRoom10: "  ",
	GetOneRandomNextRoom20: "  ",

	UpdateOneCome:   "UpdateOneCome",
	UpdateOneCome10: "",

	UpdateOneLeave:   "UpdateOneLeave",
	UpdateOneLeave10: "",
}

//=========================================
