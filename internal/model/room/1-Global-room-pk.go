package room

// =========================================
type m = uint16
type e = uint8 //error type

// -----------------------------------------
const (
	PkgName            string = "room"
	MaxSeatsStartEnd   m      = 5000
	UsedSeatsStartEnd  m      = 1000
	AllSeatsNormalRoom m      = 1
	startRoomName      m      = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

// -----------------------------------------
const ( //func or method status code
	Null = iota
	GetOneRandomNextRoom
	GetOneRandomNextRoom10
	GetOneRandomNextRoom20

	UpdateOneCome
	Room_set_oneCome_code_10

	UpdateOneLeave
	UpdateOneLeave10
)

//=========================================
