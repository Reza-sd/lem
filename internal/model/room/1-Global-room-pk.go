package room

const (
	PkgName            string = "room"
	MaxSeatsStartEnd   mtr    = 5000
	UsedSeatsStartEnd  mtr    = 1000
	AllSeatsNormalRoom mtr    = 1
	startRoomName      mtr    = 1 //always 1 (use 0 as null)

	maxLenConnectionSlice = 5
)

var ErrorsHolder = make(map[uint8]uint8)

const ( //func or method status code
	//-----getter--------------------
	null              = 0
	OneRandomNextRoom = iota + 10
	OneRandomNextRoom_code_10
	OneRandomNextRoom_Code_20

	oneCome
	oneCome_code_10

	oneLeave
	oneLeave_code_10

	//-----setter--------------------

	//--------------------------------
)

// --------------------
type mtr = uint16

// -------------------
